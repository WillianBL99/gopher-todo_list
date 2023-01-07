package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

const (
	getalltqy  = `SELECT id, user_id, title, description, status FROM public.task WHERE user_id = $1`
	getbysttqy = `SELECT id, user_id, title, description, status FROM public.task WHERE user_id = $1 AND status = $2`
	getbyidtqy = `SELECT id, user_id, title, description, status FROM public.task WHERE id = $1`
	savetqy    = `
		INSERT INTO public.task (id, user_id, title, description, status, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6)`
	updatetqy = `
		UPDATE public.task
			SET title=$2, description=$3, status=$4, updated_at=$5
			WHERE id = $1`
	deletetqy = `DELETE FROM public.task WHERE id = $1`
)

type TaskRepositoryPostgres struct {
	Server *sql.DB
}

func (tp *TaskRepositoryPostgres) GetAll(userId uuid.UUID) ([]entity.Task, error) {
	rws, err := tp.Server.Query(getalltqy, userId)
	if err != nil {
		return nil, err
	}

	tasks := make([]entity.Task, 0)

	for rws.Next() {
		t := entity.Task{}
		err = rws.Scan(
			&t.Id,
			&t.UserId,
			&t.Title,
			&t.Description,
			&t.Status,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (tp *TaskRepositoryPostgres) GetByStatus(userId uuid.UUID, status entity.Status) ([]entity.Task, error) {
	rws, err := tp.Server.Query(getbysttqy, userId, status)
	if err != nil {
		return nil, err
	}

	tasks := make([]entity.Task, 0)

	for rws.Next() {
		t := entity.Task{}
		err = rws.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (tp *TaskRepositoryPostgres) GetById(id uuid.UUID) (entity.Task, error) {
	rw, err := tp.Server.Query(getbyidtqy, id)
	if err != nil {
		return entity.Task{}, err
	}

	t := entity.Task{}

	rw.Next()
	err = rw.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.Status)
	if err != nil {
		return entity.Task{}, err
	}

	defer rw.Close()
	return t, nil
}

func (tp *TaskRepositoryPostgres) Save(t *entity.Task) error {
	rw, err := tp.Server.Query(
		savetqy,
		t.Id,
		t.UserId,
		t.Title,
		t.Description,
		t.Status,
		time.Now(),
	)
	if err != nil {
		return err
	}
	defer rw.Close()
	return nil
}

func (tp *TaskRepositoryPostgres) Delete(id uuid.UUID) error {
	rw, err := tp.Server.Query(deletetqy, id)
	if err != nil {
		return herr.NewApp().Conflict
	}
	defer rw.Close()
	return nil
}

func (tp *TaskRepositoryPostgres) Update(t *entity.Task) error {
	rw, err := tp.Server.Query(
		updatetqy,
		t.Id,
		t.Title,
		t.Description,
		t.Status,
		time.Now(),
	)
	if err != nil {
		return err
	}

	defer rw.Close()
	return nil
}
