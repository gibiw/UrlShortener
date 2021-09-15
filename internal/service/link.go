package service

import (
	"github.com/gibiw/UrlShortener/internal/repository"
	util "github.com/gibiw/UrlShortener/pkg"
)

const lenghOfString int = 6

type LinkItemService struct {
	repo repository.LinkItem
}

func NewLinkItemService(repo repository.LinkItem) *LinkItemService {
	return &LinkItemService{repo: repo}
}

func (s *LinkItemService) Create(o string) (string, error) {

	mod := util.RandString(lenghOfString)
	res, err := s.repo.Create(o, mod)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *LinkItemService) GetByHash(hash string) (string, error) {
	return s.repo.GetByHash(hash)
}
