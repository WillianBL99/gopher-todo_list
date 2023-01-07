package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/willianbl99/todo-list_api/pkg/infra/db"
	"github.com/willianbl99/todo-list_api/pkg/infra/http/middleware"
)

type RouterMod struct {
	DBModule *db.DbMod
}

func NewRouterMod(dbmod *db.DbMod) *RouterMod {
	return &RouterMod{
		DBModule: dbmod,
	}
}

func (rm *RouterMod) Start(dbmod *db.DbMod) chi.Router {
	ur := dbmod.UserRepository
	tr := dbmod.TaskRepository
	md := middleware.NewMiddleware(ur, tr)
	r := chi.NewRouter()

	r.Get("/health", healthCheck)
	UserRouter(r, ur)
	TaskRouter(r, tr, md)

	return r
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status": "ok"}`)
}