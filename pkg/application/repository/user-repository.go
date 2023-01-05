package repository

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

type UserRepository interface {
	GetById(id uuid.UUID) (*entity.User, error);
	GetByEmail(email string) (*entity.User, error);
	Save(u *entity.User) error;
}