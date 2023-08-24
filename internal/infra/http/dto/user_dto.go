package dto

import (
	"errors"
	"fmt"
	"regexp"
)

type UserCtx string

const (
	UserId UserCtx = "userId"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *RegisterUserRequest) Validate() error {
	reqFields := requiredFields(*u)
	if len(reqFields) > 0 {
		return fmt.Errorf("Missing required fields: %v", reqFields)
	}

	if err := validateName(u.Name); err != nil {
		return err
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *LoginUserRequest) Validate() error {
	reqFields := requiredFields(*u)
	if len(reqFields) > 0 {
		return fmt.Errorf("Missing required fields: %v", reqFields)
	}

	if err := validateEmail(u.Email); err != nil {
		return err
	}

	if err := validatePassword(u.Password); err != nil {
		return err
	}

	return nil
}

func validateName(n string) error {
	if len(n) < 3 {
		return errors.New("Name must be at least 3 characters")
	}

	return nil
}

func validateEmail(e string) error {
	eregx := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

	if !eregx.MatchString(e) {
		return errors.New("Invalid email")
	}

	return nil
}

func validatePassword(p string) error {
	if len(p) < 6 {
		return errors.New("Password must be at least 6 characters")
	}

	return nil
}
