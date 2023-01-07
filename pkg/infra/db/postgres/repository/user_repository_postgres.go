package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/willianbl99/todo-list_api/pkg/application/entity"
)

const (
	saveqy = `INSERT INTO public.user (id, name, email, password) VALUES ($1, $2, $3, $4)`
	gbyid  = `SELECT * FROM public.user WHERE id = $1`
	gbyem  = `SELECT * FROM public.user WHERE email = $1`
)

type UserRepositoryPostgres struct {
	Server *sql.DB
}

func (up *UserRepositoryPostgres) GetById(id uuid.UUID) (entity.User, error) {
	panic("implement me")
}

func (up *UserRepositoryPostgres) GetByEmail(email string) (entity.User, error) {
	rw, err := up.Server.Query(gbyem, email)
	if err != nil {
		fmt.Println("errz: ", err)
		return *&entity.User{}, err
	}

	u := entity.User{}

	rw.Next()
	err = rw.Scan(&u.Id, &u.Name, &u.Email, &u.Password)
	if err != nil {
		fmt.Println("errz: ", err)
		return *&entity.User{}, err
	}

	defer rw.Close()
	return u, nil
}

func (up *UserRepositoryPostgres) Save(u *entity.User) error {
	rw, err := up.Server.Query(saveqy, u.Id, u.Name, u.Email, u.Password)
	if err != nil {
		fmt.Println("errz: ", err)
		return err
	}
	defer rw.Close()
	return nil
}
