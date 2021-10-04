package File

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)


type Repository interface {
	CreateFile(file models.Filepdf) (models.Filepdf,error)
	UpdateFile(file models.Filepdf) (models.Filepdf,error)
	FindFile(ID string) (models.Filepdf,error)

}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}


func (r *repository)CreateFile(file models.Filepdf) (models.Filepdf,error){

	err := r.db.Create(&file).Error

	if err != nil {
		return file,err
	}

	return file,nil

}

func (r *repository)UpdateFile(file models.Filepdf) (models.Filepdf,error){

	err := r.db.Save(&file).Error

	if err != nil {
		return file,err
	}

	return file,nil

}

func (r *repository)FindFile(ID string) (models.Filepdf,error){

	var file models.Filepdf

	err := r.db.Where("id = ?", ID).Find(&file).Error

	if err != nil {
		return file,err
	}

	return file,nil
}

