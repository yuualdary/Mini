package location

import (
	"pasarwarga/generatornumber"
	"pasarwarga/models"
)

type Service interface {
	CreateLocation(input CreateLocationInput) (models.Locations, error)
	UpdateLocation(input DetailLocationInput, inputdata CreateLocationInput) (models.Locations, error)
	DetailLocation(ID string)(models.Locations,error)
	LocationList() ([]models.Locations, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateLocation(input CreateLocationInput) (models.Locations, error) {

	NewLocation := models.Locations{}
	NewLocation.ID = generatornumber.NewUUID()
	NewLocation.LocationCity = input.LocationCity

	SaveLocation, err := s.repository.CreateLocation(NewLocation)

	if err != nil {
		return SaveLocation, err
	}

	return SaveLocation, nil

}


func (s *service)DetailLocation(ID string)(models.Locations,error){

	DetailLocation, err := s.repository.FindLocation(ID)

	if err != nil {

		return DetailLocation,err
	}

	return DetailLocation, nil
}


func (s *service) UpdateLocation(input DetailLocationInput, inputdata CreateLocationInput) (models.Locations, error) {

	GetLocation, err := s.repository.FindLocation(input.ID)

	GetLocation.LocationCity = inputdata.LocationCity

	UpdateLocation, err := s.repository.UpdateLocation(GetLocation)

	if err != nil {
		return UpdateLocation, err
	}

	return UpdateLocation, nil

}
func (s *service) LocationList() ([]models.Locations, error) {

	List, err := s.repository.ListLocation()
	if err != nil {
		return List, err
	}

	return List, nil

}
