package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	ID            string `gorm : "size:36;not null;uniqueIndex;primary_key"`
	PedagangID    string `gorm:"size:36;not null"`
	JenisDagangID string `gorm:"size:36;not null"`
	LapakID       string `gorm:"size:36;not null"`
	AdminID       string `gorm:"size:36;not null"`
	Tanggal       string `gorm:"size:10;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
