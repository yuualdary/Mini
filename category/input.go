package category

type CategoryInput struct {
	CategoryPrefix string `json:"categoryprefIx" binding:"required"`
	CategoryName   string `json:"categoryname" binding:"required"`
}

type CategoryIDInput struct {
	ID string `uri:"id" binding:"required"`
}
