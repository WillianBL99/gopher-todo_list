package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestGetTasksByStatus(t *testing.T) {
	t.Run("Should return all tasks undone", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}
		uid := uuid.New()

		dtk := entity.NewTask(uuid.New(), "Task 1", "Description 1", uid)
		dtk.Status = entity.Done

		tr.Save(entity.NewTask(uuid.New(), "Task 2", "Description 2", uid))
		tr.Save(entity.NewTask(uuid.New(), "Task 3", "Description 3", uid))
		tr.Save(dtk)

		tasks, err := gt.Execute(uid.String(), "undone")

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if len(tasks) != 2 {
			t.Errorf("Expected 2 tasks, got %d", len(tasks))
		}
	})

	t.Run("Should return all tasks done", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}
		uid := uuid.New()

		dtk := entity.NewTask(uuid.New(), "Task 1", "Description 1", uid)
		dtk.Status = entity.Done

		tr.Save(entity.NewTask(uuid.New(), "Task 2", "Description 2", uid))
		tr.Save(entity.NewTask(uuid.New(), "Task 3", "Description 3", uid))
		tr.Save(dtk)

		tasks, err := gt.Execute(uid.String(), "done")

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if len(tasks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tasks))
		}
	})

	t.Run("Should return error if invalid status", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		gt := GetTasksByStatus{TaskRepository: &tr}
		uid := uuid.New()

		dtk := entity.NewTask(uuid.New(), "Task 1", "Description 1", uid)
		tr.Save(dtk)

		_, err := gt.Execute(uid.String(), "invalid")
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
