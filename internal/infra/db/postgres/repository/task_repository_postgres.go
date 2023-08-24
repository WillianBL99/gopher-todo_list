package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

var (
	mtfields      = `id, list, title, description, status, user_id`
	defaultfields = `created_at, modified_at, deleted_at`
	getAllQr      = fmt.Sprintf(`SELECT %s, %s FROM tasks WHERE user_id = $1 AND deleted_at IS NULL`, mtfields, defaultfields)
	getByStatusQr = fmt.Sprintf(`SELECT %s, %s FROM tasks WHERE user_id = $1 AND status = $2 AND deleted_at IS NULL`, mtfields, defaultfields)
	getByIdQr     = fmt.Sprintf(`SELECT %s, %s FROM tasks WHERE id = $1 AND deleted_at IS NULL`, mtfields, defaultfields)
	getByListQr   = fmt.Sprintf(`SELECT %s, %s FROM tasks WHERE user_id = $1 AND list = $2 AND deleted_at IS NULL`, mtfields, defaultfields)
	SaveTQr       = fmt.Sprintf(`INSERT INTO tasks (%s, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`, mtfields)
	updateTQr     = `
		UPDATE tasks
			SET title=$2, description=$3, status=$4, modified_at=$5
			WHERE id = $1`
	deleteTQr = `UPDATE tasks SET deleted_at=$2 WHERE id = $1`
)

type TaskRepositoryPostgres struct {
	Server *sql.DB
}

func (tp *TaskRepositoryPostgres) GetAll(userId uuid.UUID) ([]entity.Task, *e.Error) {
	rws, err := tp.Server.Query(getAllQr, userId)
	if err != nil {
		return nil, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rws.Close()

	tasks := make([]entity.Task, 0)

	for rws.Next() {
		t := entity.Task{}
		err = rws.Scan(
			&t.Id,
			&t.List,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.UserId,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt,
		)
		if err != nil {
			return nil, e.New().InfraDbErr(e.InternalServerError, err.Error())
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (tp *TaskRepositoryPostgres) GetByStatus(userId uuid.UUID, status entity.Status) ([]entity.Task, *e.Error) {
	rws, err := tp.Server.Query(getByStatusQr, userId, status)
	if err != nil {
		return nil, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rws.Close()

	tasks := make([]entity.Task, 0)

	for rws.Next() {
		t := entity.Task{}
		err = rws.Scan(
			&t.Id,
			&t.List,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.UserId,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt,
		)
		if err != nil {
			return nil, e.New().InfraDbErr(e.InternalServerError, err.Error())
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (tp *TaskRepositoryPostgres) GetById(id uuid.UUID) (entity.Task, *e.Error) {
	rw, err := tp.Server.Query(getByIdQr, id)
	if err != nil {
		return entity.Task{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rw.Close()
	if !rw.Next() {
		return entity.Task{}, e.New().InfraDbErr(e.TaskNotFound)
	}

	t := entity.Task{}
	err = rw.Scan(
		&t.Id,
		&t.List,
		&t.Title,
		&t.Description,
		&t.Status,
		&t.UserId,
		&t.CreatedAt,
		&t.UpdatedAt,
		&t.DeletedAt,
	)
	if err != nil {
		return entity.Task{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}

	return t, nil
}

func (tp *TaskRepositoryPostgres) GetByList(uid uuid.UUID, listName string) ([]entity.Task, *e.Error) {
	rws, err := tp.Server.Query(getByListQr, uid, listName)
	if err != nil {
		return nil, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rws.Close()

	tasks := make([]entity.Task, 0)
	for rws.Next() {
		var t entity.Task
		err := rws.Scan(
			&t.Id,
			&t.List,
			&t.Title,
			&t.Description,
			&t.Status,
			&t.UserId,
			&t.CreatedAt,
			&t.UpdatedAt,
			&t.DeletedAt,
		)
		if err != nil {
			return nil, e.New().InfraDbErr(e.InternalServerError, err.Error())
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (tp *TaskRepositoryPostgres) Save(t *entity.Task) (entity.Task, *e.Error) {
	rw, err := tp.Server.Query(
		SaveTQr,
		t.Id,
		t.List,
		t.Title,
		t.Description,
		t.Status,
		t.UserId,
		time.Now(),
	)
	if err != nil {
		return entity.Task{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rw.Close()
	return *t, nil
}

func (tp *TaskRepositoryPostgres) Delete(id uuid.UUID) *e.Error {
	rw, err := tp.Server.Query(deleteTQr, id, time.Now())
	if err != nil {
		return e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rw.Close()
	return nil
}

func (tp *TaskRepositoryPostgres) Update(t *entity.Task) *e.Error {
	rw, err := tp.Server.Query(
		updateTQr,
		t.Id,
		t.List,
		t.Title,
		t.Description,
		t.Status,
		time.Now(),
	)
	if err != nil {
		return e.New().InfraDbErr(e.InternalServerError, err.Error())
	}

	defer rw.Close()
	return nil
}
