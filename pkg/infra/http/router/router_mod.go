package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
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

	r.Get("/", version)
	r.Get("/health", healthCheck)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))
	TaskRouter(r, tr, md)
	UserRouter(r, ur)

	return r
}

func version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"version": "1.0.0"}`)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status": "ok"}`)
}
