package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/uilianlago/API-To-Do-List/application/entities"
)

type InMemoryToDoListRepository struct {
	tasks []entities.Task
}

func (r *InMemoryToDoListRepository) GetAll() ([]entities.Task, error) {
	return r.tasks, nil
}

func (r *InMemoryToDoListRepository) GetAllDone() ([]entities.Task, error) {
	td := make([]entities.Task, len(r.tasks))
	n := 0

	for _, t := range r.tasks {
		if t.Done {
			td[n] = t
			n++
		}
	}

	return td, nil
}

func (r *InMemoryToDoListRepository) GetAllUndone() ([]entities.Task, error) {
	tu := make([]entities.Task, len(r.tasks))
	n := 0

	for _, t := range r.tasks {
		if !t.Done {
			tu[n] = t
			n++
		}
	}

	return tu, nil
}

func (r *InMemoryToDoListRepository) GetOne(id uuid.UUID) (entities.Task, error) {
	tk := entities.Task{}

	for _, t := range r.tasks {
		if id == t.Id {
			tk = t
			break
		}
	}

	return tk, nil
}

func (r *InMemoryToDoListRepository) Complete(id uuid.UUID) error {
	for _, t := range r.tasks {
		if id == t.Id {
			t.Done = true
		}
	}

	return nil
}

func (r *InMemoryToDoListRepository) Save(t *entities.Task) error {
	r.tasks = append(r.tasks, *t)

	return nil
}

func (r *InMemoryToDoListRepository) Delete(id uuid.UUID) error {
	for _, t := range r.tasks {
		if id == t.Id {
			t.DeletedAt = time.Now()
		}
	}

	return nil
}

func (r *InMemoryToDoListRepository) Update(t *entities.Task) error {
	for _, rt := range r.tasks {
		if rt.Id == t.Id {
			rt.Title = t.Title
			rt.Describe = t.Describe
		}
	}

	return nil
}
