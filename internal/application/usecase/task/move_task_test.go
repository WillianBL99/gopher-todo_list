package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestMoveTask(t *testing.T) {
	t.Run("Should move task to done", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		tk := entity.NewTask(uuid.New(), "List 1", "Test", "Description", uuid.New())
		tr.Save(tk)
		err := mt.Execute(tk.Id.String(), "done")
		assert.Nil(t, err)

		ftk, err := tr.GetById(tk.Id)
		assert.Nil(t, err)
		assert.Equal(t, entity.Done, ftk.Status)
	})

	t.Run("Should move task to undone", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		tk := entity.NewTask(uuid.New(), "List 1", "Test", "Description", uuid.New())
		tk.Status = entity.Done
		tr.Save(tk)
		err := mt.Execute(tk.Id.String(), "undone")
		assert.Nil(t, err)

		ftk, err := tr.GetById(tk.Id)
		assert.Nil(t, err)
		assert.Equal(t, entity.Undone, ftk.Status)
	})

	t.Run("Should return conflict error if task is already in the status", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		type testCase struct {
			status entity.Status
		}
		tc := []testCase{
			{status: entity.Done},
			{status: entity.Undone},
		}

		for _, c := range tc {
			tk := entity.NewTask(uuid.New(), "List 1", "Test", "Description", uuid.New())
			tk.Status = c.status
			tr.Save(tk)
			err := mt.Execute(tk.Id.String(), string(c.status))
			assert.NotNil(t, err)
			assert.Equal(t, e.Conflict.Title, err.Title)
		}
	})

	t.Run("Should return error if task not found", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		err := mt.Execute(uuid.New().String(), "done")
		assert.NotNil(t, err)
		assert.Equal(t, e.TaskNotFound.Title, err.Title)
	})

	t.Run("Should return error if invalid user id", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		err := mt.Execute("invalid-uuid", "done")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should return error if invalid status", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		tk := entity.NewTask(uuid.New(), "List 1", "Test", "Description", uuid.New())
		tr.Save(tk)
		err := mt.Execute(tk.Id.String(), "invalid")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidStatus.Title, err.Title)
	})
}
