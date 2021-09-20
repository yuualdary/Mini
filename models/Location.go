package models

type Locations struct {
	ID           string `gorm:"primary_key"`
	LocationCity string

	// Categories   Category `gorm:"foreignKey:LocationID"`

	// LocationStreet string
}
