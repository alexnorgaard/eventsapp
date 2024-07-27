package db

import (
	config "github.com/alexnorgaard/eventsapp"
	"github.com/alexnorgaard/eventsapp/cmd/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conf := config.GetConf()
	dsn := "host=" + conf.Postgres.Host + " user=" + conf.Postgres.User + " password=" + conf.Postgres.Password + " dbname=" + conf.Postgres.Database + " port=" + conf.Postgres.Port + " sslmode=disable TimeZone=Europe/Copenhagen"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Event{},
		&model.User{},
		&model.Subscription{})
}
