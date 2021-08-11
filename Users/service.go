package Users

import (
	"errors"
	"fmt"
	"math/rand"
	"pasarwarga/generatornumber"
	"pasarwarga/models"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (models.Users, error)
	SaveAvatar(ID string, filelocatiion string) (models.Users, error)
	LoginUser(input LoginInput) (models.Users, error)
	GetUserById(ID string) (models.Users, error)
	ResendOTP(ID string) (models.Otps, error)
	CreateOtp(UserID string) (models.Otps, error)
	CheckOTP(input OtpInput) (models.Otps, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) CreateOtp(UserID string) (models.Otps, error) {

	Otp := models.Otps{}
	Otp.ID = generatornumber.NewUUID()
	Otp.UsersID = UserID
	Otp.Value = rand.Intn(10000)
	Otp.Expired = time.Now().Local().Add(time.Minute * 5)
	Save, err := s.repository.SaveOTP(Otp)

	if err != nil {
		return Save, err
	}
	return Save, nil
}

func (s *service) RegisterUser(input RegisterInput) (models.Users, error) {

	User := models.Users{}
	User.ID = generatornumber.NewUUID()
	User.Name = input.Name
	User.Email = input.Email
	User.Bod = input.BOD

	GetSplit := strings.Replace(User.Name, " ", "", -1)
	GetChar := GetSplit[0:3]
	rand.Seed(time.Now().Local().UnixNano())
	GetRandNum := rand.Intn(10000)
	User.Initial = GetChar + strconv.Itoa(GetRandNum)

	GenereateHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {

		return User, err
	}

	User.Password = string(GenereateHash)
	User.Role = "candidate"

	CheckMail, err := s.repository.FindUserEmail(User.Email)

	if err != nil {
		return CheckMail, err

	}

	if CheckMail.ID != "" {

		return CheckMail, errors.New("Email Already Used")
	}

	NewUser, err := s.repository.SaveUser(User)

	if err != nil {

		return NewUser, err
	}

	_, err = s.CreateOtp(NewUser.ID)

	if err != nil {
		return NewUser, err
	}

	return NewUser, nil

}

func (s *service) LoginUser(input LoginInput) (models.Users, error) {

	email := input.Email
	password := input.Password

	//get data user

	GetDataUser, err := s.repository.FindUserEmail(email)

	if err != nil {

		return GetDataUser, err
	}

	if GetDataUser.ID == "" {

		return GetDataUser, errors.New("Email Not Found")
	}

	//hash password

	err = bcrypt.CompareHashAndPassword([]byte(GetDataUser.Password), []byte(password))

	if err != nil {

		return GetDataUser, err

	}

	return GetDataUser, nil
}
func (s *service) SaveAvatar(ID string, filelocatiion string) (models.Users, error) {

	//check id

	CheckUserID, err := s.repository.FindUserById(ID)

	if err != nil {

		return CheckUserID, err
	}

	CheckUserID.ProfilePhoto = filelocatiion

	Update, err := s.repository.SaveUser(CheckUserID)

	if err != nil {
		return Update, err
	}

	return Update, nil

}
func (s *service) GetUserById(ID string) (models.Users, error) {

	user, err := s.repository.FindUserById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("No user found")
	}
	return user, nil
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

// func (s *service) CheckOTP(input OtpInput) (models.Otps, error) {

// 	fmt.Println(input.User.ID)
// 	GetOtp, err := s.repository.GetUserOtp(int(input.User.ID))
// 	GetCurrentDateTime := time.Now().Local()

// 	if err != nil {
// 		return GetOtp, errors.New("not authorized")
// 	}

// 	if input.Otp != GetOtp.Value {
// 		fmt.Println(GetOtp.Value, input.Otp)
// 		return GetOtp, errors.New("your otp not valid")
// 	}

// 	if GetCurrentDateTime.After(GetOtp.Expired) {

// 		return GetOtp, errors.New("your OTP has been expired")

// 	}

// 	fmt.Printf("ini user otp %d", int(GetOtp.UsersID))
// 	GetUser, err := s.repository.FindUserById(int(GetOtp.UsersID))

// 	if err != nil {
// 		return GetOtp, err

// 	}

// 	GetUser.IsVerif = true

// 	_, err = s.repository.UpdateUser(GetUser)

// 	if err != nil {
// 		return GetOtp, errors.New("failed update user")
// 	}
// 	return GetOtp, nil

// }

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

	fmt.Printf("ini user otp %d", GetOtp.UsersID)
	GetUser, err := s.repository.FindUserById(GetOtp.UsersID)

	if err != nil {
		return GetOtp, err

	}

	GetUser.IsVerif = true

	_, err = s.repository.UpdateUser(GetUser)

	if err != nil {
		return GetOtp, errors.New("failed update user")
	}
	return GetOtp, nil

}
