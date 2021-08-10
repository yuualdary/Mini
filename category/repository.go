package category

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {

	CreateCategory(category models.Category)(models.Category, error)
	FindCategoryID(CategoryID int)(models.Category, error)
	UpdateCategory(category models.Category)(models.Category, error)
	ListCategory()([]models.Category,error)
	DeleteCategory(CategoryID int)error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository)CreateCategory(category models.Category)(models.Category, error){

	err := r.db.Create(&category).Error

	if err != nil{

		return category, err
	}

	return category, nil
	
}

func (r *repository)FindCategoryID(CategoryID int)(models.Category, error){

	var category models.Category

	err := r.db.Where("id = ?",CategoryID).Find(&category).Error

	if err !=nil{
		return category, err
	}

	return category, nil

}

func (r *repository)UpdateCategory(category models.Category)(models.Category, error){

	err := r.db.Save(&category).Error

	if err != nil{

		return category, err
	}

	return category, nil
	
}
func(r *repository)ListCategory()([]models.Category,error){

	var category []models.Category

	err := r.db.Find(&category).Error

	if err !=nil{
		return category, err

	}

	return category, nil


}

func(r *repository)	DeleteCategory(CategoryID int)error{

	var category models.Category
	err := r.db.Where("id = ?", CategoryID).Delete(&category).Error

	if err != nil{
		return err
	}
	return nil
}


