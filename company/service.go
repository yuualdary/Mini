package Company

import (
	"errors"
	"fmt"

	"pasarwarga/Users"
	"pasarwarga/generatornumber"
	"pasarwarga/models"

	"github.com/gosimple/slug"
)
type Service interface {
	CreateCompany(input CreateCompanyInput) (models.Company, error)
//	CreateCompanyLocation(input CreateCompanyLocInput)()
	UpdateCompany(input CreateCompanyInput, inputid CompanyFindIDInput) (models.Company, error)
	ListCompany(value string) ([]models.Company, error)
	DetailCompany(input CompanyFindIDInput) (models.Company, error)
	CompanyOwner(ID string) (models.Company, error)
}

type service struct {
	repository     Repository
	UserRepository Users.Repository
}

func NewService(repository Repository, UserRepository Users.Repository) *service {
	return &service{repository, UserRepository}
}

func (s *service) CreateCompany(input CreateCompanyInput) (models.Company, error) {
	IsUserGetCompany, err := s.repository.FindCompanyOwner(input.User.ID)

	if err != nil {
		return IsUserGetCompany, err
	}

	if IsUserGetCompany.UserID != "" {
		return IsUserGetCompany, errors.New("You already have a company")
	}

	// IsUserExist, err := s.UserRepository.FindUserById(input.User.ID)

	// if IsUserExist.ID != input.User.ID {

	// 	return models.Company{}, errors.New("User Not Registered")
	// }

	CreateCompany := models.Company{}
	CreateCompany.ID = generatornumber.NewUUID()
	CreateCompany.CompanyName = input.CompanyName
	CreateCompany.CompanyDescription = input.CompanyDescription
	CreateCompany.UserID = input.User.ID
	CreateCompany.CategoryID = input.CompanyType
	CreateCompany.LocationID = input.LocationID

	CreateCompany.CompanySlug = slug.Make(input.CompanyName)


	Save, err := s.repository.CreateCompany(CreateCompany)

	if err != nil {
		return Save, err
	}

	return Save, nil

}

func (s *service) UpdateCompany(input CreateCompanyInput, inputid CompanyFindIDInput) (models.Company, error) {

	IsUserGetCompany, err := s.repository.FindCompanyOwner(inputid.ID)

	if err != nil {
		return IsUserGetCompany, err
	}
	
	fmt.Println(IsUserGetCompany.ID)
	fmt.Println(inputid.ID)
//positon detail add type company done
//status candidate 
//company list position done -formatter 
//fixing location
//TODO 
	if IsUserGetCompany.UserID != input.User.ID {
		return IsUserGetCompany, errors.New("Forbidden Access")
	}

	UpdateCompany, err := s.repository.FindCompanyID(inputid.ID)
	if err != nil {

		return UpdateCompany, err
	}

	UpdateCompany.CompanyName = input.CompanyName
	UpdateCompany.CompanyDescription = input.CompanyDescription
	UpdateCompany.CompanySlug = slug.Make(input.CompanyName)
	UpdateCompany.CategoryID = input.CompanyType
	UpdateCompany.LocationID = input.LocationID

	Save, err := s.repository.UpdateCompany(UpdateCompany)

	if err != nil {
		return Save, err
	}

	return Save, nil
}

func (s *service) DetailCompany(inputid CompanyFindIDInput) (models.Company, error) {

	FindID, err := s.repository.FindCompanyID(inputid.ID)

	if err != nil {
		return FindID, err
	}

	return FindID, nil

}
func (s *service) ListCompany(value string) ([]models.Company, error) {

	ListCompany, err := s.repository.ListCompany(value)

	if err != nil {
		return ListCompany, err
	}

	return ListCompany, nil
}
func (s *service) CompanyOwner(ID string) (models.Company, error) {

	FindOwner, err := s.repository.FindCompanyOwner(ID)

	if err != nil {
		return FindOwner, err
	}
	return FindOwner, nil
}
