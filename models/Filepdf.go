package models

import "time"

type Filepdf struct {
	ID       string `gorm:"primary_key"`
	Filename string
	UserID   string `gorm:"type:varchar(191)"`
	User    Users  `gorm:"foreignKey:UserID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

}
