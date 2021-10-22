package Users

import (
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	SaveUser(user models.Users) (models.Users, error)
	FindUserEmail(email string) (models.Users, error)
	FindUserById(ID string) (models.Users, error)
	FindUserUnverifiedById(ID string) (models.Users, error)
	UpdateUser(users models.Users) (models.Users, error)
	SaveOTP(otp models.Otps) (models.Otps, error)
	GetListUnregistered(email string)([]models.Users,error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) SaveUser(user models.Users) (models.Users, error) {

	err := r.db.Create(&user).Error

	if err != nil {

		return user, err
	}

	return user, nil

}

func (r *repository) FindUserEmail(email string) (models.Users, error) {

	var User models.Users
	err := r.db.Where("email = ?", email).Find(&User).Error

	if err != nil {
		return User, err
	}

	return User, nil

}
func (r *repository) FindUserById(ID string) (models.Users, error) {

	var User models.Users

	err := r.db.Where("id = ? AND is_verif = ? ", ID,true).Find(&User).Error

	if err != nil {
		return User, err
	}

	return User, err

}

func (r *repository) UpdateUser(users models.Users) (models.Users, error) {

	err := r.db.Save(&users).Error

	if err != nil {

		return users, err
	}

	return users, nil
}
func (r *repository) SaveOTP(otp models.Otps) (models.Otps, error) {

	err := r.db.Create(&otp).Error

	if err != nil {

		return otp, err
	}

	return otp, nil
}

func (r *repository)GetListUnregistered(email string)([]models.Users,error) {
	
	var user []models.Users
	
	err := r.db.Debug().Where("email = ? AND is_verif = ? ", email, false).Find(&user).Error
	
	if err != nil {
		return []models.Users{}, err
	}
	
	return user, nil
}

func (r *repository)FindUserUnverifiedById(ID string) (models.Users, error){

	var user models.Users

	err := r.db.Debug().Where("id = ? AND is_verif = ? ", ID, false).Find(&user).Error

	if err !=nil{
		return models.Users{}, err
	}

	return user, nil

}

