package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type SaveTask struct {
	TaskRepository repository.TaskRepository
}

func (s *SaveTask) Execute(list, title, description, userId string) (entity.Task, *e.Error) {
	appErr := e.New().SetLayer(e.Application)
	pTuid, er := uuid.Parse(userId)
	if er != nil {
		return entity.Task{}, appErr.CustomError(e.InvalidId)
	}

	if list == "" || title == "" || description == "" {
		return entity.Task{}, appErr.CustomError(e.EmptyField)
	}

	t := entity.NewTask(uuid.New(), list, title, description, pTuid)

	return s.TaskRepository.Save(t)
}
