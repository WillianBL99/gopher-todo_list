package repository

import (
	"github.com/google/uuid"

	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

type TaskRepository interface {
	GetAll(userId uuid.UUID) ([]entity.Task, error)
	GetByStatus(userId uuid.UUID, status entity.Status) ([]entity.Task, error)
	GetById(id uuid.UUID) (entity.Task, error)
	Save(t *entity.Task) error
	Delete(id uuid.UUID) error
	Update(t *entity.Task) error
}
