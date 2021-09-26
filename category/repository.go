package category

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreateCategory(category models.Category) (models.Category, error)
	FindCategoryID(CategoryID string) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	ListCategory() ([]models.Category, error)
	ListStatus() ([]models.Category, error)
	ListPositionTag() ([]models.Category, error)
	ListCategoryStatus() ([]models.Category, error)
	//kayaknya bisa pakai string trus read query pakai stringnya, tapi sepertinya berbahaya
	ListJobTypeTag()([]models.Category, error)
	DeleteCategory(CategoryID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {

	err := r.db.Create(&category).Error

	if err != nil {

		return category, err
	}

	return category, nil

}

func (r *repository) FindCategoryID(CategoryID string) (models.Category, error) {

	var category models.Category

	err := r.db.Where("id = ?", CategoryID).Find(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil

}



func (r *repository) ListStatus() ([]models.Category, error) {

	var category []models.Category

	err := r.db.Where("categoryprefix = ?", "STATUS").Error

	if err != nil {
		return []models.Category{}, nil
	}
	return category, nil

}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {

	err := r.db.Save(&category).Error

	if err != nil {

		return category, err
	}

	return category, nil

}
func (r *repository) ListCategory() ([]models.Category, error) {

	var category []models.Category

	err := r.db.Find(&category).Error

	if err != nil {
		return category, err

	}

	return category, nil

}

func (r *repository) ListCategoryStatus() ([]models.Category, error) {

	var category []models.Category

	err := r.db.Where("category_prefix = CANDIDATESTATUS").Find(&category).Error

	if err != nil {
		return category, err

	}

	return category, nil

}

func (r *repository) ListPositionTag() ([]models.Category, error){

	var category []models.Category

	err := r.db.Where("category_prefix = ?","POSITIONTAG").Find(&category).Error

	if err != nil {
		return category, err

	}

	return category, nil

}

func (r *repository) DeleteCategory(CategoryID string) error {

	var category models.Category
	err := r.db.Where("id = ?", CategoryID).Delete(&category).Error

	if err != nil {
		return err
	}
	return nil
}

func(r *repository)ListJobTypeTag()([]models.Category, error){
	
	var category []models.Category

	err := r.db.Where("category_prefix = ?","JOBTYPETAG").Find(&category).Error

	if err != nil {
		return category, err

	}

	return category, nil

}
