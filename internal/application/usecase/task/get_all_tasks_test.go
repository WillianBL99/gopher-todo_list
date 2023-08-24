package usecase

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestGetAllTasks(t *testing.T) {
	t.Run("Should return all tasks", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TaskRepository: &tr}

		uid := uuid.New()
		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 1", "Description 1", uid))
		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 2", "Description 2", uid))
		tr.Save(entity.NewTask(uuid.New(), "List 2", "Task 3", "Description 3", uid))

		tasks, err := g.Execute(uid.String())
		assert.Nil(t, err)
		assert.Equal(t, 3, len(tasks))
	})

	t.Run("Should return an error if the user id is invalid", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TaskRepository: &tr}

		_, err := g.Execute("invalid-uuid")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should not get any task if the user has no tasks", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TaskRepository: &tr}

		uid := uuid.New()
		tasks, err := g.Execute(uid.String())
		assert.Nil(t, err)
		assert.Equal(t, 0, len(tasks))
	})

	t.Run("Should not get deleted tasks", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TaskRepository: &tr}

		uid := uuid.New()

		dltk := entity.NewTask(uuid.New(), "List 1", "Deleted Task", "Description", uid)
		dltk.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}

		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 1", "Description 1", uid))
		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 2", "Description 2", uid))
		tr.Save(dltk)

		tks, err := g.Execute(uid.String())
		assert.Nil(t, err)
		assert.Equal(t, 2, len(tks))
	})
}
