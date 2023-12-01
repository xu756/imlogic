package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID             int64 `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      sql.NullTime `gorm:"index"`
	CreatedAtInt64 int64        `gorm:"-" json:"createdAtInt64"`
	UpdatedAtInt64 int64        `gorm:"-" json:"updatedAtInt64"`
}
