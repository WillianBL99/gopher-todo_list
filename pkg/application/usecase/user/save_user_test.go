package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func SaveUserTest(t *testing.T) {
	t.Run("Should save user", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{Repository: &ur}

		u := entity.NewUser(uuid.New(), "John Doe", "email@email.com", "123456")

		if err := su.Execute(u.Name, u.Email, u.Password); err != nil {
			t.Errorf("Error to save user: %v", err.Error())
		}

		fu, err := ur.GetByEmail(u.Email)

		if err != nil {
			t.Errorf("Error to get user: %v", err.Error())
		}

		if fu.Name != u.Name || fu.Email != u.Email || fu.Password != u.Password {
			t.Errorf("Expected user %v, got %v", u, fu)
		}
	})

	t.Run("Should save user and generate uuid", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{Repository: &ur}

		u := entity.NewUser(uuid.Nil, "John Doe", "john@john.doe", "123456")

		su.Execute(u.Name, u.Email, u.Password)

		fu, _ := ur.GetByEmail(u.Email)

		if fu.Id == uuid.Nil {
			t.Errorf("Expected user with uuid, got %v", fu)
		}
	})

	t.Run("Should return error if email already exists", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{Repository: &ur}
		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", "123456")

		ur.Save(u)
		
		err := su.Execute(u.Name, u.Email, u.Password)
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})
}