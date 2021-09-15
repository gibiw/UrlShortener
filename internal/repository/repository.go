package repository

import link "github.com/gibiw/UrlShortener"

type LinkItem interface {
	GetByUrl(url string) (link.LinkItem, error)
	GetByHash(hash string) (string, error)
	Create(o, e string) (string, error)
}

type Repository struct {
	LinkItem
}

func NewRepository() *Repository {
	return &Repository{
		LinkItem: NewLinkItemRepository(),
	}
}
