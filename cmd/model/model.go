package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID    `gorm:"<-:create;type:uuid;primary_key;default:uuid_generate_v4()" param:"id" validate:"omitempty,uuid4"`
	CreatedAt time.Time    `gorm:"autoCreateTime:milli"`
	UpdatedAt time.Time    `gorm:"autoUpdateTime:milli"`
	DeletedAt sql.NullTime `gorm:"index"`
}
