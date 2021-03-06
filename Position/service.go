package Position

import (
	"errors"
	"pasarwarga/Company"
	"pasarwarga/generatornumber"
	"pasarwarga/models"

	"github.com/gosimple/slug"
)

type Service interface {
	CreatePosition(input CreatePositionInput) (models.Position, error)
	UpdatePosition(inputid DetailPositionInput, inputdata CreatePositionInput) (models.Position, error)
	CreateTagPosition(inputid DetailPositionInput, inputtag CreateTagPosition) (models.PositionCategory,error)
	ListPosition(positionname string, inputjobtag string, inputprovince string, inputcity string) ([]models.Position, error)
	ListCompanyPosition(CompanyID string) ([]models.Position, error)
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

func (s *service)CreateTagPosition(inputid DetailPositionInput, inputtag CreateTagPosition) (models.PositionCategory,error){

	FindDetail, err := s.repository.DetailPosition(inputid.ID)
	//validasi gabisa 2x add tag
	//sebenarnya udah bisa auto karna gabisa 2 PK tapi dibuat validasi agar bagus
	//formatter detail position dengan tag

	if err != nil {
		return models.PositionCategory{}, err
	}


	FindPosition,err := s.repository.ListPositionTag(FindDetail.ID)

	if err != nil{
		return models.PositionCategory{}, err
	}

	// if FindPosition.PositionID != FindDetail.ID{

	// 	return models.PositionCategory{},errors.New("Cannot Update Another Position")
	// }

	for _, listposition := range FindPosition{


	if listposition.CategoryID == inputtag.ID{

		return models.PositionCategory{}, errors.New("Cannot add same category")
	}
}



	FindUser, err := s.CompanyRepository.FindCompanyOwner(inputid.Users.ID)

	if FindUser.UserID != inputid.Users.ID {
		return models.PositionCategory{}, err
	}
	
	
	CreateTag := models.PositionCategory{}
	CreateTag.CategoryID = inputtag.ID
	CreateTag.PositionID = FindDetail.ID

	SaveTag,err := s.repository.CreateTagPosition(CreateTag)

	if err !=nil{

		return SaveTag,err
	}

	
	return SaveTag,nil

}


func (s *service) CreatePosition(input CreatePositionInput) (models.Position, error) {

	FindUser, err := s.CompanyRepository.FindCompanyOwner(input.Users.ID)

	if err != nil {
		return models.Position{}, err
	}
	if FindUser.UserID != input.Users.ID {
		return models.Position{}, errors.New("Forbidden Access")
	}

	Create := models.Position{}
	Create.ID = generatornumber.NewUUID()
	Create.PositionName = input.PositionName
	Create.PositionDescription = input.PositionDescription
	Create.PositionFee = input.PositionFee
	Create.PositionLength = input.PositionLength
	Create.PositionRequirement = input.PositionRequirement
	Create.CompanyID = FindUser.ID
	Create.PositionSlug = slug.Make(input.PositionName)

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

	FindUser, err := s.CompanyRepository.FindCompanyOwner(inputid.Users.ID)

	if err != nil {
		return models.Position{}, err
	}

	if FindUser.UserID != inputid.Users.ID {
		return models.Position{}, errors.New("Forbidden Access")
	}


	FindDetail.PositionName = inputdata.PositionName
	FindDetail.PositionDescription = inputdata.PositionDescription
	FindDetail.PositionLength = inputdata.PositionLength
	FindDetail.PositionFee = inputdata.PositionFee
	FindDetail.PositionRequirement = inputdata.PositionRequirement
	FindDetail.PositionSlug = slug.Make(inputdata.PositionName)
	Update, err := s.repository.UpdatePosition(FindDetail)

	if err != nil {
		return Update, err
	}

	return Update, nil
}

func (s *service) ListPosition(positionname string, inputjobtag string, inputprovince string, inputcity string) ([]models.Position, error) {

	GetList, err := s.repository.ListPosition(positionname,inputjobtag,inputprovince,inputcity)

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

	//Check Candidate Exist

	if len(FindDetail.Candidates) != 0 {
		return errors.New("Cannot delete because you already have candidate")
	}

	return nil

}


func (s *service) ListCompanyPosition(CompanyID string) ([]models.Position, error){

	ListPosition,err := s.repository.ListCompanyPosition(CompanyID)

	if err != nil{

		return ListPosition,err
	}

	return ListPosition,nil

}
