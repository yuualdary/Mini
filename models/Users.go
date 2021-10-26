package models

type Users struct {
	ID           string `gorm:"primary_key"`
	Name         string
	Email        string `gorm:"type:varchar(191);unique_index"`
	Bod          string
	Initial      string
	Password     string
	IsVerif      bool
	ProfilePhoto string
	Role         string
	Token        string
}
