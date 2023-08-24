package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetTasksByList struct {
	TaskRepository repository.TaskRepository
}

func (g *GetTasksByList) Execute(listName string, uid string) ([]entity.Task, *e.Error) {
	ptuid, er := uuid.Parse(uid)
	if er != nil {
		return nil, e.New().AppErr(e.InvalidId, er.Error())
	}

	if listName == "" {
		return nil, e.New().AppErr(e.EmptyField)
	}

	tks, err := g.TaskRepository.GetByList(ptuid, listName)
	if err != nil {
		return nil, err
	}

	return tks, nil
}
