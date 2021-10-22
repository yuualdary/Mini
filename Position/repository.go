package Position

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	CreatePosiion(positiion models.Position) (models.Position, error)
	UpdatePosition(position models.Position) (models.Position, error)
	CreateTagPosition(positiontag models.PositionCategory)(models.PositionCategory, error)
	ListPosition() ([]models.Position, error)
	ListCompanyPosition(CompanyID string)([]models.Position,error)
	ListPositionTag(ID string)([]models.PositionCategory,error)
	DetailPosition(ID string) (models.Position, error)
	BookMarkPosition(bookmark models.Bookmarks) (models.Bookmarks,error)
	RemoveBookmark(BookmarkID string) error
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

func (r *repository)CreateTagPosition(positiontag models.PositionCategory)(models.PositionCategory, error){
	
	err := r.db.Create(&positiontag).Error

	if err != nil {
		return positiontag, err
	}
	return positiontag, nil

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

func(r *repository)	ListPositionTag(ID string)([]models.PositionCategory,error){

	var positiion []models.PositionCategory

	err := r.db.Where("position_id = ? ",ID).Find(&positiion).Error

	if err != nil {
		return positiion, err
	}
	return positiion, nil
}

func (r *repository) DetailPosition(ID string) (models.Position, error) {

	var position models.Position

	err := r.db.Preload("Companies").Preload("PositionCategories").Preload("Candidates").Where("id = ? ", ID).Find(&position).Error
	
			
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


func(r *repository)	ListCompanyPosition(CompanyID string)([]models.Position,error){

	var position []models.Position

	err := r.db.Preload("Companies").Where("company_id = ? ", CompanyID).Find(&position).Error

	if err != nil {
		return position,err
	}
	return position,nil

}


func (r *repository)BookMarkPosition(bookmark models.Bookmarks) (models.Bookmarks,error){

	err := r.db.Create(&bookmark).Error

	if err != nil {
		return models.Bookmarks{},err
	}

	return bookmark, nil
}

func (r *repository)RemoveBookmark(BookmarkID string) error {

	var bookmark models.Bookmarks

	err := r.db.Where("PositionID = ? ", BookmarkID).Delete(&bookmark).Error

	if err != nil {

		return err
	}
	return nil
}



