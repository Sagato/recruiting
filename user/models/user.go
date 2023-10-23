package user_models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey" json:"id,omitempty"`
	Firstname string  `json:"firstname,omitempty"`
	Lastname  string  `json:"lastname,omitempty"`
	Profile   Profile `json:"profile,omitempty"`
}
