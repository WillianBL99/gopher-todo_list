package repositories

import (
	"github.com/google/uuid"
	"github.com/uilianlago/API-To-Do-List/application/entities"
)

type ToDoListRepository interface {
	GetAll() ([]entities.Task, error)
	GetAllDone() ([]entities.Task, error)
	GetAllUndone() ([]entities.Task, error)
	GetOne(id uuid.UUID) (entities.Task, error)
	Complete(id uuid.UUID) error
	Save(t *entities.Task) error
	Delete(id uuid.UUID) error
	Update(t *entities.Task) error
}
