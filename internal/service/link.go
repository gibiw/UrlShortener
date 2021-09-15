package service

import (
	link "github.com/gibiw/UrlShortener"
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

func (s *LinkItemService) GetByUrl(guid string) (link.LinkItem, error) {
	return s.repo.GetByUrl(guid)
}
