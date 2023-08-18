package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestUpdateTask_Execute(t *testing.T) {
	t.Run("Should update task title, description and due date", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "Title", "Description", user_id)
		ntk := entity.NewTask(tk.Id, "New Title", "New Description", user_id)

		rp.Save(tk)

		err := ut.Execute(tk.Id.String(), ntk.Title, ntk.Description)

		if err != nil {
			t.Errorf("Error to update task: %s", err.Error())
		}

		if tks, _ := rp.GetAll(user_id); len(tks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tks))
		}

		ftk, _ := rp.GetById(tk.Id)

		if ftk.Title != ntk.Title || ftk.Description != ntk.Description {
			t.Errorf("Expected: {%s, %s}, got: {%s, %s}",
				ntk.Title, ntk.Description,
				ftk.Title, ftk.Description,
			)
		}
	})
}
