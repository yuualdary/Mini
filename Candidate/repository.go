package Candidate

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCandidate(candidate models.Candidate) (models.Candidate, error)
	UpdateCandidate(candidate models.Candidate) (models.Candidate, error)
	DetailCandidate(ID string) (models.Candidate, error)
	ListCandidate() ([]models.Candidate, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) CreateCandidate(candidate models.Candidate) (models.Candidate, error) {

	err := r.db.Create(&candidate).Error

	if err != nil {
		return candidate, err
	}
	return candidate, nil
}

func (r *repository) UpdateCandidate(candidate models.Candidate) (models.Candidate, error) {

	err := r.db.Save(&candidate).Error

	if err != nil {
		return candidate, err
	}
	return candidate, nil
}

func (r *repository) DetailCandidate(ID string) (models.Candidate, error) {

	var candidate models.Candidate

	err := r.db.Where("id = ? ", ID).Find(&candidate).Error

	if err != nil {
		return candidate, err
	}

	return candidate, nil

}

func (r *repository) ListCandidate() ([]models.Candidate, error) {

	var candidate []models.Candidate

	err := r.db.Find(&candidate).Error

	if err != nil {
		return []models.Candidate{}, err
	}
	return candidate, nil
}
