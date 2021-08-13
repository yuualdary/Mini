package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID                 string `gorm:"primary_key"`
	CompanyName        string
	CompanyDescription string
	UserID             string `gorm:"type:varchar(191)"`
	Users              Users  `gorm:"foreignKey:UserID"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
