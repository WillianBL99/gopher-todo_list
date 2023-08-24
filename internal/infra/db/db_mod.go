package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	rp "github.com/willianbl99/todo-list_api/internal/application/repository"
	"github.com/willianbl99/todo-list_api/internal/infra/db/postgres"
	"github.com/willianbl99/todo-list_api/internal/infra/db/postgres/repository"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

type DbMod struct {
	db             *sql.DB
	UserRepository rp.UserRepository
	TaskRepository rp.TaskRepository
}

// This function is used to connect to the database and return the repositories
func NewDbMod() *DbMod {
	srvdb := postgres.PostgresServer{}
	conndb := srvdb.Connect()

	urp := repository.UserRepositoryPostgres{Server: conndb}
	trp := repository.TaskRepositoryPostgres{Server: conndb}

	dbMod := &DbMod{
		db:             conndb,
		UserRepository: &urp,
		TaskRepository: &trp,
	}
	dbMod.initDatabase()
	return dbMod
}

func (db *DbMod) initDatabase() {
	existsTablesQr := "SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';"
	rw, err := db.db.Query(existsTablesQr)
	if err != nil {
		e.New().InfraDbErr(e.InternalServerError, err.Error()).Fatal()
	}
	defer rw.Close()

	var tables []string
	for rw.Next() {
		var table string
		err = rw.Scan(&table)
		if err != nil {
			e.New().InfraDbErr(e.InternalServerError, err.Error()).Fatal()
		}
		tables = append(tables, table)
	}
	if len(tables) <= 1 {
		createDatabase(db.db)
	}
}

func createDatabase(db *sql.DB) {
	sqlFileContent, err := os.ReadFile("/app/pkg/infra/db/postgres/create-tables.sql")
	if err != nil {
		e.New().InfraDbErr(e.InternalServerError, err.Error()).Fatal()
	}

	// divide the file content into SQL commands
	sqlCommands := strings.Split(string(sqlFileContent), ";")

	// execute each command one by one
	for _, cmd := range sqlCommands {
		trimmedCmd := strings.TrimSpace(cmd)
		if trimmedCmd == "" {
			continue
		}

		_, err := db.Exec(trimmedCmd)
		if err != nil {
			e.New().InfraDbErr(e.InternalServerError, err.Error()).Fatal()
		}
	}

	fmt.Println("Database created successfully!")
}
