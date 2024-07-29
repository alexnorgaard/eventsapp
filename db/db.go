package db

import (
	"fmt"

	config "github.com/alexnorgaard/eventsapp"
	"github.com/alexnorgaard/eventsapp/cmd/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conf := config.GetConf()
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Copenhagen", conf.Postgres.Host, conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Database, conf.Postgres.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Event{},
		&model.User{},
	)
}
