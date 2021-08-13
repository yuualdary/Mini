package models

type Locations struct {
	ID           string `gorm:"primary_key"`
	LocationID   string `gorm:"type:varchar(191)"`
	LocationCity string
	CompanyID    string   `gorm:"type:varchar(191)"`
	Companies    Company  `gorm:"foreignKey:CompanyID"`
	Categories   Category `gorm:"foreignKey:LocationID"`

	LocationStreet string
}
