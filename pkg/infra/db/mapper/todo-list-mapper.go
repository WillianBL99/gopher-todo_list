package mapper

import (
	"database/sql"

	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

func ToDomain(t *sql.Rows) *entity.Task {
	return entity.NewTask(
		t.Id,
		t.Title,
		t.Description,
		t.DueDate,
	)
}