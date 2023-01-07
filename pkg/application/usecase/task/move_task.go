package usecase

import (
	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type MoveTask struct {
	TaskRepository repository.TaskRepository
}

func (mt *MoveTask) Execute(tid string, st string) error {
	ptid, err := uuid.Parse(tid)
	if err != nil {
		return err
	}

	pst := entity.Status(st)
	if pst != entity.Done && pst != entity.Undone {
		return herr.NewApp().InvalidTaskStatus
	}

	task, err := mt.TaskRepository.GetById(ptid)
	if err != nil {
		return err
	}

	if task.Status == pst {
		return herr.NewApp().Conflict
	}

	task.Status = pst

	if err := mt.TaskRepository.Update(&task); err != nil {
		return err
	}

	return nil
}