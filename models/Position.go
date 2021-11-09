package models

type Position struct {
	//gorm.Model
	ID                  string `gorm:"primary_key"`
	PositionName        string
	PositionDescription string
	PositionFee         int
	PositionLength      int
	PositionRequirement string
	PositionSlug string
	Candidates []Candidate 
	PositionCategories []PositionCategory 

	CompanyID      string  `gorm:"type:varchar(191)"`
	Companies           Company `gorm:"foreignKey:CompanyID"`
}


type Bookmarks struct {
	UserID     string `gorm:"primary_key"`
	PositionID string `gorm:"primary_key"`
}