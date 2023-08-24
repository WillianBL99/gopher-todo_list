package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"github.com/willianbl99/todo-list_api/config"
	"github.com/willianbl99/todo-list_api/internal/infra/db"
	"github.com/willianbl99/todo-list_api/internal/infra/http/middleware"
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

	url := config.NewAppConf().API.URL
	filesDir := config.NewAppConf().API.WorkDir
	handlerFiles := http.FileServer(http.Dir(filesDir))
	r.Handle("/docs/*", http.StripPrefix("/docs/", handlerFiles))
	r.Get("/", version)
	r.Get("/health", healthCheck)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(url+"/docs/swagger.yml"),
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
