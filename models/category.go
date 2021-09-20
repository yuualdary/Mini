package models

type Category struct {
	ID             string `gorm:"primary_key"`
	CategoryPrefix string
	CategoryName   string
	CategorySlug   string
}
