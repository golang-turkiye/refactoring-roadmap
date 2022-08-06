package gormDB_test

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository/gormDB"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/database/localDB"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
)

func TestNewLinkRepository(t *testing.T) {
	db, err := localDB.Connection("test_new_link_repository.db")
	assert.Nil(t, err)
	linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
	assert.Nil(t, err)
	assert.NotNil(t, linkRepo)
	assert.Nil(t, os.Remove("test_new_link_repository.db"))
}

func TestLinkRepository_CreateLink(t *testing.T) {
	db, err := localDB.Connection("test_create_link.db")
	if err != nil {
		t.Error(err)
	}
	t.Run("Success", func(t *testing.T) {
		linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
		if err != nil {
			t.Error(err)
		}
		link := domain.Link{
			LongUrl:  "https://www.google.com",
			ShortUrl: "shortlink.com/shorten",
		}
		assert.Nil(t, linkRepo.CreateLink(&link))
	})
	assert.Nil(t, os.Remove("test_create_link.db"))
}

func TestLinkRepository_GetLinkByID(t *testing.T) {
	db, err := localDB.Connection("test_get_link_by_id.db")
	assert.Nil(t, err)
	linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
	assert.Nil(t, err)
	link := domain.Link{
		LongUrl:  "https://www.google.com",
		ShortUrl: "shortlink.com/shorten",
	}
	assert.Nil(t, linkRepo.CreateLink(&link))
	t.Run("Success", func(t *testing.T) {
		link, err := linkRepo.GetLinkByID(link.ID)
		assert.Nil(t, err)
		assert.Equal(t, link.LongUrl, "https://www.google.com")
		assert.Equal(t, link.ShortUrl, "shortlink.com/shorten")
	})
	assert.Nil(t, os.Remove("test_get_link_by_id.db"))
}

func TestLinkRepository_GetLinksByUserID(t *testing.T) {
	db, err := localDB.Connection("test_get_links_by_user_id.db")
	assert.Nil(t, err)
	linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
	assert.Nil(t, err)
	links := []domain.Link{
		{
			OwnerID:  1,
			LongUrl:  "https://www.google.com",
			ShortUrl: "shortlink.com/shorten",
		},
		{
			OwnerID:  1,
			LongUrl:  "https://www.google2.com",
			ShortUrl: "shortlink.com/shorten2",
		},
		{
			OwnerID:  2,
			LongUrl:  "https://www.google2.com",
			ShortUrl: "shortlink.com/shorten2",
		},
	}
	for _, l := range links {
		assert.Nil(t, linkRepo.CreateLink(&l))
	}
	t.Run("Success", func(t *testing.T) {
		links, err := linkRepo.GetLinksByUserID(links[0].OwnerID)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 2)
	})
	assert.Nil(t, os.Remove("test_get_links_by_user_id.db"))
}

func TestLinkRepository_DeactivateLink(t *testing.T) {
	db, err := localDB.Connection("test_deactivate_link.db")
	assert.Nil(t, err)
	linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
	assert.Nil(t, err)
	link := domain.Link{
		Model: gorm.Model{
			ID: 1,
		},
		LongUrl:   "https://www.google.com",
		ShortUrl:  "shortlink.com/shorten",
		IsDeleted: false,
	}
	assert.Nil(t, linkRepo.CreateLink(&link))
	t.Run("Success", func(t *testing.T) {
		err := linkRepo.DeactivateLink(&link)
		assert.Nil(t, err)
		link, err := linkRepo.GetLinkByID(link.ID)
		assert.Nil(t, err)
		assert.True(t, link.IsDeleted)
	})
	assert.Nil(t, os.Remove("test_deactivate_link.db"))
}

func TestLinkRepository_GetLinkByURL(t *testing.T) {
	db, err := localDB.Connection("test_get_link_by_url.db")
	assert.Nil(t, err)
	linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
	assert.Nil(t, err)
	link := domain.Link{
		Model: gorm.Model{
			ID: 1,
		},
		LongUrl:   "https://www.google.com",
		ShortUrl:  "shortlink.com/shorten",
		IsDeleted: false,
	}
	assert.Nil(t, linkRepo.CreateLink(&link))
	t.Run("Success", func(t *testing.T) {
		link, err := linkRepo.GetLinkByURL(link.ShortUrl)
		assert.Nil(t, err)
		assert.Equal(t, link.LongUrl, "https://www.google.com")
		assert.Equal(t, link.ShortUrl, "shortlink.com/shorten")
	})
	assert.Nil(t, os.Remove("test_get_link_by_url.db"))
}

func TestLinkRepository_UpdateLink(t *testing.T) {
	db, err := localDB.Connection("test_update_link.db")
	assert.Nil(t, err)
	linkRepo, err := gormDB.NewLinkRepository(db, logrus.New())
	assert.Nil(t, err)
	link := domain.Link{
		Model: gorm.Model{
			ID: 1,
		},
		LongUrl:   "https://www.google.com",
		ShortUrl:  "shortlink.com/shorten",
		IsDeleted: false,
	}
	assert.Nil(t, linkRepo.CreateLink(&link))
	t.Run("Success", func(t *testing.T) {
		link.ShortUrl = "shortlink.com/shorten2"
		assert.Nil(t, linkRepo.UpdateLink(&link))
		link, err := linkRepo.GetLinkByID(link.ID)
		assert.Nil(t, err)
		assert.Equal(t, link.ShortUrl, "shortlink.com/shorten2")
	})
	assert.Nil(t, os.Remove("test_update_link.db"))
}
