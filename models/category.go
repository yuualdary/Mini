package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryPrefix string
	CategoryName   string
	CategorySlug   string
}
