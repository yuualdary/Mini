package Position

import "pasarwarga/models"

type CreatePositionInput struct {
	PositionName        string `json:"positionname" binding:"required"`
	PositionDescription string `json:"positiondescription" binding:"required"`
	PositionFee         int    `json:"positionfee" binding:"required"`
	PositionLength      int    `json:"positionlength" binding:"required"`
	PositionRequirement string `json:"positionrequirement" binding:"required"`
	Users               models.Users
}

type DetailPositionInput struct {
	ID    string `uri:"id" binding:"required"`
	Users models.Users
}
type CreateTagPosition struct {
	ID    string `json:"id" binding:"required"`
	Users               models.Users
}
