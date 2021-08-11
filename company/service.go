package company

import (
	"errors"
	"pasarwarga/Users"
	"pasarwarga/models"
)

type Service interface {
	CreateCompany(input CreateCompanyInput) (models.Company, error)
	UpdateCompany(input CreateCompanyInput, inputid CompanyFindIDInput) (models.Company, error)
	ListCompany() ([]models.Company, error)
	DetailCompany(input CompanyFindIDInput) (models.Company, error)
}

type service struct {
	repository     Repository
	UserRepository Users.Repository
}

func NewService(repository Repository, UserRepository Users.Repository) *service {
	return &service{repository, UserRepository}
}

func (s *service) CreateCompany(input CreateCompanyInput) (models.Company, error) {

	CreateCompany := models.Company{}

	CreateCompany.CompanyName = input.CompanyName
	CreateCompany.CompanyDescription = input.CompanyDescription

	Save, err := s.repository.CreateCompany(CreateCompany)

	if err != nil {
		return Save, err
	}

	IsUserGetCompany, err := s.UserRepository.FindUserById(int(input.User.ID))

	if err != nil {
		return IsUserGetCompany.Companies, err
	}

	if IsUserGetCompany.CompanyID != 0 {
		return IsUserGetCompany.Companies, errors.New("You already have a company")
	}

	IsUserGetCompany.CompanyID = int(input.User.ID)
	UpdateUser, err := s.UserRepository.UpdateUser(IsUserGetCompany)

	if err != nil {
		return UpdateUser.Companies, err
	}

	return UpdateUser.Companies, nil

}

func (s *service) UpdateCompany(input CreateCompanyInput, inputid CompanyFindIDInput) (models.Company, error) {

	IsUserGetCompany, err := s.UserRepository.FindUserById(int(input.User.ID))

	if err != nil {
		return IsUserGetCompany.Companies, err
	}

	if IsUserGetCompany.CompanyID != int(input.User.ID) {
		return IsUserGetCompany.Companies, errors.New("Cannot Update Company")
	}
	UpdateCompany, err := s.repository.FindCompanyID(inputid.ID)
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
