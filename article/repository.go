package article

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {

	CreateArticle(article models.Article)(models.Article, error)
	FindArticleID(ArticleID int)(models.Article, error)
	UpdateArticle(article models.Article)(models.Article, error)
	ListArticle()([]models.Article,error)
	DeleteArticle(ArticleID int) error
	DeleteArticleByCategoryID(CategoryID int) error
	DetailArticle(ArticleID int)(models.Article, error)
	


}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository)CreateArticle(article models.Article)(models.Article, error){

	err := r.db.Create(&article).Error

	if err != nil{

		return article, err
	}

	return article, nil
	
}

func (r *repository)FindArticleID(ArticleID int)(models.Article, error){

	var article models.Article

	err := r.db.Where("id = ?",ArticleID).Find(&article).Error

	if err !=nil{
		return article, err
	}

	return article, nil

}

func (r *repository)DetailArticle(ArticleID int)(models.Article, error){
	var article models.Article

	err := r.db.Preload("Categories").Where("id = ?",ArticleID).Find(&article).Error

	if err !=nil{
		return article, err
	}

	return article, nil

}

func (r *repository)UpdateArticle(category models.Article)(models.Article, error){

	err := r.db.Save(&category).Error

	if err != nil{

		return category, err
	}

	return category, nil
	
}
func(r *repository)ListArticle()([]models.Article,error){

	var article []models.Article

	err := r.db.Preload("Categories").Find(&article).Error

	if err !=nil{
		return article, err

	}

	return article, nil


}

func(r *repository)DeleteArticle(ArticleID int) error{

	var article models.Article

	err:= r.db.Where("id = ?", ArticleID).Delete(&article).Error

	if err != nil{

		return err
	}
	return nil
}

func(r *repository)	DeleteArticleByCategoryID(CategoryID int) error{
	var article models.Article

	err:= r.db.Where("category_id = ?",CategoryID).Delete(&article).Error

	if err !=nil{
		return err
	}
	return nil
}

