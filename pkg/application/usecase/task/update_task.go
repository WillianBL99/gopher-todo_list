package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type UpdateTask struct {
	Repository repository.TaskRepository
}

func (u *UpdateTask) Execute(id string, title string, describe string, dueDate time.Time, userId string) error {
	var err error

	tid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("Error to parse taskId: %v", err.Error())
	}

	uid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("Error to parse userId: %v", err.Error())
	}

	t := entity.NewTask(tid, title, describe, dueDate, uid)
	
	err = u.Repository.Update(t)
	if err != nil {
		return fmt.Errorf("Error to update task: %v", err.Error())
	}

	return err
}