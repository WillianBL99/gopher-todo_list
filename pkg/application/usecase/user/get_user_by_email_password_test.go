package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func GetUserByEmailPasswordTest(t *testing.T) {
	t.Run("Should return user", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		gu := GetUserByEmailPassword{Repository: &ur}

		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", "123456")

		ur.Save(u)

		fu, err := gu.Execute(u.Email, u.Password)
		if err != nil {
			t.Errorf("Error to get user: %v", err.Error())
		}

		if fu.Name != u.Name || fu.Email != u.Email || fu.Password != u.Password {
			t.Errorf("Expected user %v, got %v", u, fu)
		}
	})

	t.Run("Should return error if password is invalid", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		gu := GetUserByEmailPassword{Repository: &ur}

		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", "123456")
		ur.Save(u)

		_, err := gu.Execute(u.Email, "invalid_password")
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})

	t.Run("Should return error if email is invalid", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		gu := GetUserByEmailPassword{Repository: &ur}

		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", "123456")
		ur.Save(u)

		_, err := gu.Execute("invalid_email", u.Password)
		if err == nil {
			t.Errorf("Expected error, got %v", err)
		}
	})
}
