package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
	"github.com/willianbl99/todo-list_api/pkg/herr"
)

const (
	saveqy = `INSERT INTO users (id, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`
	gbyid  = `SELECT * FROM users WHERE id = $1`
	gbyem  = `SELECT * FROM users WHERE email = $1`
)

type UserRepositoryPostgres struct {
	Server *sql.DB
}

func (up *UserRepositoryPostgres) GetById(id uuid.UUID) (entity.User, error) {
	rw, err := up.Server.Query(gbyid, id)
	if err != nil {
		return entity.User{}, err
	}

	u := entity.User{}

	rw.Next()
	err = rw.Scan(&u.Id, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return entity.User{}, err
	}

	if u.Email == "" {
		return entity.User{}, herr.NewApp().UserNotFound
	}

	defer rw.Close()
	return u, nil
}

func (up *UserRepositoryPostgres) GetByEmail(email string) (entity.User, error) {
	rw, err := up.Server.Query(gbyem, email)
	if err != nil {
		return entity.User{}, err
	}

	u := entity.User{}

	rw.Next()
	err = rw.Scan(&u.Id, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return entity.User{}, err
	}

	if u.Email == "" {
		return entity.User{}, herr.NewApp().UserNotFound
	}

	defer rw.Close()
	return u, nil
}

func (up *UserRepositoryPostgres) Save(u *entity.User) error {
	rw, err := up.Server.Query(saveqy, u.Id, u.Name, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return err
	}
	defer rw.Close()
	return nil
}
