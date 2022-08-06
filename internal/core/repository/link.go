package repository

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"

// LinkRepository is the repository for links.
type LinkRepository interface {
	GetLinkByID(linkID uint) (*domain.Link, error)
	GetLinkByURL(url string) (*domain.Link, error)
	GetLinksByUserID(userID uint) ([]*domain.Link, error)
	CreateLink(link *domain.Link) error
	DeactivateLink(link *domain.Link) error
}
