package service

import (
	"github.com/gibiw/UrlShortener/internal/repository"
)

type LinkItem interface {
	Create(o string) (string, error)
	GetByHash(guid string) (string, error)
}

type Service struct {
	LinkItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		LinkItem: NewLinkItemService(repos.LinkItem),
	}
}
