package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/willianbl99/todo-list_api/config"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

type PostgresServer struct {
	Db *sql.DB
}

func (ps *PostgresServer) Connect() *sql.DB {

	dbcf := config.NewAppConf().Database

	var err error
	fmt.Printf("Conn. string: %s\n", dbcf.ConnStr())
	ps.Db, err = sql.Open("postgres", dbcf.ConnStr())
	herr.CheckError(err)

	if err := ps.Db.Ping(); err != nil {
		herr.CheckError(err)
	}

	fmt.Println("Database connected")
	return ps.Db
}

func (ps *PostgresServer) Stop() {
	ps.Db.Close()
}
