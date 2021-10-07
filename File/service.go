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
	ListAll(ID string)([]models.Filepdf,error)
	UpdatePrimary(NewID string, UserID string)([]models.Filepdf,error)
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
	CreateFile.IsPrimary = true	
	CreateUser, err := s.repository.CreateFile(CreateFile)

	//Update Primary to unprimary

	if err != nil {
		
		return CreateUser,err
	}

	_,err = s.UpdatePrimary(CreateFile.ID ,FindUserFromID.ID)

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
		//isprimary done
		//check pdf dan ukuran file kalau bisa
		//list dan validasi max 3x file
		//update file set primary, delete file 
		//buat validasi kalau yg access dari company tidak forbidden!
		//buat hanya bisa 1 x cv atau bisa banyak upload tapi di set jadi primary (?) done
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

func(s *service)ListAll(ID string)([]models.Filepdf,error){

	GetList, err := s.repository.ListAllFile(ID)

	if err != nil {

		return GetList,err
	}

	return GetList,nil
}


func(s *service)UpdatePrimary(NewID string, UserID string)([]models.Filepdf,error){

	GetList, err := s.repository.ListAllFile(UserID)

	if err != nil {

		return GetList,err
	}

	for _, ListFile := range GetList{

		if ListFile.ID != NewID && ListFile.IsPrimary{

			ListFile.IsPrimary = false
			_,err := s.repository.UpdateFile(ListFile)
			
			if err != nil {

				return []models.Filepdf{}, err
			}


		}


	}

	return GetList,nil

}


