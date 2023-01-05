package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/controller"
	"github.com/willianbl99/todo-list_api/test/repositories/inmemory"
)

func UserRouter(r chi.Router) {
	ur := inmemory.UserRepositoryInMemory{}
	uc := controller.NewUserController(&ur)

	r.Post("/sign-up", uc.Register)
	r.Post("/sign-in", uc.Login)
}
