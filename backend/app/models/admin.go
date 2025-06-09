package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID         string `gorm : "size:36;not null;uniqueIndex;primary_key"`
	Username   string `gorm:"size:100;not null;uniqueIndex"`
	Password   string `gorm:"size:255;not null"`
	Email      string `gorm:"size:100;not null;uniqueIndex"`
	NoHp       string `gorm:"size:15;not null;uniqueIndex"`
	SuperAdmin bool   `gorm:"default:false"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
