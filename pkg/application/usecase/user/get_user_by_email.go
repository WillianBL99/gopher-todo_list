package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetUserById struct {
	Repository repository.UserRepository
}

func (gu *GetUserById) Execute(id string) (entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.User{}, herr.NewApp().InvalidUserId
	}

	u, err := gu.Repository.GetById(uid)
	if err != nil {
		return entity.User{}, err
	}

	return u, nil
}
