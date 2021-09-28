package Position

import (
	"pasarwarga/models"
)



type CompanyFormatter struct {
	ID                 string `json:"id"`
	CompanyName        string `json:"companyname"`
	CompanyDescription string `json:"companydescription"`
	SubPosition []SubPositionFormatter `json:"companyposition"`

	// User               CompanyOwner `json:"user"`

}

type SubPositionFormatter struct {
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
}

type PositionFormatter struct {
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
	Company             CompanyName `json:"company"`
	Count int `json:"candidate"`
}

type DetailPositionFormatter struct{
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
	PositionDescription string `json:"description"`
	PositionFee int `json:"fee"`
	PositionRequirement string `json:"requirement"`
	PositionLength int `json:"length"`
	Company             CompanyName `json:"company"`
	CompanyType CompanyType `json:"type"`
	Tag	[]PositionTag `json:"tag"`
	Count int `json:"candidate"`

}
//buat gaji,buat validasi input, formatter lain, tambahin RP di gaji


type PositionTag struct{
	ID          string `json:"id"`
	PositionTag string `json:"tag"`
}
type CompanyType struct{
	ID          string `json:"id"`
	CompanyType string `json:"companytype"`
}

type CompanyName struct {
	ID          string `json:"id"`
	CompanyName string `json:"companyname"`
}
//buat file model
//formatter dicoba
type CandidateCount struct{

	ID string `json:"id"`

}


func FormatCompany(position []models.Position) CompanyFormatter {

	CompanyFormatter := CompanyFormatter{}

	
	for i:= 0 ; i < 1; i++{

		CompanyFormatter.ID = position[0].Companies.ID
		CompanyFormatter.CompanyName =   position[0].Companies.CompanyName
		CompanyFormatter.CompanyDescription = position[0].Companies.CompanyDescription


	}

	PositionCompany := []SubPositionFormatter{}


	for _, listjob := range position{

		SubPos := SubPositionFormatter{}
		SubPos.ID = listjob.ID
		SubPos.PositionName = listjob.PositionName

		PositionCompany = append(PositionCompany,SubPos)
				
	}

	CompanyFormatter.SubPosition = PositionCompany

	return CompanyFormatter
}

func FormatDetailPosition(position models.Position, category []models.Category) DetailPositionFormatter {

	DetailPositionFormatter := DetailPositionFormatter{}
	DetailPositionFormatter.ID = position.ID
	DetailPositionFormatter.PositionName = position.PositionName
	DetailPositionFormatter.PositionDescription = position.PositionDescription
	DetailPositionFormatter.PositionFee = position.PositionFee
	DetailPositionFormatter.PositionLength = position.PositionLength

	company := position.Companies

	companyname := CompanyName{}
	companyname.ID = company.ID
	companyname.CompanyName = company.CompanyName
	DetailPositionFormatter.Company = companyname



	if position.CompanyID != "" {

		
		for _, listjobtype := range category{
			
			if listjobtype.ID == company.CategoryID{
			//	fmt.Println(listjobtype.CategoryName)

			//	fmt.Println(company.CategoryID)
					companytype := CompanyType{}
					companytype.ID = listjobtype.ID
					companytype.CompanyType = listjobtype.CategoryName
					DetailPositionFormatter.CompanyType = companytype
					
			}
		}
	}



	if (len(position.Candidates)>0){

		for start:=0 ; start < len(position.Candidates); start++{

			DetailPositionFormatter.Count ++

		}
	}


	tagname := []PositionTag{}


		//logic loop 2 x, loop pertama itu get list [catid, posid] loop ke 2 get categoryall [catid] loop pertama 
		//loop [catid,posid] 1 ke range di category, jadi dia akan cek loop dari list category yang id nya sama seperti catid loop pertama
		//dan juga masukkan pos id yang sekarang untuk cocoking posisi yang sesuai
	if (len(position.PositionCategories)>0){

			for _, listpositiontag := range position.PositionCategories{

				for _, getcategory := range category{

					if listpositiontag.PositionID == position.ID && getcategory.ID == listpositiontag.CategoryID {

						postagformatter := PositionTag{}
						postagformatter.ID = getcategory.ID
						postagformatter.PositionTag = getcategory.CategoryName
						tagname = append(tagname, postagformatter)
				}

			}
		
		}
	}
	DetailPositionFormatter.Tag = tagname
	

	//buat type job [se,design]
	return DetailPositionFormatter

}

func FormatListPosition(position models.Position) PositionFormatter {

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

			CandidateFormatter := FormatListPosition(position)//get each position
			ListPositionFormatter = append(ListPositionFormatter, CandidateFormatter)


		}
	

	return ListPositionFormatter

}


