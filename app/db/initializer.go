package db

import (
	"github.com/Serbroda/distance-challenge/models"
	"github.com/Serbroda/distance-challenge/security"
	"gorm.io/gorm"
)

func InitializeData(db *gorm.DB) error {
	var admin models.User
	res := db.Where("lower(username) = lower(?)", "admin").First(&admin)
	if res.Error != nil {
		admin = models.User{
			Username: "admin",
			Password: security.MustHashPassword("d1st@nce"),
			Active:   true,
			IsAdmin:  true,
		}
		res := db.Create(&admin)
		if res.Error != nil {
			return res.Error
		}
	}
	return nil
}
