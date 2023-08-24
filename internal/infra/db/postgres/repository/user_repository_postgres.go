package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/internal/application/entity"
	e "github.com/willianbl99/todo-list_api/pkg/herr"
)

var (
	mUFields   = `id, name, email, password`
	dUFields   = `created_at, modified_at, deleted_at`
	saveQr     = fmt.Sprintf(`INSERT INTO users (%s, created_at) VALUES ($1, $2, $3, $4, $5)`, mUFields)
	getById    = fmt.Sprintf(`SELECT %s, %s FROM users WHERE id = $1`, mUFields, dUFields)
	getByEmail = fmt.Sprintf(`SELECT %s, %s FROM users WHERE email = $1`, mUFields, dUFields)
)

type UserRepositoryPostgres struct {
	Server *sql.DB
}

func (up *UserRepositoryPostgres) GetById(id uuid.UUID) (entity.User, *e.Error) {
	rw, err := up.Server.Query(getById, id)
	if err != nil {
		return entity.User{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	u := entity.User{}

	rw.Next()
	err = rw.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return entity.User{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rw.Close()

	if u.Email == "" {
		return entity.User{}, e.New().InfraDbErr(e.UserNotFound)
	}

	return u, nil
}

func (up *UserRepositoryPostgres) GetByEmail(email string) (entity.User, *e.Error) {
	rw, err := up.Server.Query(getByEmail, email)
	if err != nil {
		return entity.User{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rw.Close()

	u := entity.User{}

	rw.Next()
	err = rw.Scan(&u.Id, &u.Name, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt)
	if err != nil {
		return entity.User{}, e.New().InfraDbErr(e.InternalServerError, err.Error())
	}

	if u.Email == "" {
		return entity.User{}, e.New().InfraDbErr(e.UserNotFound)
	}

	return u, nil
}

func (up *UserRepositoryPostgres) Save(u *entity.User) *e.Error {
	rw, err := up.Server.Query(saveQr, u.Id, u.Name, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return e.New().InfraDbErr(e.InternalServerError, err.Error())
	}
	defer rw.Close()
	return nil
}
