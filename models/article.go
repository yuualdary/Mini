package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string 
	Slug string
	CategoryID int 
	Categories Category  `gorm:"foreignKey:CategoryID"`
	Content string
	

}