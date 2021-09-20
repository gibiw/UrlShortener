package service

import (
	"github.com/gibiw/UrlShortener/internal/repository"
	util "github.com/gibiw/UrlShortener/pkg"
)

type LinkItemService struct {
	repo          repository.LinkItem
	lenghOfString int
}

func NewLinkItemService(repo repository.LinkItem, lenghOfString int) *LinkItemService {
	return &LinkItemService{repo: repo, lenghOfString: lenghOfString}
}

func (s *LinkItemService) Create(o string) (string, error) {

	link, err := s.repo.GetByUrl(o)

	if err == nil {
		return link.Modification, nil
	}

	if _, ok := err.(*util.NotFoundError); !ok {
		return "", err
	}

	mod := util.RandString(s.lenghOfString)

	var exitFor = true
	for exitFor {
		if ok, err := s.isHashExist(mod); err == nil && !ok {
			exitFor = false
		} else if err != nil && !ok {
			return "", err
		} else {
			mod = util.RandString(s.lenghOfString)
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
