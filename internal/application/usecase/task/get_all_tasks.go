package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetAllTasks struct {
	TaskRepository repository.TaskRepository
}

func (g *GetAllTasks) Execute(uid string) ([]entity.Task, *e.Error) {
	appErr := e.New().SetLayer(e.Application)
	puid, er := uuid.Parse(uid)
	if er != nil {
		return nil, appErr.CustomError(e.InvalidId)
	}

	tasks, err := g.TaskRepository.GetAll(puid)
	if err != nil {
		return nil, appErr.CustomError(e.InternalServerError, err.ToSubErr()...)
	}

	return tasks, nil
}
