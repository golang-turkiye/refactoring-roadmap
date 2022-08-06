package postgres

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LinkRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewLinkRepository(db *gorm.DB, logger *logrus.Logger) (*LinkRepository, error) {
	if err := db.AutoMigrate(&domain.Link{}); err != nil {
		logger.Error("Error while auto migrating link table")
		return nil, err
	}
	return &LinkRepository{db: db, logger: logger}, nil
}

// GetLinkByID returns a link by its ID.
func (r *LinkRepository) GetLinkByID(linkID uint) (*domain.Link, error) {
	tx := r.db.Begin()
	link := &domain.Link{}
	result := tx.First(link, linkID)
	if result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return link, nil
}

// GetLinkByURL returns a link by its URL.
func (r *LinkRepository) GetLinkByURL(url string) (*domain.Link, error) {
	tx := r.db.Begin()
	link := &domain.Link{}
	result := tx.Where("shorten_url_path = ?", url).First(link)
	if result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return link, nil
}

// GetLinksByUserID returns a list of links by its user ID.
func (r *LinkRepository) GetLinksByUserID(userID uint) ([]*domain.Link, error) {
	tx := r.db.Begin()
	links := []*domain.Link{}
	result := tx.Where("owner_id = ?", userID).Find(&links)
	if result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return links, nil
}

// CreateLink creates a new link.
func (r *LinkRepository) CreateLink(link *domain.Link) error {
	tx := r.db.Begin()
	result := tx.Create(link)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// UpdateLink updates an existing link.
func (r *LinkRepository) UpdateLink(link *domain.Link) error {
	tx := r.db.Begin()
	result := tx.Model(&domain.Link{}).Where("id = ?", link.ID).Updates(link)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}

// DeactivateLink deactivates an existing link.
func (r *LinkRepository) DeactivateLink(link *domain.Link) error {
	tx := r.db.Begin()
	result := tx.Model(&domain.Link{}).Where("id = ? and is_deleted = false", link.ID).Update("is_deleted", true)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	tx.Commit()
	return nil
}
