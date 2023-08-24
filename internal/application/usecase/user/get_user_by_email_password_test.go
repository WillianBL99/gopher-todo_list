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

func TestGetUserByEmailPassword(t *testing.T) {
	t.Run("Should return user", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		gu := GetUserByEmailPassword{UserRepository: &ur}

		p := "123456"
		bc := server.NewBcryptService()
		hash, _ := bc.Encrypt(p)
		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", hash)

		ur.Save(u)

		fu, err := gu.Execute(u.Email, p)
		assert.Nil(t, err)
		assert.Equal(t, u.Id, fu.Id)
		assert.Equal(t, u.Name, fu.Name)
		assert.Equal(t, u.Email, fu.Email)
	})

	t.Run("Should return error if password is invalid", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		gu := GetUserByEmailPassword{UserRepository: &ur}

		p := "123456"
		bc := server.NewBcryptService()
		hash, _ := bc.Encrypt(p)
		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", hash)

		_, err := gu.Execute(u.Email, "invalid_password")
		assert.NotNil(t, err)
		assert.Equal(t, e.EmailOrPasswordInvalid.Title, err.Title)
	})

	t.Run("Should return error if email is invalid", func(t *testing.T) {
		ur := inmemory.UserRepositoryInMemory{}
		gu := GetUserByEmailPassword{UserRepository: &ur}

		p := "123456"
		bc := server.NewBcryptService()
		hash, _ := bc.Encrypt(p)

		u := entity.NewUser(uuid.New(), "John Doe", "john@john.doe", hash)
		ur.Save(u)

		_, err := gu.Execute("invalid_email", p)
		assert.NotNil(t, err)
		assert.Equal(t, e.EmailOrPasswordInvalid.Title, err.Title)
	})
}
