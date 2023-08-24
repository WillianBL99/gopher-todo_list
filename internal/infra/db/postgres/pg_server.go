package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/willianbl99/todo-list_api/config"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type PostgresServer struct {
	Db *sql.DB
}

func (ps *PostgresServer) Connect() *sql.DB {

	dbcf := config.NewAppConf().Database

	var err error
	ps.Db, err = sql.Open("postgres", dbcf.ConnStr())
	if err != nil {
		e.New().InfraDbErr(e.ErrorOnConnectDb, err.Error()).Fatal()
	}
	if err := ps.Db.Ping(); err != nil {
		e.New().InfraDbErr(e.ErrorOnConnectDb, err.Error()).Fatal()
	}

	fmt.Println("Database connected")
	return ps.Db
}

func (ps *PostgresServer) Stop() {
	ps.Db.Close()
}
