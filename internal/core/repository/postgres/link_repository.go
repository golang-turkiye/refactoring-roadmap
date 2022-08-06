package postgres

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

// GetLinkByID returns a link by its ID.
func (r *LinkRepository) GetLinkByID(linkID uint) (*domain.Link, error) {
	return nil, nil
}

// GetLinkByURL returns a link by its URL.
func (r *LinkRepository) GetLinkByURL(url string) (*domain.Link, error) {
	return nil, nil
}

// GetLinksByUserID returns a list of links by its user ID.
func (r *LinkRepository) GetLinksByUserID(userID uint) ([]*domain.Link, error) {
	return nil, nil
}

// CreateLink creates a new link.
func (r *LinkRepository) CreateLink(link *domain.Link) error {
	return nil
}

// UpdateLink updates an existing link.
func (r *LinkRepository) UpdateLink(link *domain.Link) error {
	return nil
}

// DeactivateLink deactivates an existing link.
func (r *LinkRepository) DeactivateLink(link *domain.Link) error {
	return nil
}
