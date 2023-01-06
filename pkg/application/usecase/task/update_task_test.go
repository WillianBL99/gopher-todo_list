package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestUpdateTask_Execute(t *testing.T) {
	t.Run("Should update task title, describe and due date", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{Repository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "Title", "Describe", user_id)
		ntk := entity.NewTask(tk.Id, "New Title", "New Describe", user_id)

		rp.Save(tk)

		err := ut.Execute(tk.Id.String(), ntk.Title, ntk.Describe)

		if err != nil {
			t.Errorf("Error to update task: %s", err.Error())
		}

		if tks, _ := rp.GetAll(user_id); len(tks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tks))
		}

		ftk, _ := rp.GetById(tk.Id)

		if ftk.Title != ntk.Title || ftk.Describe != ntk.Describe {
			t.Errorf("Expected: {%s, %s}, got: {%s, %s}",
				ntk.Title, ntk.Describe,
				ftk.Title, ftk.Describe,
			)
		}
	})
}
