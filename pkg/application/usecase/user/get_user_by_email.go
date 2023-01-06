package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type GetUserById struct {
	Repository repository.UserRepository
}

func (gu *GetUserById) Execute(id string) (entity.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("Error to parse uuid: %v", err.Error())
	}

	u, err := gu.Repository.GetById(uid)
	if err != nil {
		return entity.User{}, fmt.Errorf("Error to get user: %v", err.Error())
	}

	return u, nil
}
