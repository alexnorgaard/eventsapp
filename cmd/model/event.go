package models

import (
	"time"

	"github.com/codingsince1985/geo-golang"
)

type Event struct {
	Model
	Title         string        `json:"title" gorm:"not null"`
	Description   string        `json:"description" gorm:"default:NULL"`
	Owners        []User        `json:"owner" gorm:"many2many:event_owners"`
	Private_event bool          `json:"private_event" gorm:"default:false"`
	Time_start    time.Time     `json:"time_start"`
	Time_end      time.Time     `json:"time_end" gorm:"default:NULL;check:time_end > time_start"`
	Address       *geo.Address  `json:"address" gorm:"embedded"`
	Geolocation   *geo.Location `json:"geolocation" gorm:"embedded"`
}

type Subscription struct {
	Model
	Event Event `gorm:"embedded;not null"`
	User  User  `gorm:"embedded;not null"`
}
