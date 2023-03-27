package model

import (
	"time"
)

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Author    string
	Title     string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
