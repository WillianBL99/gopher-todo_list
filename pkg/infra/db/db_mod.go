package db

import (
	rp "github.com/willianbl99/todo-list_api/pkg/application/repository"
	"github.com/willianbl99/todo-list_api/pkg/infra/db/postgres"
	"github.com/willianbl99/todo-list_api/pkg/infra/db/postgres/repository"
)

type DbMod struct {
	UserRepository rp.UserRepository
	TaskRepository rp.TaskRepository
}

// NewDbMod returns a new DbMod
// This function is used to connect to the database and return the repositories
// @title Gopher Todo-list modulo
func NewDbMod() *DbMod {
	srvdb := postgres.PostgresServer{}
 	conndb := srvdb.Connect()

	urp := repository.UserRepositoryPostgres{Server: conndb}
	trp := repository.TaskRepositoryPostgres{Server: conndb}

	return &DbMod{
		UserRepository: &urp,
		TaskRepository: &trp,
	}
}
