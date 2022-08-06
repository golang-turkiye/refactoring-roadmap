package service

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"

type LinkService interface {
	GetLink(ownerID, shortPath string) (*domain.Link, error)
	GetAllLinks(ownerID uint) ([]*domain.Link, error)
	CreateLink(link *domain.Link) error
	DeactivateLink(linkID uint) error
}
