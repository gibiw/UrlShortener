package service

import (
	link "github.com/gibiw/UrlShortener"
	"github.com/gibiw/UrlShortener/internal/repository"
)

type LinkItem interface {
	Create(o string) (string, error)
	GetByUrl(guid string) (link.LinkItem, error)
}

type Service struct {
	LinkItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		LinkItem: NewLinkItemService(repos.LinkItem),
	}
}
