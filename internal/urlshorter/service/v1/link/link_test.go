package linkservice_test

import (
	"errors"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/helpers/utils/customerror"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/service/v1/link"
	"github.com/Golang-Turkiye/refactoring-roadmap/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestNewLinkService(t *testing.T) {
	linkRepo := new(mocks.MockLinkRepository)
	linkService := linkservice.New(linkRepo)
	assert.NotNil(t, linkService)
}

func TestLinkService_CreateLink(t *testing.T) {
	linkRepo := new(mocks.MockLinkRepository)
	linkService := linkservice.New(linkRepo)
	assert.NotNil(t, linkService)
	testCases := []struct {
		name     string
		link     *domain.Link
		expected error
	}{
		{
			name: "Success",
			link: &domain.Link{
				LongUrl:  "https://www.google.com",
				ShortUrl: "shortlink.com/shorten",
			},
			expected: nil,
		},
		{
			name: "Fail",
			link: &domain.Link{
				LongUrl:  "",
				ShortUrl: "shortlink.com/shorten",
			},
			expected: customerror.ErrInvalidLongURL,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			linkRepo.On("CreateLink", tc.link).Return(tc.expected)
			err := linkService.CreateLink(tc.link)
			assert.Equal(t, err, tc.expected)
		})
	}
}

func TestLinkService_DeactivateLink(t *testing.T) {
	linkRepo := new(mocks.MockLinkRepository)
	linkService := linkservice.New(linkRepo)
	assert.NotNil(t, linkService)
	testCases := []struct {
		name     string
		link     *domain.Link
		expected error
	}{
		{
			name: "Success by False Deactivate",
			link: &domain.Link{
				Model: gorm.Model{
					ID: 1,
				},
				LongUrl:   "https://www.google.com",
				ShortUrl:  "shortlink.com/shorten",
				IsDeleted: false,
			},
			expected: nil,
		},
		{
			name: "Fail by True Deactivate",
			link: &domain.Link{
				Model: gorm.Model{
					ID: 2,
				},
				LongUrl:   "",
				ShortUrl:  "shortlink.com/shorten",
				IsDeleted: true,
			},
			expected: customerror.ErrInvalidLinkID,
		},
		{
			name:     "Fail by Not Found",
			link:     &domain.Link{},
			expected: customerror.ErrInvalidLinkID,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			linkRepo.On("GetLinkByID", uint(1)).Return(tc.link, nil)
			linkRepo.On("GetLinkByID", uint(2)).Return(nil, errors.New("not found"))
			if !tc.link.IsDeleted {
				linkRepo.On("DeactivateLink", tc.link).Return(nil).Once()
			} else {
				linkRepo.On("DeactivateLink", tc.link).Return(tc.expected).Once()
			}
			err := linkService.DeactivateLink(tc.link.ID)
			assert.Equal(t, err, tc.expected)
		})
	}
}

func TestLinkService_GetLink(t *testing.T) {
	link := &domain.
		Link{
		Model: gorm.Model{
			ID: 1,
		},
		LongUrl:   "https://www.google.com",
		ShortUrl:  "shortlink.com/shorten",
		IsDeleted: false,
		OwnerID:   1,
	}
	linkRepo := new(mocks.MockLinkRepository)
	linkRepo.On("GetLinkByURL", "shortlink.com/shorten").Return(link, nil)
	linkRepo.On("GetLinkByURL", "").Return(nil, customerror.ErrInvalidShortPath)
	linkService := linkservice.New(linkRepo)
	assert.NotNil(t, linkService)
	testCases := []struct {
		name     string
		shorturl string
		link     *domain.Link
		expected error
	}{
		{
			name:     "Success",
			shorturl: "shortlink.com/shorten",
			link:     link,
			expected: nil,
		},
		{
			name:     "Fail",
			shorturl: "",
			link:     nil,
			expected: customerror.ErrInvalidShortPath,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			link, err := linkService.GetLink(tc.shorturl)
			assert.Equal(t, link, tc.link)
			assert.Equal(t, err, tc.expected)
		})
	}
}
