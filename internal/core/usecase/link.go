package usecase

type LinkResponse struct {
	ID       uint   `json:"id"`
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}

type LinkCreateRequest struct {
	URL string `json:"url"`
}
