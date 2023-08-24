package inmemory

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type TaskRepositoryInMemory struct {
	tasks []entity.Task
}

func (r *TaskRepositoryInMemory) GetAll(userId uuid.UUID) ([]entity.Task, *e.Error) {
	tks := make([]entity.Task, 0, 5)

	for _, t := range r.tasks {
		if userId == t.UserId && !t.DeletedAt.Valid {
			tks = append(tks, t)
		}
	}

	return tks, nil
}

func (r *TaskRepositoryInMemory) GetByStatus(userId uuid.UUID, status entity.Status) ([]entity.Task, *e.Error) {
	tks := make([]entity.Task, 0, 5)

	for _, t := range r.tasks {
		if userId == t.UserId && status == t.Status {
			tks = append(tks, t)
		}
	}

	return tks, nil
}

func (r *TaskRepositoryInMemory) GetById(id uuid.UUID) (entity.Task, *e.Error) {
	for _, t := range r.tasks {
		if id == t.Id {
			return t, nil
		}
	}

	return entity.Task{}, e.New().SetLayer(e.InfraTest).CustomError(e.TaskNotFound)
}

func (r *TaskRepositoryInMemory) GetByList(id uuid.UUID, listName string) ([]entity.Task, *e.Error) {
	tks := make([]entity.Task, 0, 5)

	for _, t := range r.tasks {
		if id == t.UserId && listName == t.List && !t.DeletedAt.Valid {
			tks = append(tks, t)
		}
	}

	return tks, nil
}

func (r *TaskRepositoryInMemory) Save(t *entity.Task) (entity.Task, *e.Error) {
	r.tasks = append(r.tasks, *t)

	return *t, nil
}

func (r *TaskRepositoryInMemory) Delete(id uuid.UUID) *e.Error {
	for n, t := range r.tasks {
		if id == t.Id {
			if t.DeletedAt.Valid {
				return e.New().SetLayer(e.InfraTest).CustomError(e.Conflict)
			}
			r.tasks[n].DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
			break
		}
	}

	return nil
}

func (r *TaskRepositoryInMemory) Update(t *entity.Task) *e.Error {
	for n, rt := range r.tasks {
		if rt.Id == t.Id {
			r.tasks[n] = *t
			r.tasks[n].UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
		}
	}

	return nil
}
