package usecase

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type SaveTask struct {
	Repository repository.TaskRepository
}

func (s *SaveTask) Execute(title string, describe string, dueDate time.Time, userId string) error {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("Error to parse userId: %v", err.Error())
	}

	t := entity.NewTask(uuid.New(), title, describe, dueDate, uid)
	
	if err := s.Repository.Save(t); err != nil {
		return fmt.Errorf("Error to save task: %v", err.Error())
	}

	return nil
}
