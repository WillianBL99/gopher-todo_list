package usecase

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/server"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func TestSaveUser(t *testing.T) {
	t.Run("Should save user", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{UserRepository: &ur}

		u := entity.NewUser(uuid.New(), "John Doe", "email@email.com", "123456")
		err := su.Execute(u.Name, u.Email, u.Password)
		assert.Nil(t, err)

		fu, err := ur.GetByEmail(u.Email)
		assert.Nil(t, err)
		assert.Equal(t, u.Name, fu.Name)
		assert.Equal(t, u.Email, fu.Email)
	})

	t.Run("Should save user and generate uuid", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{UserRepository: &ur}

		u := entity.NewUser(uuid.Nil, "John Doe", "john@john.doe", "123456")

		su.Execute(u.Name, u.Email, u.Password)

		fu, err := ur.GetByEmail(u.Email)
		assert.Nil(t, err)
		assert.NotEqual(t, uuid.Nil, fu.Id)
	})

	t.Run("Should return error if email already exists", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{UserRepository: &ur}
		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", "123456")

		ur.Save(u)

		err := su.Execute(u.Name, u.Email, u.Password)
		assert.NotNil(t, err)
		assert.Equal(t, e.EmailAlreadyExists.Title, err.Title)
	})

	t.Run("Password should be encripted", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		su := SaveUser{UserRepository: &ur}
		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", "123456")

		su.Execute(u.Name, u.Email, u.Password)

		fu, _ := ur.GetByEmail(u.Email)

		bc := server.NewBcryptService()
		if !bc.Compare(fu.Password, u.Password) {
			t.Error("Password not encrypted")
		}
	})
}
