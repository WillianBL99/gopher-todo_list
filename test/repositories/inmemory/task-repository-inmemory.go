package inmemory

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

type TaskRepositoryInMemory struct {
	tasks []entity.Task
}

func (r *TaskRepositoryInMemory) GetAll(userId uuid.UUID) ([]entity.Task, error) {
	tks := make([]entity.Task, 0, 5)

	for _, t := range r.tasks {
		if userId == t.UserId {
			fmt.Printf("entrou no if: %v == %v\n", userId, t.UserId)
			tks = append(tks, t)
			fmt.Println("tks len: ", len(tks))
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
	tk := entity.Task{}

	for _, t := range r.tasks {
		if id == t.Id {
			tk = t
			break
		}
	}

	return tk, nil
}

func (r *TaskRepositoryInMemory) Save(t *entity.Task) error {
	r.tasks = append(r.tasks, *t)

	return nil
}

func (r *TaskRepositoryInMemory) Delete(id uuid.UUID) error {
	for _, t := range r.tasks {
		if id == t.Id {
			t.DeletedAt = time.Now()
		}
	}

	return nil
}

func (r *TaskRepositoryInMemory) Update(t *entity.Task) error {
	for _, rt := range r.tasks {
		if rt.Id == t.Id {
			rt.Title = t.Title
			rt.Describe = t.Describe
		}
	}

	return nil
}
