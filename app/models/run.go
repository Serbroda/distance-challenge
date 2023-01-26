package models

import (
	"gorm.io/gorm"
)

type Run struct {
	gorm.Model
	UserId   uint  `json:"user_id"`
	Distance int64 `json:"distance"`
	Time     int64 `json:"time"`
}
