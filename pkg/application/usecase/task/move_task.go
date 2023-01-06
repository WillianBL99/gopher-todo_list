package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
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
		return fmt.Errorf("Invalid status")
	}

	task, err := mt.TaskRepository.GetById(ptid)
	if err != nil {
		return err
	}

	if task.Status == pst {
		return fmt.Errorf("Task is already %v", pst)
	}

	task.Status = pst

	if err := mt.TaskRepository.Update(&task); err != nil {
		return fmt.Errorf("Error chaging task status: %w", err)
	}

	return nil
}