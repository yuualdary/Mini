package Company

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCompany(company models.Company) (models.Company, error)
	UpdateCompany(company models.Company) (models.Company, error)
	FindCompanyID(CompanyID string) (models.Company, error)
	ListCompany(value string, inputjobtag string, inputprovince string, inputcity string)  ([]models.Company, error)
	FindCompanyOwner(CompanyID string) (models.Company, error)
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

func (r *repository) ListCompany(value string, inputjobtag string, inputprovince string, inputcity string)  ([]models.Company, error){

	var company []models.Company
	//fmt.Println(value + "%")
	// err := r.db.Preload("Users").Where("company_name LIKE ?", "%"+value+"%").Find(&company).Error

	err := r.db.Preload("Categories").Where("company_name LIKE ?", "%"+value+"%").Where("location_id LIKE ?","%"+inputcity+"%").
									Where("location_province LIKE ?", "%"+inputprovince+"%").Where("category_id LIKE ?", "%"+inputjobtag+"%").
									Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (r *repository) FindCompanyID(CompanyID string) (models.Company, error) {

	var company models.Company

	err := r.db.Preload("Categories").Where("id = ?", CompanyID).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil
}

func (r *repository) FindCompanyOwner(CompanyID string) (models.Company, error) {

	var company models.Company

	err := r.db.Where("id = ?", CompanyID).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}
