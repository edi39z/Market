package models

type Lapak struct {
	ID   string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Nama string `gorm:"size:100;not null;uniqueIndex"`
}
