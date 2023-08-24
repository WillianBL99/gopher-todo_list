package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestGetTasksByStatus(t *testing.T) {
	t.Run("Should return all tasks undone", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}
		uid := uuid.New()

		dtk := entity.NewTask(uuid.New(), "List 1", "Task 1", "Description 1", uid)
		dtk.Status = entity.Done

		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 2", "Description 2", uid))
		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 3", "Description 3", uid))
		tr.Save(dtk)

		tasks, err := gt.Execute(uid.String(), "undone")
		assert.Nil(t, err)
		assert.Equal(t, 2, len(tasks))
		assert.Equal(t, entity.Undone, tasks[0].Status)
	})

	t.Run("Should return all tasks done", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}
		uid := uuid.New()

		dtk := entity.NewTask(uuid.New(), "List 1", "Task 1", "Description 1", uid)
		dtk.Status = entity.Done

		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 2", "Description 2", uid))
		tr.Save(entity.NewTask(uuid.New(), "List 1", "Task 3", "Description 3", uid))
		tr.Save(dtk)

		tasks, err := gt.Execute(uid.String(), "done")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(tasks))
		assert.Equal(t, entity.Done, tasks[0].Status)
	})

	t.Run("Should return error if invalid user id", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}

		_, err := gt.Execute("invalid-uuid", "undone")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should return error if invalid status", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}
		uid := uuid.New()

		dtk := entity.NewTask(uuid.New(), "List 1", "Task 1", "Description 1", uid)
		tr.Save(dtk)

		_, err := gt.Execute(uid.String(), "invalid")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidStatus.Title, err.Title)
	})
}
