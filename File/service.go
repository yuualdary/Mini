package File

import (
	"errors"
	"fmt"
	"pasarwarga/Users"
	"pasarwarga/generatornumber"
	"pasarwarga/models"
)


type Service interface {
	CreateService(userID string, file string) (models.Filepdf, error)
	UpdateFile(inputid DetailFile, file string) (models.Filepdf,error)
	DetailFile(inputid DetailFile) (models.Filepdf,error)
}



type service struct{
	repository Repository
	UserRepository Users.Repository
}

func NewService(repository Repository, 	UserRepository Users.Repository) *service {
	return &service{repository, UserRepository}
}

func (s *service)CreateService(userID string, file string) (models.Filepdf, error){

	FindUserFromID, err := s.UserRepository.FindUserById(userID)

	if err != nil {

		return models.Filepdf{},err
	}


	CreateFile := models.Filepdf{}
	CreateFile.ID =  generatornumber.NewUUID()
	CreateFile.Filename = CreateFile.ID+".pdf"
	CreateFile.UserID = FindUserFromID.ID
	
	CreateUser, err := s.repository.CreateFile(CreateFile)

	if err != nil {
		
		return CreateUser,err
	}

	return CreateUser,nil

}

func (s *service)UpdateFile(inputid DetailFile, file string) (models.Filepdf,error){
	
	FindUserFromID, err := s.UserRepository.FindUserById(inputid.User.ID)

	if err != nil {

		return models.Filepdf{},err
	}

	FindFileID,err := s.repository.FindFile(inputid.PdfFile)

	if err != nil {

		return models.Filepdf{},err
	}


	if FindUserFromID.ID != FindFileID.UserID{

		return models.Filepdf{},errors.New("Forbidden Access")
		//TODO
		//buat validasi kalau yg access dari company tidak forbidden!
		//buat hanya bisa 1 x cv atau bisa banyak upload tapi di set jadi primary (?)
	}

	FindFileID.Filename =FindFileID.ID+".pdf"
	FindFileID.UserID = FindUserFromID.ID
	
	UpdateFile, err := s.repository.UpdateFile(FindFileID)

	if err != nil {

		return UpdateFile,err
	}

	return UpdateFile,nil

}

func (s *service)DetailFile(inputid DetailFile) (models.Filepdf,error){
	fmt.Println(inputid.PdfFile)

	FindUserFromID, err := s.UserRepository.FindUserById(inputid.User.ID)

	if err != nil {

		return models.Filepdf{},err
	}

	FindFileID,err := s.repository.FindFile(inputid.PdfFile)

	if err != nil {

		return models.Filepdf{},err
	}


	if FindUserFromID.ID != FindFileID.UserID{

		return models.Filepdf{},errors.New("Forbidden Access")

	}

	return FindFileID,nil


}


