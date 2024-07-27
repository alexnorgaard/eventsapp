package model

type User struct {
	Model
	Username string `json:"username" gorm:"unique;not null"`
}
