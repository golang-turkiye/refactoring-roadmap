package v1

import (
	"errors"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/repository"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/usecase/constant"
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

// GetLinkByID returns a link by its ID.
func (s *LinkService) GetLinkByID(linkID uint) (*domain.Link, error) {
	return nil, nil
}

// GetLink returns a link.
func (s *LinkService) GetLink(shortPath string) (*domain.Link, error) {
	if shortPath == "" {
		return nil, errors.New(constant.ErrInvalidShortPath)
	}
	return s.linkRepository.GetLinkByURL(shortPath)
}

// GetAllLinks returns all links.
func (s *LinkService) GetAllLinks(ownerID uint) ([]*domain.Link, error) {
	return nil, nil
}

// CreateLink creates a new link.
func (s *LinkService) CreateLink(link *domain.Link) error {
	if link.LongUrl == "" {
		return errors.New(constant.ErrInvalidLongURL)
	}
	return s.linkRepository.CreateLink(link)
}

// DeactivateLink deactivates a link.
func (s *LinkService) DeactivateLink(linkID uint) error {
	if linkID == 0 {
		return errors.New(constant.ErrInvalidLinkID)
	}
	link, err := s.linkRepository.GetLinkByID(linkID)
	if err != nil {
		return errors.New(constant.ErrInvalidLinkID)
	}
	return s.linkRepository.DeactivateLink(link)
}
