package inmemory

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type TaskRepositoryInMemory struct {
	tasks []entity.Task
}

func (r *TaskRepositoryInMemory) GetAll(userId uuid.UUID) ([]entity.Task, error) {
	tks := make([]entity.Task, 0, 5)

	for _, t := range r.tasks {
		if userId == t.UserId && !t.DeletedAt.Valid {
			tks = append(tks, t)
		}
	}

	return tks, nil
}

func (r *TaskRepositoryInMemory) GetByStatus(userId uuid.UUID, status entity.Status) ([]entity.Task, error) {
	tks := make([]entity.Task, 0, 5)

	for _, t := range r.tasks {
		if userId == t.UserId && status == t.Status {
			tks = append(tks, t)
		}
	}

	return tks, nil
}

func (r *TaskRepositoryInMemory) GetById(id uuid.UUID) (entity.Task, error) {
	for _, t := range r.tasks {
		if id == t.Id {
			return t, nil
		}
	}

	return entity.Task{}, herr.NewApp().TaskNotFound
}

func (r *TaskRepositoryInMemory) Save(t *entity.Task) error {
	r.tasks = append(r.tasks, *t)

	return nil
}

func (r *TaskRepositoryInMemory) Delete(id uuid.UUID) error {
	for n, t := range r.tasks {
		if id == t.Id {
			if t.DeletedAt.Valid {
				return herr.NewApp().Conflict
			}
			r.tasks[n].DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
			break
		}
	}

	return nil
}

func (r *TaskRepositoryInMemory) Update(t *entity.Task) error {
	for n, rt := range r.tasks {
		if rt.Id == t.Id {
			r.tasks[n] = *t
			r.tasks[n].UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
		}
	}

	return nil
}
