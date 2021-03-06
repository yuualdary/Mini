package Candidate

import "pasarwarga/models"

type CreateCandidateInput struct {
	PositionID string `json:"positionid"  binding:"required"`
	CategoryID string `json:"categoryid"  binding:"required"`
	User       models.Users
}

type DetailCandidateInput struct {
	ID      string `uri:"id"  binding:"required"`
	Company models.Company
}

type UpdateCandidateInput struct{
	PositionID string `json:"positionid"  binding:"required"`
	CategoryID string `json:"categoryid"  binding:"required"`
	CandidateID string `json:"candidateid"  binding:"required"`
	User       models.Users
}

 
