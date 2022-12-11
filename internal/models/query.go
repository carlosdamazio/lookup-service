package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Query struct {
	ID        uuid.UUID      `json:"-"`
	Domain    string         `json:"domain"`
	ClientIP  string         `json:"client_ip"`
	Addresses pq.StringArray `json:"addresses" gorm:"type:text[]"`
	CreatedAt time.Time      `json:"created_at"`
}

func (q *Query) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}
	q.ID = id

	return nil
}
