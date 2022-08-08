package linkservice

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/helpers/utils/customerror"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/repository"
)

type linkService struct {
	linkRepository repository.LinkRepository
}

// New creates a new linkService.
func New(linkRepository repository.LinkRepository) *linkService {
	return &linkService{
		linkRepository: linkRepository,
	}
}

// GetLinkByID returns a link by its ID.
func (s *linkService) GetLinkByID(linkID uint) (*domain.Link, error) {
	return nil, nil
}

// GetLink returns a link.
func (s *linkService) GetLink(shortPath string) (*domain.Link, error) {
	if shortPath == "" {
		return nil, customerror.ErrInvalidShortPath
	}
	return s.linkRepository.GetLinkByURL(shortPath)
}

// GetAllLinks returns all links.
func (s *linkService) GetAllLinks(ownerID uint) ([]*domain.Link, error) {
	return nil, nil
}

// CreateLink creates a new link.
func (s *linkService) CreateLink(link *domain.Link) error {
	if link.LongUrl == "" {
		return customerror.ErrInvalidLongURL
	}
	return s.linkRepository.CreateLink(link)
}

// DeactivateLink deactivates a link.
func (s *linkService) DeactivateLink(linkID uint) error {
	if linkID == 0 {
		return customerror.ErrInvalidLinkID
	}
	link, err := s.linkRepository.GetLinkByID(linkID)
	if err != nil {
		return customerror.ErrInvalidLinkID
	}
	return s.linkRepository.DeactivateLink(link)
}
