package usecase

import (
	"fmt"

	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type GetUserByEmailPassword struct {
	Repository repository.UserRepository
}

func (gu *GetUserByEmailPassword) Execute(e string, p string) (entity.User, error) {
	u, err := gu.Repository.GetByEmail(e)
	if err != nil {
		return u, fmt.Errorf("Error to get user: %v", err.Error())
	}

	bc := server.NewBcryptService()
	if !bc.Compare(u.Password, p) {
		hashed, _ := bc.Encrypt(p)
		fmt.Printf("Compare: %v\n", bc.Compare(hashed, "123456"))
		fmt.Printf("Compare: %v\n", bc.Compare(u.Password, p))
		fmt.Printf("Hash: %v - Pass: %v\n", u.Password, p)
		return u, fmt.Errorf("Email or password invalid")
	}

	return u, nil
}
