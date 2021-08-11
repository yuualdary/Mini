package category

type CategoryInput struct {
	CategoryPrefix string `json:"categoryprefIx" binding:"required"`
	CategoryName   string `json:"categoryname" binding:"required"`
}

type CategoryIDInput struct {
	ID int `uri:"id" binding:"required"`
}
