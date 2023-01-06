package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/controller"
)

func UserRouter(r chi.Router, ur repository.UserRepository) {
	uc := controller.NewUserController(ur)

	r.Post("/sign-up", uc.Register)
	r.Post("/sign-in", uc.Login)
}
