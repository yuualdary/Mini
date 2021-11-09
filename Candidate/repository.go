package Candidate

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCandidate(candidate models.Candidate) (models.Candidate, error)
	UpdateCandidate(candidate models.Candidate) (models.Candidate, error)
	DetailCandidatePosition(ID string) (models.Candidate, error)
	DetailCandidate(ID string)(models.Candidate,error)
	ListCandidate(positionid string) ([]models.Candidate, error)
	ListUserApplication(UserID string,value string)([]models.Candidate,error)
	
	
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


func (r *repository) DetailCandidatePosition(ID string) (models.Candidate, error) {

	var candidate models.Candidate

	err := r.db.Preload("Users").Where("position_id = ? ", ID).Find(&candidate).Error

	if err != nil {
		return candidate, err
	}

	return candidate, nil

}
func (r *repository)DetailCandidate(ID string)(models.Candidate,error){

	var candidate models.Candidate

	err := r.db.Preload("Candidates").Preload("Users").Where("id = ? ", ID).Find(&candidate).Error

	if err != nil {
		return candidate, err
	}

	return candidate, nil

}

func (r *repository) ListCandidate(positioid string) ([]models.Candidate, error) {

	var candidate []models.Candidate

	err := r.db.Preload("Positions").Where("position_id = ?", positioid).Find(&candidate).Error

	if err != nil {
		return []models.Candidate{}, err
	}
	return candidate, nil
}


func (r *repository)ListUserApplication(UserID string,value string)([]models.Candidate,error){

	var candidate []models.Candidate

	err := r.db.Joins("Categories").Preload("Positions").Where("user_id = ? ", UserID).Where("categories.category_prefix = ?","CANDIDATESTATUS").Where("categories.category_name LIKE ?", "%"+value+"%").Find(&candidate).Error

	if err != nil {
		return candidate,err
	}

	return candidate, nil

}
