package inmemory

import (
	"errors"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

type UserRepositoryInMemory struct {
	users []entity.User
}

func (r *UserRepositoryInMemory) GetById(id uuid.UUID) (entity.User, error) {
	u := entity.User{}

	for _, user := range r.users {
		if id == user.Id {
			u = user
			break
		}
	}

	return u, nil
}

func (r *UserRepositoryInMemory) GetByEmail(email string) (entity.User, error) {
	for _, u := range r.users {
		if email == u.Email {
			return u, nil
		}
	}

	return entity.User{}, errors.New("User not found")
}

func (r *UserRepositoryInMemory) Save(u *entity.User) error {
	r.users = append(r.users, *u)

	return nil
}
