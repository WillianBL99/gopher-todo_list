package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type GetTasksByStatus struct {
	TaskRepository repository.TaskRepository
}

func (g *GetTasksByStatus) Execute(uid string, st string) ([]entity.Task, error) {
	puid, err := uuid.Parse(uid)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %v", err)
	}

	est := entity.Status(st)
	if est != entity.Done && est != entity.Undone {
		return nil, fmt.Errorf("invalid status: %v", est)
	}

	tks, err := g.TaskRepository.GetByStatus(puid, est)
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %v", err)
	}

	return tks, nil
}