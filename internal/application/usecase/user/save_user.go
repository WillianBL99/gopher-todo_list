package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type SaveUser struct {
	UserRepository repository.UserRepository
}

func (s *SaveUser) Execute(name string, email string, password string) *e.Error {
	appErr := e.New().SetLayer(e.Application)

	if name == "" || email == "" || password == "" {
		return appErr.CustomError(e.EmptyField)
	}

	if _, err := s.UserRepository.GetByEmail(email); err == nil {
		return appErr.CustomError(e.EmailAlreadyExists)
	}

	bc := server.NewBcryptService()
	hashedPass, er := bc.Encrypt(password)
	if er != nil {
		return appErr.CustomError(e.InternalServerError, er.Error())
	}

	u := entity.NewUser(uuid.New(), name, email, hashedPass)

	return s.UserRepository.Save(u)
}
