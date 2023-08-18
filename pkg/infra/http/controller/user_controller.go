package controller

import (
	"fmt"
	"net/http"

	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	usecase "github.com/willianbl99/todo-list_api/pkg/application/usecase/user"
	"github.com/willianbl99/todo-list_api/pkg/herr"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/dto"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type UserController struct {
	Providers struct {
		SaveUser               usecase.SaveUser
		GetUserByEmailPassword usecase.GetUserByEmailPassword
	}
	Repository repository.UserRepository
}

func NewUserController(r repository.UserRepository) *UserController {
	uc := &UserController{}
	uc.Providers.SaveUser = usecase.SaveUser{UserRepository: r}
	uc.Providers.GetUserByEmailPassword = usecase.GetUserByEmailPassword{UserRepository: r}
	return uc
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var ru dto.RegisterUserRequest
	if err := dto.ToDTO(r, &ru); err != nil {
		herr.BadBodyRequest(w, err)
		return
	}
	err := uc.Providers.SaveUser.Execute(ru.Name, ru.Email, ru.Password)
	if err != nil {
		herr.AppToHttp(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "User created successfully"}`)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var lu dto.LoginUserRequest
	if err := dto.ToDTO(r, &lu); err != nil {
		herr.BadBodyRequest(w, err)
		return
	}
	u, err := uc.Providers.GetUserByEmailPassword.Execute(lu.Email, lu.Password)
	if err != nil {
		herr.AppToHttp(w, err)
		return
	}
	jwt := server.NewJwtService()
	token, err := jwt.GenerateToken(u.Id.String())
	if err != nil {
		herr.NewHttp().InternalServerError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}
