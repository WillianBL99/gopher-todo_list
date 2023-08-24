package inmemory

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type UserRepositoryInMemory struct {
	users []entity.User
}

func (r *UserRepositoryInMemory) GetById(id uuid.UUID) (entity.User, *e.Error) {
	for _, user := range r.users {
		if id == user.Id {
			return user, nil
		}
	}

	return entity.User{}, e.New().SetLayer(e.InfraTest).CustomError(e.UserNotFound)
}

func (r *UserRepositoryInMemory) GetByEmail(email string) (entity.User, *e.Error) {
	for _, u := range r.users {
		if email == u.Email {
			return u, nil
		}
	}

	return entity.User{}, e.New().SetLayer(e.InfraTest).CustomError(e.UserNotFound)
}

func (r *UserRepositoryInMemory) Save(u *entity.User) *e.Error {
	r.users = append(r.users, *u)

	return nil
}
