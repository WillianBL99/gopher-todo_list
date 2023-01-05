package repository

import (
	"database/sql"

	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/infra/db/mapper"
)

type PostgresRepository struct {
	Server *sql.DB
}

func (r *PostgresRepository) GetAll() ([]entity.Task, error) {
	tks_raw, err := r.Server.Query("SELECT * FROM tasks")

	defer tks_raw.Close()

	if err != nil {
		return nil, err
	}

	tks := make([]entity.Task, 0)

	for tks_raw.Next() {
		if err != nil {
			return nil, err
		}

		tk := mapper.ToDomain(tks_raw)
		tks = append(tks, *tk)
	}

	return tks, nil
}

