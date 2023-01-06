package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestDeleteTask(t *testing.T) {
	t.Run("Should delete a task", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		d := DeleteTask{TaskRepository: &tr}

		uid := uuid.New()
		tk := entity.NewTask(uuid.New(), "Task 1", "Description 1", uid)
		tr.Save(tk)

		err := d.Execute(tk.Id.String())
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		tks, _ := tr.GetAll(uid)
		if len(tks) != 0 {
			t.Errorf("Expected 0 tasks, got %d", len(tks))
		}
	})
}
