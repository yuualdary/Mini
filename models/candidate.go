package models

import "time"

type Candidate struct {
	ID          string   `gorm:"primary_key"`
	PositionID  string   `gorm:"type:varchar(191)"`
	Positions   Position `gorm:"foreignKey:PositionID"`
	UserID      string   `gorm:"type:varchar(191)"`
	Users       Users    `gorm:"foreignKey:UserID"`
	CategoryID  string   `gorm:"type:varchar(191)"`
	Categories  Category `gorm:"foreignKey:CategoryID"`
	UpdatedByID string   `gorm:"type:varchar(191)"`
	UsersFrom   Users    `gorm:"foreignKey:UpdatedByID"`
	CandidateFile string//buat file table sendiri
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
