package Company

import (
	"pasarwarga/models"
)

type CreateCompanyInput struct {
	CompanyName        string `json:"companyname" binding:"required"`
	CompanyDescription string `json:"companydescription" binding:"required"`
	CompanyType string `json:"companytype"`
	LocationID string `json:"locationid"`

	User               models.Users
}

type CompanyFindIDInput struct {
	ID string `uri:"id" binding:"required"`
}

type SearchCompany struct {
	Value string `uri:"companyname"`
}

type CreateCompanyLocInput struct {

	CompanyID string `json:"companyid" binding:"required"`
	LocationID string `json:"locationid" binding:"required"`


}
