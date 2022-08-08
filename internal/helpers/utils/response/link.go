package responseutils

import "github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/domain"

type LinkResponse struct {
	ID       uint   `json:"id"`
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}

type LinkCreateRequest struct {
	URL string `json:"url"`
}

func MapLinkListResponse(links []domain.Link) []LinkResponse {
	var linkResponses []LinkResponse
	for _, link := range links {
		linkResponses = append(linkResponses, MapLinkResponse(link))
	}
	return linkResponses
}

func MapLinkResponse(link domain.Link) LinkResponse {
	return LinkResponse{
		ID:       link.Model.ID,
		URL:      link.LongUrl,
		ShortURL: link.ShortUrl,
	}
}
