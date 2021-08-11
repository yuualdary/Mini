package location

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateLocation(location models.Locations) (models.Locations, error)
	UpdateLocation(location models.Locations) (models.Locations, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
