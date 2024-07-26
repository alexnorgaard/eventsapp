package db

import (
	model "github.com/alexnorgaard/eventsapp/cmd/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=192.168.0.49 user=appaccess password=eventsapp dbname=eventsappdb port=5432 sslmode=disable TimeZone=Europe/Copenhagen"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Event{},
		&model.User{},
		&model.Subscription{})
}
