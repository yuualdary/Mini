package Company

import (
	"fmt"
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCompany(company models.Company) (models.Company, error)
	UpdateCompany(company models.Company) (models.Company, error)
	FindCompanyID(CompanyID string) (models.Company, error)
	ListCompany(value string) ([]models.Company, error)
	FindCompanyOwner(UserID string) (models.Company, error)
	//	DeleteCompany(company models.Company) error
	//FindUserWithCompany(ID string) (models.Company, error)
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

	err := r.db.Save(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (r *repository) ListCompany(value string) ([]models.Company, error) {

	var company []models.Company
	fmt.Println(value + "%")
	// err := r.db.Preload("Users").Where("company_name LIKE ?", "%"+value+"%").Find(&company).Error

	err := r.db.Where("company_name LIKE ?", "%"+value+"%").Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (r *repository) FindCompanyID(CompanyID string) (models.Company, error) {

	var company models.Company

	err := r.db.Where("id = ?", CompanyID).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil
}

func (r *repository) FindCompanyOwner(UserID string) (models.Company, error) {

	var company models.Company

	err := r.db.Where("user_id = ?", UserID).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}
