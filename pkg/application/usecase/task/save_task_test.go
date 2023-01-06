package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestSaveTask(t *testing.T) {
	t.Run("Should save task", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{Repository: &rp}
		user_id := uuid.New()

		err := st.Execute("Title", "Description", user_id.String())

		if err != nil {
			t.Errorf("Error to save task: %s", err.Error())
		}

		if l, _ := rp.GetAll(user_id); len(l) != 1 {
			t.Errorf("Expected 1 task, got %d", len(l))
		}
	})

	t.Run("All tasks should be undone", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{Repository: &rp}
		user_id := uuid.New()
		
		st.Execute("Title1", "Description1", user_id.String())
		st.Execute("Title2", "Description2", user_id.String())
		st.Execute("Title3", "Description3", user_id.String())

		if l, _ := rp.GetByStatus(user_id, entity.Undone); len(l) != 3 {
			t.Errorf("Expected 3 undone tasks, got %d", len(l))
		}
	})

	t.Run("Should save task and generate uuid", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{Repository: &rp}
		user_id := uuid.New()

		err := st.Execute("Title", "Description", user_id.String())

		if err != nil {
			t.Errorf("Error to save task: %s", err.Error())
		}

		if l, _ := rp.GetAll(user_id); len(l) != 1 {
			t.Errorf("Expected 1 task, got %d", len(l))
		}

		if l, _ := rp.GetAll(user_id); l[0].Id == uuid.Nil {
			t.Errorf("Expected uuid nil, got %s", l[0].Id)
		}
	})

	t.Run("Should not save task with invalid user id", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{Repository: &rp}

		err := st.Execute("Title", "Description", "invalid")

		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
