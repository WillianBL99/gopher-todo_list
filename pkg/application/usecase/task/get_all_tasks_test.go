package usecase

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestGetAllTasks(t *testing.T) {
	t.Run("Should return all tasks", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TasksRepository: &tr}

		uid := uuid.New()
		tr.Save(entity.NewTask(uuid.New(), "Task 1", "Description 1", uid))
		tr.Save(entity.NewTask(uuid.New(), "Task 2", "Description 2", uid))
		tr.Save(entity.NewTask(uuid.New(), "Task 3", "Description 3", uid))

		tasks, err := g.Execute(uid.String())
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(tasks) != 3 {
			t.Errorf("Expected 3 tasks, got %d", len(tasks))
		}
	})

	t.Run("Should return an error if the user id is invalid", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TasksRepository: &tr}

		_, err := g.Execute("invalid-uuid")
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})

	t.Run("Should not get any task if the user has no tasks", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TasksRepository: &tr}

		uid := uuid.New()
		tasks, err := g.Execute(uid.String())
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if len(tasks) != 0 {
			t.Errorf("Expected 0 tasks, got %d", len(tasks))
		}
	})

	t.Run("Should not get deleted tasks", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		g := GetAllTasks{TasksRepository: &tr}

		uid := uuid.New()
		
		dltk := entity.NewTask(uuid.New(), "Deleted Task", "Description", uid)
		dltk.DeletedAt = time.Now()

		tr.Save(entity.NewTask(uuid.New(), "Task 1", "Description 1", uid))
		tr.Save(entity.NewTask(uuid.New(), "Task 2", "Description 2", uid))
		tr.Save(dltk)

		tks,_ := g.Execute(uid.String())

		if len(tks) != 2 {
			t.Errorf("Expected 2 tasks, got %d", len(tks))
		}
	})
}