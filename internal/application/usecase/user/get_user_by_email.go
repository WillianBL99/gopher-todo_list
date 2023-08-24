package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetUserById struct {
	Repository repository.UserRepository
}

func (gu *GetUserById) Execute(id string) (entity.User, *e.Error) {
	appErr := e.New().SetLayer(e.Application)
	uid, er := uuid.Parse(id)
	if er != nil {
		return entity.User{}, appErr.CustomError(e.InvalidId, er.Error())
	}

	u, err := gu.Repository.GetById(uid)
	if err != nil {
		return entity.User{}, appErr.CustomError(e.UserNotFound, err.ToSubErr()...)
	}

	return u, nil
}
