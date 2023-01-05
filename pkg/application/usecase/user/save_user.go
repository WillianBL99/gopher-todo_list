package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
)

type SaveUser struct {
	Repository repository.UserRepository
}

func (s *SaveUser) Execute(name string, email string, password string) error {
	if _, err := s.Repository.GetByEmail(email); err == nil {
		return fmt.Errorf("Email already exists")
	}

	u := entity.NewUser(uuid.New(), name, email, password)

	if err := s.Repository.Save(u); err != nil {
		return fmt.Errorf("Error to save user: %v", err.Error())
	}

	return nil
}