package company

import "pasarwarga/models"

type CompanyFormatter struct {
	ID                 string `json:"id"`
	CompanyName        string `json:"companyname"`
	CompanyDescription string `json:"companydescription"`
	// User               CompanyOwner `json:"user"`
}

type CompanyOwner struct {
	User string `json:"user"`
}

func FormatCompany(company models.Company) CompanyFormatter {

	CompanyFormatter := CompanyFormatter{}
	CompanyFormatter.ID = company.ID
	CompanyFormatter.CompanyName = company.CompanyName
	CompanyFormatter.CompanyDescription = company.CompanyDescription

	// user := company.Users

	// GetOwner := CompanyOwner{}
	// GetOwner.User = user.Name

	// CompanyFormatter.User = GetOwner

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
