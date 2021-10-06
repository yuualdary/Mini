package Candidate

import (
	"errors"
	"fmt"
	"pasarwarga/Company"
	"pasarwarga/Position"
	"pasarwarga/Users"
	"pasarwarga/generatornumber"
	"pasarwarga/models"
)

type Service interface {
	CreateCandidate(input CreateCandidateInput) (models.Candidate, error)
	ListCandidate(inputid DetailCandidateInput) ([]models.Candidate, error)
	UpdateCandidateStatus(inputid DetailCandidateInput, input CreateCandidateInput) (models.Candidate, error)
	CheckOwnerValidation(ID string, CompanyID string) (bool, error)
	DetailCandidate(inputid DetailCandidateInput) (models.Candidate,error)
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
	//check apakah current id sudah pernah dipakai

	CheckCandidate, _ := s.repository.DetailCandidatePosition(input.PositionID)

	fmt.Println(FindUser.ID)
	fmt.Println(CheckCandidate.UserID)
	if FindUser.ID == CheckCandidate.UserID && CheckCandidate.PositionID == input.PositionID && CheckCandidate.UserID != "" {

		return models.Candidate{}, errors.New("You've applied in this position")
	}

	CreateCandidate := models.Candidate{}
	CreateCandidate.ID = generatornumber.NewUUID()
	CreateCandidate.CategoryID = input.CategoryID
	CreateCandidate.PositionID = input.PositionID
	CreateCandidate.UserID = FindUser.ID
	CreateCandidate.UpdatedByID = FindUser.ID
	//buat validasi gabisa apply 2x done;
	//detail list cabdidate, formatter total
	//buat pdfnya
	SaveCandidate, err := s.repository.CreateCandidate(CreateCandidate)

	if err != nil {
		return SaveCandidate, err
	}

	return SaveCandidate, nil
}

func (s *service) UpdateCandidateStatus(inputid DetailCandidateInput, input CreateCandidateInput) (models.Candidate, error) {

	
	FindCandidate, err := s.repository.DetailCandidate(inputid.ID)
	
	if err != nil {

		return models.Candidate{},err
	} 

	
	CheckOwner,err := s.CheckOwnerValidation(FindCandidate.PositionID, inputid.ID)

	if err != nil {

		return models.Candidate{},err
	}

	if !CheckOwner{

		return models.Candidate{},errors.New("Forbidden Access")
	}

	
	FindCandidate.CategoryID = input.CategoryID
	FindCandidate.UpdatedByID = input.User.ID
	//buat pdfnya
	SaveCandidate, err := s.repository.UpdateCandidate(FindCandidate)

	if err != nil {
		return SaveCandidate, err
	}

	return SaveCandidate, nil
}

func (s *service)DetailCandidate(inputid DetailCandidateInput) (models.Candidate,error){



	FindCandidate, err := s.repository.DetailCandidate(inputid.ID)
	
	if err != nil {

		return models.Candidate{},err
	}

	CheckOwner, err := s.CheckOwnerValidation(FindCandidate.Positions.ID,inputid.Company.ID)

	if err != nil {
		return models.Candidate{}, err
	}

	if !CheckOwner {
		return models.Candidate{}, errors.New("Forbidden Access")
	}


	return FindCandidate,nil
}


func (s *service) ListCandidate(inputid DetailCandidateInput) ([]models.Candidate, error) {
	//harusnya pakai middleware, buat misahin manya yg owner
	FindCandidate, err := s.repository.DetailCandidate(inputid.ID)
	
	if err != nil {

		return []models.Candidate{},err
	} 

	CheckOwner, err := s.CheckOwnerValidation(FindCandidate.PositionID,inputid.Company.ID)

	if err != nil {
		return []models.Candidate{}, err
	}

	if !CheckOwner {
		return []models.Candidate{}, errors.New("Forbidden Access")
	}

	ListCandidateFromPosition, err := s.repository.ListCandidate(inputid.ID)

	if err != nil {
		return ListCandidateFromPosition, err

	}

	return ListCandidateFromPosition, nil

}

func (s *service) CheckOwnerValidation(ID string, CompanyID string) (bool, error) {
	//2 //3
	CheckPosition, err := s.PositionRepository.DetailPosition(ID)
	//get companyid positionid

	if err != nil {
		return false, err
	}
	//2 //4
	if CheckPosition.CompanyID != CompanyID {
		return false, err
	}

	return true, nil

}
