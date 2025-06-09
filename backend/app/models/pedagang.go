package models

import (
	"time"

	"gorm.io/gorm"
)

type Pedagang struct {
	ID            string `gorm : "size:36;not null;uniqueIndex;primary_key"`
	firstName     string `gorm:"size:100;not null"`
	lastName      string `gorm:"size:100;not null"`
	JenisDagangID string `gorm:"size:36;not null"`
	LapakID       string `gorm:"size:36;not null"`
	NoHp          string `gorm:"size:15;not null;uniqueIndex"`

	Email         string `gorm:"size:100;not null;uniqueIndex"`
	Password      string `gorm:"size:255;not null"`
	rememberToken string `gorm:"size:255;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
