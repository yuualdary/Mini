package Otp

import (
	"errors"
	"fmt"
	"math/rand"
	"pasarwarga/Users"
	"pasarwarga/models"
	"strconv"
	"time"
)

type Service interface {
	CheckOTP(input OtpInput) (models.Otps, error)
	ResendOTP(ID string) (models.Otps, error)
}

type service struct {
	repository     Repository
	UserRepositroy Users.Repository
}

func NewService(repository Repository, UserRepositroy Users.Repository) *service {
	return &service{repository, UserRepositroy}
}

func (s *service) ResendOTP(ID string) (models.Otps, error) {

	//CheckExpired, err := s.repository.GetUserOtp(ID)
	GetCurrentDateTime := time.Now().Local()
	GetOTPCurrentUser, err := s.repository.GetUserOtp(ID)
	if err != nil {
		return GetOTPCurrentUser, err
	}
	t := GetOTPCurrentUser.Expired.String()

	// Prints output
	out, err := fmt.Printf("%v\n", t)
	fmt.Println(out)
	if err != nil {
		return GetOTPCurrentUser, err
	}
	if GetCurrentDateTime.Before(GetOTPCurrentUser.Expired) {

		return GetOTPCurrentUser, errors.New("current otp still valid, request again after" + strconv.Itoa(out))

	}

	//otpcode generator and local date
	rand.Seed(GetCurrentDateTime.UnixNano())
	GetRandNum := rand.Intn(10000)
	//

	GetOTPCurrentUser.Value = GetRandNum
	GetOTPCurrentUser.UsersID = ID
	GetOTPCurrentUser.Expired = GetCurrentDateTime.Add(time.Minute * 5)

	SaveOtps, err := s.repository.UpdateOTP(GetOTPCurrentUser)

	if err != nil {

		return SaveOtps, err
	}

	return SaveOtps, nil

}

func (s *service) CheckOTP(input OtpInput) (models.Otps, error) {

	fmt.Println(input.User.ID)
	GetOtp, err := s.repository.GetUserOtp(input.User.ID)
	GetCurrentDateTime := time.Now().Local()

	if err != nil {
		return GetOtp, errors.New("not authorized")
	}

	if input.Otp != GetOtp.Value {
		fmt.Println(GetOtp.Value, input.Otp)
		return GetOtp, errors.New("your otp not valid")
	}

	if GetCurrentDateTime.After(GetOtp.Expired) {

		return GetOtp, errors.New("your OTP has been expired")

	}

	fmt.Printf("ini user otp %s", GetOtp.UsersID)
	GetUser, err := s.UserRepositroy.FindUserById(GetOtp.UsersID)

	if err != nil {
		return GetOtp, err

	}

	GetUser.IsVerif = true

	_, err = s.UserRepositroy.UpdateUser(GetUser)

	if err != nil {
		return GetOtp, errors.New("failed update user")
	}
	return GetOtp, nil

}
