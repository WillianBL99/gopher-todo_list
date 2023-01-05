package controller

import (
	"encoding/json"
	"net/http"

	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	usecase "github.com/willianbl99/todo-list_api/pkg/application/usecase/user"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/dto"
)

type UserController struct {
	Providers struct {
		SaveUser usecase.SaveUser
	}
	Repository repository.UserRepository	
}

func NewUserController(r repository.UserRepository) *UserController {
	uc := &UserController{}
	uc.Providers.SaveUser = usecase.SaveUser{r}
	
	return uc
}

func (uc *UserController) Save(w http.ResponseWriter, r *http.Request) {
	var ur dto.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&ur); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := ur.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := uc.Providers.SaveUser.Execute(ur.Name, ur.Email, ur.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User saved"))
}