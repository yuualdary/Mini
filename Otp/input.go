package Otp

import "pasarwarga/models"

type OtpInput struct {
	Otp  int `json:"otp" binding:"required"`
	User models.Users
}
