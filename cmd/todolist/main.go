package main

import (
	_ "github.com/lib/pq"
	"github.com/willianbl99/todo-list_api/internal/infra/db"
	"github.com/willianbl99/todo-list_api/internal/infra/http"
)

func main() {
	db := db.NewDbMod()
	ht := http.NewHttpMod(db)
	ht.Start()
}
