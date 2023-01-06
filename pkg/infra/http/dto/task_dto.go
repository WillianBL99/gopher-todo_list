package dto

import "errors"

type SaveTask struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

func (s *SaveTask) Validate() error {
	if s.Title == "" {
		return errors.New("Title is required")
	}

	if len(s.Title) > 50 {
		return errors.New("Title must be less than 50 characters")
	}

	if s.Description == "" {
		return errors.New("Description is required")
	}

	if len(s.Description) > 255 {
		return errors.New("Description must be less than 255 characters")
	}

	return nil
}