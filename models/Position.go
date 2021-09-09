package models

type Position struct {
	//gorm.Model
	ID                  string `gorm:"primary_key"`
	PositionName        string
	PositionDescription string
	PositionFee         int
	PositionLength      int
	PositionRequirement string
	CompanyID           string  `gorm:"type:varchar(191)"`
	Companies           Company `gorm:"foreignKey:CompanyID"`
}
