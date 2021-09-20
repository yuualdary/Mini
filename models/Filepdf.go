package models

type Filepdf struct {
	ID       string `gorm:"primary_key"`
	Filename string
	UserID   string `gorm:"type:varchar(191)"`
	Users    Users  `gorm:"foreignKey:UserID"`
}
