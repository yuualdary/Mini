package Candidate

import (
	"errors"
	"pasarwarga/Company"
	"pasarwarga/Position"
	"pasarwarga/Users"
	"pasarwarga/models"
)

type Service interface {
	CreateCandidate(input CreateCandidateInput) (models.Candidate, error)
	ListCandidate(inputid DetailCandidateInput) ([]models.Candidate, error)
	UpdateCandidateStatus(inputid DetailCandidateInput, input CreateCandidateInput) (models.Candidate, error)
}

type service struct {
	repository         Repository
	UserRepository     Users.Repository
	PositionRepository Position.Repository
	CompanyRepository  Company.Repository
}

func NewService(repository Repository, UserRepository Users.Repository, PositionRepository Position.Repository, CompanyRepository Company.Repository) *service {
	return &service{repository, UserRepository, PositionRepository, CompanyRepository}
}

func (s *service) CreateCandidate(input CreateCandidateInput) (models.Candidate, error) {

	FindUser, err := s.UserRepository.FindUserById(input.User.ID)

	if err != nil {
		return models.Candidate{}, errors.New("User Not Found")
	}

	CreateCandidate := models.Candidate{}
	CreateCandidate.CategoryID = input.CategoryID
	CreateCandidate.PositionID = input.PositionID
	CreateCandidate.UserID = FindUser.ID
	//buat pdfnya
	SaveCandidate, err := s.repository.CreateCandidate(CreateCandidate)

	if err != nil {
		return SaveCandidate, err
	}

	return SaveCandidate, nil
}

func (s *service) UpdateCandidateStatus(inputid DetailCandidateInput, input CreateCandidateInput) (models.Candidate, error) {

	CreateCandidate := models.Candidate{}
	CreateCandidate.CategoryID = input.CategoryID
	// CreateCandidate.UpdatedByID =
	//buat pdfnya
	SaveCandidate, err := s.repository.CreateCandidate(CreateCandidate)

	if err != nil {
		return SaveCandidate, err
	}

	return SaveCandidate, nil
}

func (s *service) ListCandidate(inputid DetailCandidateInput) ([]models.Candidate, error) {
	//harusnya pakai middleware, buat misahin manya yg owner

	CheckCompanyOwner, err := s.CompanyRepository.FindCompanyOwner(inputid.User.ID)
	//userid 3 - companyid 2
	if err != nil {
		return []models.Candidate{}, err
	}
	CheckPosition, err := s.PositionRepository.DetailPosition(inputid.ID)
	//pos 1 - company id 3
	if err != nil {
		return []models.Candidate{}, err
	}

	if CheckCompanyOwner.ID != CheckPosition.Companies.ID {
		return []models.Candidate{}, errors.New("Not Owner of This Position")
	}

	ListCandidateFromPosition, err := s.repository.ListCandidate(CheckPosition.ID)

	if err != nil {
		return ListCandidateFromPosition, err
	}

	return ListCandidateFromPosition, nil

}
