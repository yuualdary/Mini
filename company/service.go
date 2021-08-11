package company

import (
	"errors"
	"pasarwarga/generatornumber"
	"pasarwarga/models"
)

type Service interface {
	CreateCompany(input CreateCompanyInput) (models.Company, error)
	UpdateCompany(input CreateCompanyInput, inputid CompanyFindIDInput) (models.Company, error)
	ListCompany() ([]models.Company, error)
	DetailCompany(input CompanyFindIDInput) (models.Company, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateCompany(input CreateCompanyInput) (models.Company, error) {

	IsUserGetCompany, err := s.repository.FindCompanyOwner(input.User.ID)

	if err != nil {
		return IsUserGetCompany, err
	}

	if IsUserGetCompany.Users.ID != "" {
		return IsUserGetCompany, errors.New("You already have a company")
	}

	CreateCompany := models.Company{}
	CreateCompany.ID = generatornumber.NewUUID()
	CreateCompany.CompanyName = input.CompanyName
	CreateCompany.CompanyDescription = input.CompanyDescription
	CreateCompany.UserID = IsUserGetCompany.UserID

	Save, err := s.repository.CreateCompany(CreateCompany)

	if err != nil {
		return Save, err
	}

	return Save, nil

}

func (s *service) UpdateCompany(input CreateCompanyInput, inputid CompanyFindIDInput) (models.Company, error) {

	IsUserGetCompany, err := s.repository.FindCompanyOwner(input.User.ID)

	if err != nil {
		return IsUserGetCompany, err
	}

	if IsUserGetCompany.UserID != input.User.ID {
		return IsUserGetCompany, errors.New("Cannot Update Company")
	}

	UpdateCompany, err := s.repository.FindCompanyID(inputid.ID)
	if err != nil {

		return UpdateCompany, err
	}

	UpdateCompany.CompanyName = input.CompanyName
	UpdateCompany.CompanyDescription = input.CompanyDescription

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
func (s *service) ListCompany() ([]models.Company, error) {

	ListCompany, err := s.repository.ListCompany()

	if err != nil {
		return ListCompany, err
	}

	return ListCompany, nil
}
