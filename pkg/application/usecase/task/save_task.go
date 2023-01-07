package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type SaveTask struct {
	Repository repository.TaskRepository
}

func (s *SaveTask) Execute(title string, description string, userId string) error {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return herr.NewApp().BadRequest
	}

	t := entity.NewTask(uuid.New(), title, description, uid)

	if err := s.Repository.Save(t); err != nil {
		return err
	}

	return nil
}
