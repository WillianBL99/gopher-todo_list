package repository

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type TaskRepository interface {
	GetAll(userId uuid.UUID) ([]entity.Task, *e.Error)
	GetByStatus(userId uuid.UUID, status entity.Status) ([]entity.Task, *e.Error)
	GetById(id uuid.UUID) (entity.Task, *e.Error)
	GetByList(userId uuid.UUID, listName string) ([]entity.Task, *e.Error)
	Save(t *entity.Task) (entity.Task, *e.Error)
	Delete(id uuid.UUID) *e.Error
	Update(t *entity.Task) *e.Error
}
