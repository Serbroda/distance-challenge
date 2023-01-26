package models

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Username  string     `json:"username"`
	Password  string     `json:"-"`
	Active    bool       `json:"active"`
	IsAdmin   bool       `json:"is_admin"`
}
