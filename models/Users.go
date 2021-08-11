package models

import (
	"time"
)

type Users struct {
	ID           string `gorm:"primary_key"`
	Name         string
	Email        string
	Bod          time.Time
	Initial      string
	Password     string
	IsVerif      bool
	ProfilePhoto string
	Role         string
	Token        string
}
