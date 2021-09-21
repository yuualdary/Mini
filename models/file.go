package models

import "time"

type file struct {
	ID      string `gorm:"primary_key"`
	UserID  string `gorm:"type:varchar(191)"`
	User    Users  `gorm:"foreignKey:UsersID"`
	Expired time.Time
}



