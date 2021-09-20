package Position

import (
	"pasarwarga/models"
)

type PositionFormatter struct {
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
	Company             CompanyName `json:"company"`
	Count int `json:"candidate"`
}
//buat gaji,buat validasi input, formatter lain, tambahin RP di gaji

type CompanyName struct {
	ID          string `json:"id"`
	CompanyName string `json:"companyname"`
}
//buat file model
//formatter dicoba
type CandidateCount struct{

	ID string `json:"id"`

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


	if (len(position.Candidates)>0){

		for start:=0 ; start < len(position.Candidates); start++{

			PositionFormatter.Count ++

		}
	}

	
	return PositionFormatter

}

func FormatListCandidate(listcandidate []models.Position) []PositionFormatter{

	ListPositionFormatter := []PositionFormatter{}

		for _, position := range listcandidate{

			CandidateFormatter := FormatDetailPosition(position)//get each position
			ListPositionFormatter = append(ListPositionFormatter, CandidateFormatter)


		}
	

	return ListPositionFormatter

}


