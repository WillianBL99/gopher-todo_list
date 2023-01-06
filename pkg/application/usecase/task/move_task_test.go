package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestMoveTask(t *testing.T) {
	t.Run("Should move task to done", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		mt := MoveTask{TaskRepository: &tr}

		tk := entity.NewTask(uuid.New(), "Test", "Description", uuid.New())
		tr.Save(tk)

		if err := mt.Execute(tk.Id.String(), "done"); err != nil {
			t.Errorf("Error moving task to done: %v", err)
		}

		ftk, _ := tr.GetById(tk.Id)
		if ftk.Status != entity.Done {
			t.Errorf("Task status should be done but is %v", ftk.Status)
		}
	})
}
