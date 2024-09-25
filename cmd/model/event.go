package model

import (
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/lib/pq"
)

type Event struct {
	Model
	Title            string         `json:"title" gorm:"not null;index:title_idx,type:GIN" validate:"required"`
	Description      string         `json:"description" gorm:"default:NULL"`
	Banner_url       string         `json:"banner_url" gorm:"default:NULL" validate:"omitempty,url"`
	Owners           []User         `json:"owners" gorm:"many2many:event_owners"`
	Is_private_event bool           `json:"private_event" gorm:"default:false"`
	Time_start       time.Time      `json:"time_start" gorm:"required" validate:"required,gt"` //gt - For time.Time ensures the time value is greater than time.Now.UTC()
	Time_end         time.Time      `json:"time_end" gorm:"default:NULL;check:time_end > time_start" validate:"omitempty,gtefield=Time_start"`
	Address          *geo.Address   `json:"address" gorm:"embedded" validate:"required"`
	Geolocation      *geo.Location  `json:"geolocation" gorm:"embedded"`
	Tags             pq.StringArray `json:"tags" query:"tags" gorm:"type:text[];index:tags_idx,type:GIN"`
	//TODO: Want this to be a separate table, but reverse declaration of FK with gorm makes it not possible
	Subscribers []User `json:"subscribers" gorm:"many2many:event_subscribers"`
	Is_enabled  bool   `json:"is_enabled" gorm:"default:true"`

	//Not stored in DB
	Image []byte `form:"image"`
}

// type Subscription struct {
// 	Model
// 	Event Event
// 	User  User
// }
