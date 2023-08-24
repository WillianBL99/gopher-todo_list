package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestGetTaskById(t *testing.T) {
	t.Run("Should return a task", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTaskById{TaskRepository: &tr}
		st := SaveTask{TaskRepository: &tr}
		tk, err := st.Execute("List 1", "Title 1", "Description 1", uuid.New().String())
		assert.Nil(t, err)
		gtk, err := gt.Execute(tk.Id.String())
		assert.Nil(t, err)
		assert.Equal(t, tk.Id.String(), gtk.Id.String())
		assert.Equal(t, tk.List, gtk.List)
	})

	t.Run("Should return an error when id is invalid", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTaskById{TaskRepository: &tr}
		_, err := gt.Execute("invalid-id")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should return an error when task not found", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTaskById{TaskRepository: &tr}
		_, err := gt.Execute(uuid.New().String())
		assert.NotNil(t, err)
		assert.Equal(t, e.TaskNotFound.Title, err.Title)
	})
}
