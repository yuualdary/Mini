package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
	Title      string
	Slug       string
	CategoryID string   `gorm:"type:varchar(191)"`
	Categories Category `gorm:"foreignKey:CategoryID"`
	Content    string
}
