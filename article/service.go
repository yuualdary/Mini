package article

import (
	"errors"
	"fmt"
	"pasarwarga/models"
	"strconv"

	"github.com/gosimple/slug"
)


type Service interface {
	CreateArticle(input CreateArticleInput) (models.Article, error)
	UpdateArticle(inputid ArticleDetailInput, inputdata CreateArticleInput)(models.Article, error)
	ListArticle()([]models.Article, error)
	DetailArticle(inputid ArticleDetailInput)(models.Article, error)
	DeleteArticle(inputid ArticleDetailInput)error

}
type service struct {
	repository Repository

}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service)CreateArticle(input CreateArticleInput) (models.Article, error){
	

	
	GetArticle := models.Article{}
	GetArticle.Title = input.Title
	GetArticle.Content = input.Content
	conv := strconv.Itoa(input.CategoryID)

	SlugArticle:= fmt.Sprintf("%s %s", input.Title, conv)
	GetArticle.Slug = slug.Make(SlugArticle)
	GetArticle.CategoryID= input.CategoryID

	SaveArticle, err := s.repository.CreateArticle(GetArticle)

	if err != nil{
		return SaveArticle,err
	}

	return SaveArticle,nil

}

func(s *service)ListArticle()([]models.Article,error){

	ListAllArticle,err:= s.repository.ListArticle()

	if err !=nil{
		return ListAllArticle,err
	}

	return ListAllArticle,nil
}

func (s *service)UpdateArticle(inputid ArticleDetailInput, inputdata CreateArticleInput)(models.Article, error){
	
	GetArticle, err := s.repository.FindArticleID(inputid.ID)
	if err !=nil{
		return GetArticle, err
	}

	GetArticle.Title = inputdata.Title
	GetArticle.Content = inputdata.Content
	fmt.Printf("%d adalah",inputdata.CategoryID)

	GetArticle.CategoryID =inputdata.CategoryID

	conv := strconv.Itoa(inputdata.CategoryID)

	SlugArticle:= fmt.Sprintf("%s %s", inputdata.Title, conv)
	GetArticle.Slug = slug.Make(SlugArticle)
	//fmt.Println(inputdata.CategoryID)

	SaveArticle, err := s.repository.UpdateArticle(GetArticle)

	if err != nil{
		return SaveArticle,err
	}

	return SaveArticle,nil

}

func(s *service)DetailArticle(inputid ArticleDetailInput)(models.Article, error){

	GetArticle,err:=s.repository.DetailArticle(inputid.ID)

	if err != nil{

		return GetArticle,errors.New("Cannot Find Data")
	}

	return GetArticle,nil

}

func(s *service)DeleteArticle(inputid ArticleDetailInput)error{

	
	GetArticle,err:=s.repository.FindArticleID(inputid.ID)

	if err != nil{

		return err
	}

	DeleteArticleErr:=s.repository.DeleteArticle(int(GetArticle.ID))
	
	if DeleteArticleErr!=nil{
		return DeleteArticleErr
	}

	return nil

}

