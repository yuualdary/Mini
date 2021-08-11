package company

import (
	"pasarwarga/models"
)

type CreateCompanyInput struct {
	CompanyName        string `json:"companyname" binding:"required"`
	CompanyDescription string `json:"companydescription" binding:"required"`
	User               models.Users
}

type CompanyFindIDInput struct {
	ID int `uri:"id" binding:"required"`
}
