package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type GetTaskById struct {
	TaskRepository repository.TaskRepository
}

func (g *GetTaskById) Execute(id string) (entity.Task, *e.Error) {
	pUid, er := uuid.Parse(id)
	if er != nil {
		return entity.Task{}, e.New().AppErr(e.InvalidId, er.Error())
	}

	tks, err := g.TaskRepository.GetById(pUid)
	if err != nil {
		return entity.Task{}, err
	}

	return tks, nil
}
