package model

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint64 `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
