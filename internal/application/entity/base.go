package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Base struct {
	Id        uuid.UUID    `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt sql.NullTime `json:"-"`
	UpdatedAt sql.NullTime `json:"modified_at"`
}

func (b *Base) New(id uuid.UUID) {
	if id != uuid.Nil {
		id = uuid.New()
	}

	b.Id = id
	b.CreatedAt = time.Now()
}
