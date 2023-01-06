package server

import "golang.org/x/crypto/bcrypt"

type BcryptService struct{
	Cost int
}

func NewBcryptService() *BcryptService {
	return &BcryptService{
		Cost: 14,
	}
}

func (s *BcryptService) Encript(p string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), s.Cost)
	return string(bytes), err
}

func (s *BcryptService) Compare(hashedpass, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(pass))
	return err == nil
}