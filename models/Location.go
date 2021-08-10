package models

type Locations struct {
	ID           string `gorm:"primary_key"`
	LocationID   int
	LocationCity string
	CompanyID    int
	Companies    Company  `gorm:"foreignKey:CompanyID"`
	Categories   Category `gorm:"foreignKey:LocationID"`

	LocationStreet string
}
