package repository

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type UserRepository interface {
	GetById(id uuid.UUID) (entity.User, *e.Error)
	GetByEmail(email string) (entity.User, *e.Error)
	Save(u *entity.User) *e.Error
}
