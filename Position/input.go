package Position

import "pasarwarga/models"

type CreatePositionInput struct {
	PositionName        string `json:"positionname" binding:"required"`
	PositionDescription string `json:"positiondescription" binding:"required"`
	Users               models.Users
}

type DetailPositionInput struct {
	ID   string `uri:"id" binding:"required"`
	Users models.Users
}
