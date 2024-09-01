package model

import (
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/lib/pq"
)

type Event struct {
	Model
	Title         string         `json:"title" gorm:"not null;index:title_idx,type:GIN" validate:"required"`
	Description   string         `json:"description" gorm:"default:NULL"`
	Owners        []User         `json:"owner" gorm:"many2many:event_owners"`
	Private_event bool           `json:"private_event" gorm:"default:false"`
	Time_start    time.Time      `json:"time_start" validate:"required,gt"`
	Time_end      time.Time      `json:"time_end" gorm:"default:NULL;check:time_end > time_start" validate:"omitempty,gtefield=Time_start"`
	Address       *geo.Address   `json:"address" gorm:"embedded" validate:"required"`
	Geolocation   *geo.Location  `json:"geolocation" gorm:"embedded"`
	Tags          pq.StringArray `json:"tags" query:"tags" gorm:"type:text[];index:tags_idx,type:GIN"`
	//TODO: Want this to be a separate table, but reverse declaration of FK with gorm makes it not possible
	Subscribers []User `json:"subscribers" gorm:"many2many:event_subscribers"`
	Is_enabled  bool   `json:"is_enabled" gorm:"default:true"`
}

// type Subscription struct {
// 	Model
// 	Event Event
// 	User  User
// }
