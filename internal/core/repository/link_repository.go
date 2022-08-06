package repository

type LinkRepository interface {
	GetLinkByID(linkID uint) (*Link, error)
	GetLinkByURL(url string) (*Link, error)
	GetLinksByUserID(userID uint) ([]*Link, error)
	CreateLink(link *Link) error
	UpdateLink(link *Link) error
	DeactivateLink(link *Link) error
}
