package v1

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository"
)

type LinkService struct {
	linkRepository repository.LinkRepository
}

// NewLinkService creates a new LinkService.
func NewLinkService(linkRepository repository.LinkRepository) *LinkService {
	return &LinkService{
		linkRepository: linkRepository,
	}
}

func (s *LinkService) GetLinkByID(linkID uint) (*domain.Link, error) {
	return nil, nil
}
func (s *LinkService) GetLink(ownerID, shortPath string) (*domain.Link, error) {
	return nil, nil
}
func (s *LinkService) GetAllLinks(ownerID uint) ([]*domain.Link, error) {
	return nil, nil
}
func (s *LinkService) CreateLink(link *domain.Link) error {
	return nil
}
func (s *LinkService) DeactivateLink(linkID uint) error {
	return nil
}
