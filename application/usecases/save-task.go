package usecases

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uilianlago/API-To-Do-List/application/entities"
	"github.com/uilianlago/API-To-Do-List/application/repositories"
)

type SaveTask struct {
	Repository repositories.ToDoListRepository
}

func (s *SaveTask) Execute(title string, describe string, dueDate time.Time) error {
	t := entities.NewTask(uuid.New(), title, describe, dueDate)
	err := s.Repository.Save(t)

	if err != nil {
		return fmt.Errorf("Error to save task: %v", err.Error())
	}

	return err
}
