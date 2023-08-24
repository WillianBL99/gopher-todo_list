package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestDeleteTask(t *testing.T) {
	t.Run("Should delete a task", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		d := DeleteTask{TaskRepository: &tr}

		uid := uuid.New()
		tk := entity.NewTask(uuid.New(), "Some title", "Task 1", "Description 1", uid)
		tr.Save(tk)

		err := d.Execute(tk.Id.String())
		assert.Nil(t, err)
		tks, _ := tr.GetAll(uid)
		assert.Equal(t, 0, len(tks))
	})

	t.Run("Should return error when task not found", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		d := DeleteTask{TaskRepository: &tr}

		err := d.Execute(uuid.New().String())
		assert.NotNil(t, err)
		assert.Equal(t, e.TaskNotFound.Title, err.Title)
	})
}
