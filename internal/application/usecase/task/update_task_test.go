package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestUpdateTask_Execute(t *testing.T) {
	t.Run("Should update task title, description and due date", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "List 1", "Title", "Description", user_id)
		ntk := entity.NewTask(tk.Id, "List 1", "New Title", "New Description", user_id)

		rp.Save(tk)
		err := ut.Execute(tk.Id.String(), ntk.List, ntk.Title, ntk.Description)
		assert.Nil(t, err)
		tks, err := rp.GetAll(user_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(tks))
		assert.Equal(t, false, tks[0].UpdatedAt.Time.IsZero())

		currtk, err := rp.GetById(tk.Id)
		assert.Nil(t, err)
		assert.Equal(t, ntk.Title, currtk.Title)
		assert.Equal(t, ntk.Description, currtk.Description)
		assert.Equal(t, false, currtk.UpdatedAt.Time.IsZero())

		if currtk.Title != ntk.Title || currtk.Description != ntk.Description {
			t.Errorf("Expected: {%s, %s}, got: {%s, %s}",
				ntk.Title, ntk.Description,
				currtk.Title, currtk.Description,
			)
		}
	})

	t.Run("Should update task title, description and due date", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "List 1", "Title", "Description", user_id)
		ntk := entity.NewTask(tk.Id, "List 1", "New Title", "New Description", user_id)

		rp.Save(tk)
		err := ut.Execute(tk.Id.String(), ntk.List, ntk.Title, ntk.Description)
		assert.Nil(t, err)
		tks, err := rp.GetAll(user_id)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(tks))

		currtk, err := rp.GetById(tk.Id)
		assert.Nil(t, err)
		assert.Equal(t, ntk.Title, currtk.Title)
		assert.Equal(t, ntk.Description, currtk.Description)
		assert.Equal(t, false, currtk.UpdatedAt.Time.IsZero())

		if currtk.Title != ntk.Title || currtk.Description != ntk.Description {
			t.Errorf("Expected: {%s, %s}, got: {%s, %s}",
				ntk.Title, ntk.Description,
				currtk.Title, currtk.Description,
			)
		}
	})

	t.Run("Should not update task with invalid user id", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}

		err := ut.Execute("Invalid", "List 1", "Title", "Description")
		assert.NotNil(t, err)
		assert.Equal(t, e.InvalidId.Title, err.Title)
	})

	t.Run("Should not update task with empty title", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "List 1", "Title", "Description", user_id)

		rp.Save(tk)
		err := ut.Execute(tk.Id.String(), "List 1", "", "Description")
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})

	t.Run("Should not update task with empty description", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "List 1", "Title", "Description", user_id)

		rp.Save(tk)
		err := ut.Execute(tk.Id.String(), "List 1", "Title", "")
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})

	t.Run("Should not update task with empty list", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}
		user_id := uuid.New()
		tk := entity.NewTask(uuid.New(), "List 1", "Title", "Description", user_id)

		rp.Save(tk)
		err := ut.Execute(user_id.String(), "", "Title", "Description")
		assert.NotNil(t, err)
		assert.Equal(t, e.EmptyField.Title, err.Title)
	})

	t.Run("Should not update task if task not found", func(t *testing.T) {
		rp := inmemory.TaskRepositoryInMemory{}
		ut := UpdateTask{TaskRepository: &rp}

		err := ut.Execute(uuid.New().String(), "List 1", "Title", "Description")
		assert.NotNil(t, err)
		assert.Equal(t, e.TaskNotFound.Title, err.Title)
	})
}
