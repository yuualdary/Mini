package Position

import (
	"errors"
	"pasarwarga/Company"
	"pasarwarga/generatornumber"
	"pasarwarga/models"
)

type Service interface {
	CreatePosition(input CreatePositionInput) (models.Position, error)
	UpdatePosition(inputid DetailPositionInput, inputdata CreatePositionInput) (models.Position, error)
	ListPosition() ([]models.Position, error)
	DetailPosition(inputid DetailPositionInput) (models.Position, error)
	DeletePosition(input DetailPositionInput) error
}

type service struct {
	repository        Repository
	CompanyRepository Company.Repository
}

func NewService(repository Repository, CompanyRepository Company.Repository) *service {
	return &service{repository, CompanyRepository}
}

func (s *service) CreatePosition(input CreatePositionInput) (models.Position, error) {

	FindUser, err := s.CompanyRepository.FindCompanyOwner(input.Users.ID)

	if FindUser.UserID != input.Users.ID {
		return models.Position{}, err
	}

	Create := models.Position{}
	Create.ID = generatornumber.NewUUID()
	Create.PositionName = input.PositionName
	Create.PositionDescription = input.PositionDescription
	Create.PositionFee = input.PositionFee
	Create.PositionLength = input.PositionLength
	Create.PositionRequirement = input.PositionRequirement
	Create.CompanyID = FindUser.ID
	NewPosition, err := s.repository.CreatePosiion(Create)

	if err != nil {
		return NewPosition, err
	}

	return NewPosition, nil
}

func (s *service) UpdatePosition(inputid DetailPositionInput, inputdata CreatePositionInput) (models.Position, error) {

	FindDetail, err := s.repository.DetailPosition(inputid.ID)

	if err != nil {
		return FindDetail, err
	}

	FindDetail.PositionName = inputdata.PositionName
	FindDetail.PositionDescription = inputdata.PositionDescription
	FindDetail.PositionLength = inputdata.PositionLength
	FindDetail.PositionFee = inputdata.PositionFee
	FindDetail.PositionRequirement = inputdata.PositionRequirement
	NewPosition, err := s.repository.CreatePosiion(FindDetail)

	if err != nil {
		return NewPosition, err
	}

	return NewPosition, nil
}

func (s *service) ListPosition() ([]models.Position, error) {

	GetList, err := s.repository.ListPosition()

	if err != nil {
		return GetList, err
	}

	return GetList, nil

}
func (s *service) DetailPosition(inputid DetailPositionInput) (models.Position, error) {

	FindPosition, err := s.repository.DetailPosition(inputid.ID)

	if err != nil {
		return FindPosition, err
	}

	return FindPosition, nil
}

func (s *service) DeletePosition(input DetailPositionInput) error {

	FindDetail, err := s.repository.DetailPosition(input.ID)

	if err != nil {
		return err
	}

	FindOwner, err := s.CompanyRepository.FindCompanyOwner(input.Users.ID)

	if err != nil {
		return err
	}

	if FindDetail.CompanyID != FindOwner.UserID {

		return errors.New("Not an owner")
	}
	return nil

}
