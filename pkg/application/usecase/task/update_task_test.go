package usecase

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestUpdateTask_Execute(t *testing.T) {
	t.Run("Should update task title, describe and due date", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{Repository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "Title", "Describe", time.Now(), user_id)
		ntk := entity.NewTask(uuid.New(), "New Title", "New Describe", time.Now().AddDate(0,0,1), user_id)

		rp.Save(tk)

		err := ut.Execute(ntk.Id.String(), ntk.Title, ntk.Describe, ntk.DueDate, user_id.String())

		if err != nil {
			t.Errorf("Error to update task: %s", err.Error())
		}

		tks, _ := rp.GetAll(user_id)

		if len(tks) != 1 {
			t.Errorf("Expected 1 task, got %d", len(tks))
		}

		if tks[0].Title != ntk.Title {
			t.Errorf("Expected title %s, got %s", ntk.Title, tks[0].Title)
		}

		if tks[0].Describe != ntk.Describe {
			t.Errorf("Expected describe %s, got %s", ntk.Describe, tks[0].Describe)
		}

		if tks[0].DueDate != ntk.DueDate {
			t.Errorf("Expected due date %s, got %s", ntk.DueDate, tks[0].DueDate)
		}
	})

}