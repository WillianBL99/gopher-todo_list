package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestSaveTask(t *testing.T) {
	t.Run("Should save task", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &tr}
		user_id := uuid.New()

		tk, err := st.Execute("List 1", "Title", "Description", user_id.String())
		assert.Nil(t, err)
		assert.Equal(t, true, tk.Id != uuid.Nil)
		assert.Equal(t, "List 1", tk.List)
		tks, err := tr.GetAll(user_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(tks))
	})

	t.Run("All tasks should be undone", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &rp}
		user_id := uuid.New()

		st.Execute("List 1", "Title1", "Description1", user_id.String())
		st.Execute("List 1", "Title2", "Description2", user_id.String())
		st.Execute("List 1", "Title3", "Description3", user_id.String())

		l, err := rp.GetByStatus(user_id, entity.Undone)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(l))
		assert.Equal(t, entity.Undone, l[0].Status)
		assert.Equal(t, entity.Undone, l[1].Status)
		assert.Equal(t, entity.Undone, l[2].Status)
	})

	t.Run("Should save task and generate uuid", func(t *testing.T) {
		tr := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &tr}
		user_id := uuid.New()

		tk, err := st.Execute("List 1", "Title", "Description", user_id.String())
		assert.Nil(t, err)
		assert.Equal(t, true, tk.Id != uuid.Nil)
		assert.Equal(t, "List 1", tk.List)
		tks, err := tr.GetAll(user_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(tks))
		assert.NotEqual(t, uuid.Nil, tks[0].Id)
	})

	t.Run("Should not save task with invalid user id", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &rp}

		_, err := st.Execute("List 1", "Title", "Description", "invalid")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should not save task with empty title", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &rp}
		user_id := uuid.New()

		_, err := st.Execute("List 1", "", "Description", user_id.String())
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})

	t.Run("Should not save task with empty description", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &rp}
		user_id := uuid.New()

		_, err := st.Execute("List 1", "Title", "", user_id.String())
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})

	t.Run("Should not save task with empty list", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		st := SaveTask{TaskRepository: &rp}
		user_id := uuid.New()

		_, err := st.Execute("", "Title", "Description", user_id.String())
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})
}
