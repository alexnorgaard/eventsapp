package models

type User struct {
	BaseModel
	Name string `json:"name" gorm:"unique"`
}
