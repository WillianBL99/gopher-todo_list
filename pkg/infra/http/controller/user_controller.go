package controller

import (
	"encoding/json"
	"net/http"

	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	usecase "github.com/willianbl99/todo-list_api/pkg/application/usecase/user"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/dto"
	"github.com/willianbl99/todo-list_api/pkg/server"
)

type UserController struct {
	Providers struct {
		SaveUser usecase.SaveUser
		GetUserByEmailPassword usecase.GetUserByEmailPassword
	}
	Repository repository.UserRepository	
}

func NewUserController(r repository.UserRepository) *UserController {
	uc := &UserController{}
	uc.Providers.SaveUser = usecase.SaveUser{r}
	uc.Providers.GetUserByEmailPassword = usecase.GetUserByEmailPassword{r}
	
	return uc
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var ru dto.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&ru); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := ru.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := uc.Providers.SaveUser.Execute(ru.Name, ru.Email, ru.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User saved"))
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var lu dto.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&lu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := lu.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := uc.Providers.GetUserByEmailPassword.Execute(lu.Email, lu.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt := server.NewJwtService()
	token, err := jwt.GenerateToken(u.Id.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNetworkAuthenticationRequired)
		return
	}

	res := make(map[string]string)
	res["token"] = token
	resp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}