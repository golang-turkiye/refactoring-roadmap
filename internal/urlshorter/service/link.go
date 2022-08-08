package service

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"

type LinkService interface {
	GetLink(shortPath string) (*domain.Link, error)
	GetAllLinks(ownerID uint) ([]*domain.Link, error)
	CreateLink(link *domain.Link) error
	DeactivateLink(linkID uint) error
}
