package usecases

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/uilianlago/API-To-Do-List/infra/db/inmemory/repositories"
)

func TestSaveTask(t *testing.T) {
	t.Run("Should save task", func(t *testing.T) {
		rp := repositories.InMemoryToDoListRepository{}
		st := SaveTask{Repository: &rp}

		err := st.Execute("Title", "Describe", time.Now())

		if err != nil {
			t.Errorf("Error to save task: %s", err.Error())
		}

		if l, _ := rp.GetAllDone(); len(l) != 1 {
			t.Errorf("Expected 1 task, got %d", len(l))
		}
	})

	t.Run("Should generate a uuid", func(t *testing.T) {
		rp := repositories.InMemoryToDoListRepository{}
		st := SaveTask{Repository: &rp}

		err := st.Execute("Title", "Describe", time.Now())

		if err != nil {
			t.Errorf("Error to save task: %s", err.Error())
		}

		if l, _ := rp.GetAllDone(); len(l) != 1 {
			t.Errorf("Expected 1 task, got %d", len(l))
		}

		if l, _ := rp.GetAllDone(); l[0].Id != uuid.Nil {
			t.Errorf("Expected uuid nil, got %s", l[0].Id)
		}
	})
}
