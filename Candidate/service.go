package Candidate

import (
	"errors"
	"pasarwarga/Users"
	"pasarwarga/models"
)

type Service interface {
	CreateCandidate(input CreateCandidateInput) (models.Candidate, error)
	DetailCandidate(inputid DetailCandidateInput) (models.Candidate, error)
	UpdateCandidateStatus(inputid DetailCandidateInput, input CreateCandidateInput) (models.Candidate, error)
}

type service struct {
	repository     Repository
	UserRepository Users.Repository
}

func NewService(repository Repository, UserRepository Users.Repository) *service {
	return &service{repository, UserRepository}
}

func (s *service) CreateCandidate(input CreateCandidateInput) (models.Candidate, error) {

	FindUser, err := s.UserRepository.FindUserById(input.UserID)

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
	//buat pdfnya
	SaveCandidate, err := s.repository.CreateCandidate(CreateCandidate)

	if err != nil {
		return SaveCandidate, err
	}

	return SaveCandidate, nil
}
