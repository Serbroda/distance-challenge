package models

import (
	"time"
)

type Run struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	UserId    uint       `json:"user_id"`
	User      User       `json:"-"`
	Distance  int64      `json:"distance"`
	Time      int64      `json:"time"`
}
