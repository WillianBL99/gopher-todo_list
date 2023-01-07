package main

import (
	_ "github.com/lib/pq"
	"github.com/willianbl99/todo-list_api/pkg/infra/db"
	"github.com/willianbl99/todo-list_api/pkg/infra/http"
)

// @title Gopher Todo-list
// @description Esta é uma API que permite lidar com tarefas. Esta api está implementada com validação de rotas com Json Web Token (JWT) para melhor segurança ao acessar as rotas.
// @version 1.0.0
// @BasePath /
// @schemes http
func main() {
	db := db.NewDbMod()
	ht := http.NewHttpMod(db)
	ht.Start()
}
