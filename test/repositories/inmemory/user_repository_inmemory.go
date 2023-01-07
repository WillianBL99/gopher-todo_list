package inmemory

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type UserRepositoryInMemory struct {
	users []entity.User
}

func (r *UserRepositoryInMemory) GetById(id uuid.UUID) (entity.User, error) {
	for _, user := range r.users {
		if id == user.Id {
			return user, nil
		}
	}

	return entity.User{}, herr.NewApp().UserNotFound
}

func (r *UserRepositoryInMemory) GetByEmail(email string) (entity.User, error) {
	for _, u := range r.users {
		if email == u.Email {
			return u, nil
		}
	}

	return entity.User{}, herr.NewApp().UserNotFound
}

func (r *UserRepositoryInMemory) Save(u *entity.User) error {
	r.users = append(r.users, *u)

	return nil
}
