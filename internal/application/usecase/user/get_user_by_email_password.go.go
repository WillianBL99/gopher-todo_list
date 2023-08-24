package usecase

import (
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type GetUserByEmailPassword struct {
	UserRepository repository.UserRepository
}

func (gu *GetUserByEmailPassword) Execute(em string, p string) (entity.User, *e.Error) {
	u, err := gu.UserRepository.GetByEmail(em)
	if err != nil {
		return u, e.New().AppErr(e.EmailOrPasswordInvalid, err.ToSubErr()...)
	}

	bc := server.NewBcryptService()
	if !bc.Compare(u.Password, p) {
		return u, e.New().AppErr(e.InternalServerError)
	}

	return u, nil
}
