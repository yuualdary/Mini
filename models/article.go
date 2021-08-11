package models

type Article struct {
	ID string `gorm:"primary_key"`

	Title      string
	Slug       string
	CategoryID string
	Categories Category `gorm:"foreignKey:CategoryID"`
	Content    string
}

type ArticleToCategory struct {
}
