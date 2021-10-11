package Candidate

import "pasarwarga/models"

type DetailCandidateFormatter struct {
	ID              string                  `json:"id"`
	CandidateStatus string                  `json:"status"`
	Position     PositionDetailFormatter `json:"position"`
	User UserCandidate `json:"user"`
	
}


type DetailUserApplicationFormatter struct{
	ID string  `json:"id"`
	CandidateStatus string `json:"status"`
	Position     PositionDetailFormatter `json:"position"`
	CompanyName string `json:"companyname"`


}

type UserCandidate struct {
	ID           string `json:"id"`
	UserName string `json:"username"`
}

type PositionDetailFormatter struct {
	ID           string `json:"id"`
	PositionName string `json:"positionname"`
}

func FormatCandidateDetail(candidate models.Candidate) DetailCandidateFormatter {

	CandidateFormatter := DetailCandidateFormatter{}
	CandidateFormatter.ID = candidate.ID
	CandidateFormatter.CandidateStatus = candidate.Categories.CategoryName

	user := candidate.Users

	getuser := UserCandidate{}
	getuser.ID = user.ID
	getuser.UserName = user.Name

	CandidateFormatter.User = getuser

	position := candidate.Positions

	GetPosition := PositionDetailFormatter{}
	GetPosition.ID = position.ID
	GetPosition.PositionName = position.PositionName

	CandidateFormatter.Position = GetPosition

	return CandidateFormatter
	
}

func FormatDetailApplication(candidate models.Candidate, company []models.Company) DetailUserApplicationFormatter {

	ListApplicationFormatter := DetailUserApplicationFormatter{}

	ListApplicationFormatter.ID = candidate.ID
	ListApplicationFormatter.CandidateStatus = candidate.Categories.CategoryName

	GetPosition := PositionDetailFormatter{}

	GetPosition.ID = candidate.Positions.ID
	GetPosition.PositionName = candidate.Positions.PositionName

	ListApplicationFormatter.Position = GetPosition


	for _, findcompany := range company{

		if findcompany.ID == candidate.Positions.CompanyID{

			ListApplicationFormatter.CompanyName = findcompany.CompanyName
		}
	}


	return ListApplicationFormatter


}

func FormatListApplication(candidate []models.Candidate, company []models.Company) []DetailUserApplicationFormatter{

	ListApplicationFormatter := []DetailUserApplicationFormatter{}

		for _, listcandidate := range candidate{

			DetailApplicationFormatter := FormatDetailApplication(listcandidate,company)//get each position
			ListApplicationFormatter = append(ListApplicationFormatter, DetailApplicationFormatter)


		}
	

	return ListApplicationFormatter

}