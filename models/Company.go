package models

type Company struct {
	ID                 string `gorm:"primary_key"`
	CompanyName        string
	CompanyDescription string
}
