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

	link, err := s.repo.GetByUrl(o)

	if err == nil {
		return link.Modification, nil
	}

	if _, ok := err.(*util.NotFoundError); !ok {
		return "", err
	}

	mod := util.RandString(lenghOfString)

	var exitFor = true
	for exitFor {
		if ok, err := s.isHashExist(mod); err == nil && !ok {
			exitFor = false
		} else if err != nil && !ok {
			return "", err
		} else {
			mod = util.RandString(lenghOfString)
		}
	}

	res, err := s.repo.Create(o, mod)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *LinkItemService) GetByHash(hash string) (string, error) {
	return s.repo.GetByHash(hash)
}

func (s *LinkItemService) isHashExist(hash string) (bool, error) {
	_, err := s.repo.GetByHash(hash)

	if err == nil {
		return true, nil
	}

	if _, ok := err.(*util.NotFoundError); ok {
		return false, nil
	}

	return false, err
}
