package Position

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreatePosiion(positiion models.Position) (models.Position, error)
	UpdatePosition(position models.Position) (models.Position, error)
	ListPosition() ([]models.Position, error)
	DetailPosition(ID string) (models.Position, error)
	DeletePosition(ID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}
func (r *repository) CreatePosiion(positiion models.Position) (models.Position, error) {

	err := r.db.Create(&positiion).Error

	if err != nil {
		return positiion, err
	}
	return positiion, nil
}

func (r *repository) UpdatePosition(positiion models.Position) (models.Position, error) {

	err := r.db.Save(&positiion).Error

	if err != nil {
		return positiion, err
	}
	return positiion, nil
}

func (r *repository) ListPosition() ([]models.Position, error) {

	var position []models.Position

	err := r.db.Preload("Companies").Preload("Candidates").Find(&position).Error

	if err != nil {
		return []models.Position{}, err
	}

	return position, nil
}
func (r *repository) DetailPosition(ID string) (models.Position, error) {

	var position models.Position

	err := r.db.Preload("Companies").Where("id = ?", ID).Find(&position).Error

	if err != nil {
		return position, err
	}

	return position, nil
}

func (r *repository) DeletePosition(ID string) error {

	var position models.Position

	err := r.db.Where("id = ?", ID).Delete(&position).Error

	if err != nil {
		return err
	}
	return nil
}
