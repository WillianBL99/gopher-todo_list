package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetTasksByStatus struct {
	TaskRepository repository.TaskRepository
}

func (g *GetTasksByStatus) Execute(uid string, st string) ([]entity.Task, *e.Error) {
	appErr := e.New().SetLayer(e.Application)
	pUid, er := uuid.Parse(uid)
	if er != nil {
		return nil, appErr.CustomError(e.InvalidId)
	}

	est := entity.Status(st)
	if est != entity.Done && est != entity.Undone {
		return nil, appErr.CustomError(e.InvalidStatus)
	}

	tks, err := g.TaskRepository.GetByStatus(pUid, est)
	if err != nil {
		return nil, err
	}

	return tks, nil
}
