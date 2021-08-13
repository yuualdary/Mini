package models

import (
	"time"
)

type Otps struct {
	ID      string `gorm:"primary_key"`
	Value   int
	UsersID string `gorm:"type:varchar(191)"`
	User    Users  `gorm:"foreignKey:UsersID"`
	Expired time.Time
}
