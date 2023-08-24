package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type DeleteTask struct {
	TaskRepository repository.TaskRepository
}

func (d *DeleteTask) Execute(id string) *e.Error {
	appErr := e.New().SetLayer(e.Application)
	pUid, er := uuid.Parse(id)
	if er != nil {
		return appErr.CustomError(e.InvalidId)
	}
	_, err := d.TaskRepository.GetById(pUid)
	if err != nil {
		return err
	}
	return d.TaskRepository.Delete(pUid)
}
