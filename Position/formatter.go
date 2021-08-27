package Position

import (
	"pasarwarga/models"
)

type PositionFormatter struct {
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
	PositionDescription string      `json:"positiondescription"`
	Company             CompanyName `json:"company"`
}

type CompanyName struct {
	ID          string `json:"id"`
	CompanyName string `json:"companyname"`
}

func FormatDetailPosition(position models.Position) PositionFormatter {

	PositionFormatter := PositionFormatter{}
	PositionFormatter.ID = position.ID
	PositionFormatter.PositionName = position.PositionName

	company := position.Companies

	companyname := CompanyName{}
	companyname.ID = company.ID
	companyname.CompanyName = company.CompanyName

	PositionFormatter.Company = companyname

	return PositionFormatter

}
