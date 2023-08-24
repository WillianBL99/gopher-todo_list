package server

import (
	"github.com/willianbl99/todo-list_api/config"
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct {
	Cost int
}

func NewBcryptService() *BcryptService {
	return &BcryptService{
		Cost: int(config.NewAppConf().BCryptCost),
	}
}

func (s *BcryptService) Encrypt(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), s.Cost)
	return string(bytes), err
}

func (s *BcryptService) Compare(hashedpass string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(pass))
	return err == nil
}
