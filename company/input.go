package Company

import (
	"pasarwarga/models"
)

type CreateCompanyInput struct {
	CompanyName        string `json:"companyname" binding:"required"`
	CompanyDescription string `json:"companydescription" binding:"required"`
	CompanyType string `json:"companytype" binding:"required"`
	LocationID int `json:"locationid" binding:"required"`
	LocationProvince int `json:"locationprovince" binding:"required"`

	User               models.Users
}

type CompanyFindIDInput struct {
	ID string `uri:"id" binding:"required"`
}

// type SearchCompany struct {
// 	Value string `uri:"companyname"`
// 	locationprovince int `uri:"locationprovince"`
// 	locationcity int `uri:"locationcity"`
// }

type CreateCompanyLocInput struct {

	CompanyID string `json:"companyid" binding:"required"`
	LocationID string `json:"locationid" binding:"required"`


}
