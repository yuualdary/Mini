package article

type CreateArticleInput struct {
	Title      string `json:"title" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
}

type ArticleDetailInput struct {
	ID int `uri:"id" binding:"required"`
}