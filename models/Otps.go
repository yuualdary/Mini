package models

import (
	"time"
	"gorm.io/gorm"
)


type Otps struct {
	gorm.Model
	Value   int
	UsersID int 
	User Users `gorm:"foreignKey:UsersID"`
	Expired time.Time
}