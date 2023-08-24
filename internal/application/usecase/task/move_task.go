package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	"github.com/willianbl99/todo-list_api/internal/application/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type MoveTask struct {
	TaskRepository repository.TaskRepository
}

func (mt *MoveTask) Execute(tid string, st string) *e.Error {
	appErr := e.New().SetLayer(e.Application)
	pTid, er := uuid.Parse(tid)
	if er != nil {
		return appErr.CustomError(e.InvalidId)
	}

	pst := entity.Status(st)
	if pst != entity.Done && pst != entity.Undone {
		return appErr.CustomError(e.InvalidStatus)
	}

	task, err := mt.TaskRepository.GetById(pTid)
	if err != nil {
		return err
	}

	if task.Status == pst {
		return appErr.CustomError(e.Conflict)
	}

	task.Status = pst
	return mt.TaskRepository.Update(&task)
}
