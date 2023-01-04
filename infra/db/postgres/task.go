package postgres

import "github.com/google/uuid"

type Task struct {
	ID          int64
	Title       string
	Description string
}

func (t *Task) Get(id uuid.UUID) (*Task, error) {
	return nil, nil
}

func (t *Task) Save() error {
	return nil
}

func (t *Task) Delete() error {
	return nil
}

func (t *Task) Update() error {
	return nil
}

func (t *Task) GetAll() ([]*Task, error) {
	return nil, nil
}
