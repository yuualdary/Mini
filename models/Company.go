package models

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID                 string `gorm:"primary_key"`
	CompanyName        string
	CompanyDescription string
	UserID             string    `gorm:"type:varchar(191)"`
	Users              Users     `gorm:"foreignKey:UserID"`
	LocationID         string    `gorm:"type:varchar(191)"`
	Location           Locations `gorm:"foreignKey:LocationID"`
	CategoryID         string    `gorm:"type:varchar(191)"`
	Categories           Category `gorm:"foreignKey:CategoryID"`
	CompanySlug string	
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
//list candidate formatter
//accept reject