package Position

import (
	"pasarwarga/company"
	"pasarwarga/generatornumber"
	"pasarwarga/models"
)

type Service interface {
	CreatePosition(input CreatePositionInput) (models.Position, error)
	UpdatePosition(inputid DetailPositionInput, inputdata CreatePositionInput) (models.Position, error)
	ListPosition() ([]models.Position, error)
	DetailPosition(inputid DetailPositionInput) (models.Position, error)
}

type service struct {
	repository        Repository
	CompanyRepository company.Repository
}

func NewService(repository Repository, CompanyRepository company.Repository) *service {
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