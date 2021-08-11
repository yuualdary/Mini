package models

import (
	"time"

	"gorm.io/gorm"
)

type Otps struct {
	gorm.Model
	ID      string `gorm:"primary_key"`
	Value   int
	UsersID string
	User    Users `gorm:"foreignKey:UsersID"`
	Expired time.Time
}
