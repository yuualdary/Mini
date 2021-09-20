package location

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateLocation(location models.Locations) (models.Locations, error)
	UpdateLocation(location models.Locations) (models.Locations, error)
	ListLocation() ([]models.Locations, error)
	FindLocation(ID string) (models.Locations, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateLocation(location models.Locations) (models.Locations, error) {

	err := r.db.Save(&location).Error

	if err != nil {
		return location, err
	}

	return location, nil
}

func (r *repository) CreateLocation(location models.Locations) (models.Locations, error) {

	err := r.db.Create(&location).Error

	if err != nil {
		return location, err
	}

	return location, nil
}

func (r *repository) ListLocation() ([]models.Locations, error) {

	var location []models.Locations

	err := r.db.Find(&location).Error

	if err != nil {
		return location, err
	}

	return location, nil
}

func (r *repository) FindLocation(ID string) (models.Locations, error) {

	var location models.Locations

	err := r.db.Where("id = ?", ID).Find(&location).Error

	if err != nil {
		return location, err
	}

	return location, nil
}
