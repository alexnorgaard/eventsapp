package model

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
	//TODO: Want this to be a separate table, but reverse declaration of FK with gorm makes it not possible
	Subscribers []User `json:"subscribers" gorm:"many2many:event_subscribers"`
}

// type Subscription struct {
// 	Model
// 	Event Event
// 	User  User
// }
