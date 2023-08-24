package usecase

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestGetTaskByList(t *testing.T) {
	t.Run("Should return all tasks from a list", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetTasksByList{TaskRepository: &tr}
		uid := uuid.New()
		for i := 0; i < 5; i++ {
			title := fmt.Sprint("Task ", i+1)
			desc := fmt.Sprint("Description ", i+1)
			tr.Save(entity.NewTask(uuid.New(), "List 1", title, desc, uid))
		}
		tks, err := g.Execute("List 1", uid.String())
		assert.Nil(t, err)
		assert.Equal(t, 5, len(tks))
		assert.Equal(t, "Task 1", tks[0].Title)
	})

	t.Run("Should return an error if the user id is invalid", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetTasksByList{TaskRepository: &tr}
		_, err := g.Execute("List 1", "invalid-uuid")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should return an error if the list name is empty", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetTasksByList{TaskRepository: &tr}
		uid := uuid.New()
		_, err := g.Execute("", uid.String())
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})
}
