package Otp

import (
	"fmt"
	"pasarwarga/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserOtp(UserID string) (models.Otps, error)
	UpdateOTP(otp models.Otps) (models.Otps, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {

	return &repository{db}
}

func (r *repository) UpdateOTP(otp models.Otps) (models.Otps, error) {

	err := r.db.Save(&otp).Error

	if err != nil {

		return otp, err
	}

	return otp, nil
}
func (r *repository) GetUserOtp(UserID string) (models.Otps, error) {

	var Otp models.Otps
	err := r.db.Where("users_id = ?", UserID).Find(&Otp).Error
	fmt.Println(err)
	if err != nil {
		return Otp, err
	}

	return Otp, err
}
