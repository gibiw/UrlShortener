package service

import (
	"crypto/md5"
	"encoding/hex"

	link "github.com/gibiw/UrlShortener"
	"github.com/gibiw/UrlShortener/internal/repository"
)

type LinkItemService struct {
	repo repository.LinkItem
}

func NewLinkItemService(repo repository.LinkItem) *LinkItemService {
	return &LinkItemService{repo: repo}
}

func (s *LinkItemService) Create(o string) (string, error) {

	mod := getMD5Hash(o)
	res, err := s.repo.Create(o, mod)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *LinkItemService) GetByUrl(guid string) (link.LinkItem, error) {
	return s.repo.GetByUrl(guid)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
