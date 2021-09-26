package models

type Category struct {
	ID                 string `gorm:"primary_key"`
	CategoryPrefix     string
	CategoryName       string
	CategorySlug       string
	PositionCategories []PositionCategory
}

type PositionCategory struct {
	CategoryID string `gorm:"primary_key"`
	PositionID string `gorm:"primary_key"`
}

//buat tag di detail position
