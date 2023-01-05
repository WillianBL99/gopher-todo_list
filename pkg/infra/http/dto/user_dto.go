package dto

import (
	"errors"
	"regexp"
)

type UserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func (u *UserRequest) Validate() error {
	var err error
	err = u.validateName()
	if err != nil {
		return err
	}

	err = u.validateEmail()
	if err != nil {
		return err
	}

	err = u.validatePassword()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRequest) validateName() error {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if len(u.Name) < 3 {
		return errors.New("Name must be at least 3 characters")
	}

	return nil
}

func (u *UserRequest) validateEmail() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}

	eregx := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

	if !eregx.MatchString(u.Email) {
		return errors.New("Invalid email")
	}

	return nil
}

func (u *UserRequest) validatePassword() error {
	if u.Password == "" {
		return errors.New("Password is required")
	}

	if len(u.Password) < 6 {
		return errors.New("Password must be at least 6 characters")
	}

	return nil
}