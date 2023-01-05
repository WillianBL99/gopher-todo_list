package entity

import "github.com/google/uuid"

type User struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(id uuid.UUID, name string, email string, password string) *User {
	b := Base{}
	b.New(id)

	u := User{
		Name:     name,
		Email:    email,
		Password: password,
		Base:     b,
	}

	return &u
}
