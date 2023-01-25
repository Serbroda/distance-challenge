package db

import (
	"log"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect(name string) *gorm.DB {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", name, err)
			panic(err)
		}
		DB = db
	})
	return DB
}

func Migrate(db *gorm.DB) {
	//db.AutoMigrate(&models.User{}, &models.Group{}, &models.Link{}, &models.GroupSubscription{})
}
