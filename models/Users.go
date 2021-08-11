package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name         string
	Email        string
	Bod          time.Time
	Initial      string
	Password     string
	IsVerif      bool
	ProfilePhoto string
	Role         string
	Token        string
	CompanyID    int
	Companies    Company `gorm:"foreignKey:CompanyID"`
}
