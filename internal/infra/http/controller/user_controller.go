package controller

import (
	"fmt"
	"net/http"

	"github.com/willianbl99/todo-list_api/internal/application/repository"
	usecase "github.com/willianbl99/todo-list_api/internal/application/usecase/user"
	"github.com/willianbl99/todo-list_api/internal/infra/http/dto"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
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
		e.New().InfraHttpErr(e.BadBodyRequest, err.Error()).ToHttp(w)
		return
	}
	err := uc.Providers.SaveUser.Execute(ru.Name, ru.Email, ru.Password)
	if err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"message": "User created successfully"}`)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var lu dto.LoginUserRequest
	if err := dto.ToDTO(r, &lu); err != nil {
		e.New().InfraHttpErr(e.BadBodyRequest, err.Error()).ToHttp(w)
		return
	}
	u, err := uc.Providers.GetUserByEmailPassword.Execute(lu.Email, lu.Password)
	if err != nil {
		e.New().InfraHttpErrRec(err).ToHttp(w)
		return
	}
	jwt := server.NewJwtService()
	token, er := jwt.GenerateToken(u.Id.String())
	if er != nil {
		e.New().InfraHttpErr(e.InternalServerError, er.Error()).ToHttp(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}
