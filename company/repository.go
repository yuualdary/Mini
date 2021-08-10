package company

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCompany(company models.Company) (models.Company, error)
	UpdateCompany(company models.Company) (models.Company, error)
	FindCompanyID(CompanyID int) (models.Company, error)
	ListCompany() ([]models.Company, error)
	DeleteCompany(company models.Company) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) CreateCompany(company models.Company) (models.Company, error) {

	err := r.db.Create(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (r *repository) UpdateCompany(company models.Company) (models.Company, error) {

	err := r.db.Create(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (r *repository) ListCompany() ([]models.Company, error) {

	var company []models.Company
	err := r.db.Create(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (r *repository) FindCompanyID(CompanyID int) (models.Company, error) {

	var company models.Company

	err := r.db.Where("id = ?", CompanyID).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil
}
