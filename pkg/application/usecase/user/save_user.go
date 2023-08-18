package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type SaveUser struct {
	UserRepository repository.UserRepository
}

func (s *SaveUser) Execute(name string, email string, password string) error {
	if _, err := s.UserRepository.GetByEmail(email); err == nil {
		return herr.NewApp().EmailAlreadyExists
	}

	bc := server.NewBcryptService()
	hashedPass, err := bc.Encrypt(password)
	if err != nil {
		return herr.NewApp().BadRequest
	}

	u := entity.NewUser(uuid.New(), name, email, hashedPass)

	if err := s.UserRepository.Save(u); err != nil {
		return err
	}

	return nil
}
