package Company

import (
	"pasarwarga/models"
)

type CompanyFormatter struct {
	ID                 string `json:"id"`
	CompanyName        string `json:"companyname"`
	CompanyDescription string `json:"companydescription"`
	Type CompanyType
	// User               CompanyOwner `json:"user"`
}


type PositionFormatter struct {
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
	//Count int `json:"candidate"`
}


type CompanyType struct {
	ID string `json:"id"`
	CompanyType string `json:"companytype"`
}

type CompanyOwner struct {
	User string `json:"user"`
}

func FormatCompany(company models.Company) CompanyFormatter {

	CompanyFormatter := CompanyFormatter{}
	CompanyFormatter.ID = company.ID
	CompanyFormatter.CompanyName = company.CompanyName
	CompanyFormatter.CompanyDescription = company.CompanyDescription


	category := company.Categories

	CompanyTypeFormatter := CompanyType{}

	CompanyTypeFormatter.ID = category.ID
	CompanyTypeFormatter.CompanyType = category.CategoryName

	// user := company.Users

	// GetOwner := CompanyOwner{}
	// GetOwner.User = user.Name

	// CompanyFormatter.User = GetOwner

	CompanyFormatter.Type = CompanyTypeFormatter

	

	return CompanyFormatter
}

func FormatListCompany(listCompany []models.Company) []CompanyFormatter {

	ListCompanyFormatter := []CompanyFormatter{}

	for _, company := range listCompany {

		CompanyFormatter := FormatCompany(company)
		ListCompanyFormatter = append(ListCompanyFormatter, CompanyFormatter)

	}

	return ListCompanyFormatter
}
