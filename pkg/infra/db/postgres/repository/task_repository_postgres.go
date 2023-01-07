package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

const (
	getalltqy = `SELECT * FROM public.task WHERE user_id = $1`
	getbysttqy = `SELECT * FROM public.task WHERE user_id = $1 AND status = $2`
	getbyidtqy = `SELECT * FROM public.task WHERE id = $1`
	savetqy = `INSERT INTO public.task (id, user_id, title, description, status) VALUES ($1, $2, $3, $4, $5)`
	updatetqy = `UPDATE public.task SET title = $1, description = $2, status = $3 WHERE id = $4`
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
		err = rws.Scan(&t.Id, &t.UserId, &t.Title, &t.Description, &t.Status)
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
	rw, err := tp.Server.Query(savetqy, t.Id, t.UserId, t.Title, t.Description, t.Status)
	if err != nil {
		return err
	}
	defer rw.Close()
	return nil
}

func (tp *TaskRepositoryPostgres) Delete(id uuid.UUID) error {
	rw, err := tp.Server.Query(deletetqy, id)
	if err != nil {
		return err
	}
	defer rw.Close()
	return nil
}

func (tp *TaskRepositoryPostgres) Update(t *entity.Task) error {
	rw, err := tp.Server.Query(updatetqy, t.Title, t.Description, t.Status, t.Id)
	if err != nil {
		return err
	}

	defer rw.Close()
	return nil
}
