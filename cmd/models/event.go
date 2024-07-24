package models

import (
	"time"

	"github.com/codingsince1985/geo-golang"
)

type Event struct {
	BaseModel
	Title         string       `json:"title" gorm:"not null"`
	Description   string       `json:"description" gorm:"default:NULL"`
	Owner         [3]User      `json:"owner"`
	Private_event bool         `json:"private_event" gorm:"default:false"`
	Time_start    time.Time    `json:"time_start"`
	Time_end      time.Time    `json:"time_end" gorm:"default:NULL;check:time_end > time_start"`
	Address       geo.Address  `json:"address"`
	Geolocation   geo.Location `json:"geolocation" gorm:"default:NULL"`
}
