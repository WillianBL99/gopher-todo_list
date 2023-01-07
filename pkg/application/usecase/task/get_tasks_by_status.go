package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetTasksByStatus struct {
	TaskRepository repository.TaskRepository
}

func (g *GetTasksByStatus) Execute(uid string, st string) ([]entity.Task, error) {
	puid, err := uuid.Parse(uid)
	if err != nil {
		return nil, herr.NewApp().InvalidUserId
	}

	est := entity.Status(st)
	if est != entity.Done && est != entity.Undone {
		return nil, herr.NewApp().InvalidTaskStatus
	}

	tks, err := g.TaskRepository.GetByStatus(puid, est)
	if err != nil {
		return nil, err
	}

	return tks, nil
}