package location

type CreateLocationInput struct {
	LocationCity string `json:"locationcity" binding:"required"`
}

type DetailLocationInput struct {
	ID string `uri:"id" binding:"required"`
}
