package Company

import (
	"encoding/json"
	"pasarwarga/apiformat"
	"pasarwarga/fetch"
	"pasarwarga/models"
	"strconv"
	"strings"
)

type CompanyFormatter struct {
	ID                 string `json:"id"`
	CompanyName        string `json:"companyname"`
	CompanyDescription string `json:"companydescription"`
	LocationData LocationDetail
	Type CompanyType
	// User               CompanyOwner `json:"user"`
}


type PositionFormatter struct {
	ID                  string      `json:"id"`
	PositionName        string      `json:"positionname"`
	//Count int `json:"candidate"`
}
type LocationDetail struct {
	Id int `json:"id"`
	Province string `json:"province"`
	City string `json:"city"`
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
	// LocationFormatter := []LocationDetail{}

	
	GetDetail := strconv.Itoa(company.LocationID)

	GetLocation, err := fetch.LocationGet("api/daerahindonesia/kota/" + GetDetail)
	
	if err != nil {
		return CompanyFormatter
	}

	var lokasi apiformat.GetFormatLokasiKota
	err = json.Unmarshal(GetLocation, &lokasi)

	if err != nil {
		return CompanyFormatter
	}
	DetailLocationFormatter := LocationDetail{}
	DetailLocationFormatter.Id = lokasi.Id
	DetailLocationFormatter.Province = lokasi.Id_provinsi
	city := strings.Replace(lokasi.Name,"Kabupaten ", "", -1)
	DetailLocationFormatter.City = city



	CompanyFormatter.Type = CompanyTypeFormatter
	CompanyFormatter.LocationData = DetailLocationFormatter

	

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
