package Company

import (
	"pasarwarga/models"
)

type CreateCompanyInput struct {
	CompanyName        string `json:"companyname" binding:"required"`
	CompanyDescription string `json:"companydescription" binding:"required"`
	User               models.Users
}

type CompanyFindIDInput struct {
	ID string `uri:"id" binding:"required"`
}

type SearchCompany struct {
	Value string `uri:"companyname"`
}
