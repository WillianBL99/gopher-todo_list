package usecase

import (
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type GetUserByEmailPassword struct {
	UserRepository repository.UserRepository
}

func (gu *GetUserByEmailPassword) Execute(e string, p string) (entity.User, error) {
	u, err := gu.UserRepository.GetByEmail(e)
	if err != nil {
		return u, err
	}

	bc := server.NewBcryptService()
	if !bc.Compare(u.Password, p) {
		return u, herr.NewApp().EmailOrPasswordInvalid
	}

	return u, nil
}
