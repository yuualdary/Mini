package Users

import (
	"pasarwarga/models"
)

type UserFormatter struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Initial      string `json:"initial"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	ProfilePhoto string `json:"image_url"`
}

func FormatUser(user models.Users, token string) UserFormatter {
	formatter := UserFormatter{
		ID:           user.ID,
		Name:         user.Name,
		Initial:      user.Initial,
		Email:        user.Email,
		Token:        token,
		ProfilePhoto: user.ProfilePhoto,
	}
	return formatter
}

type DetailUserFormatter struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Initial      string `json:"initial"`
	Email        string `json:"email"`
	ProfilePhoto string `json:"image_url"`
	IsVerif      bool   `json:"isverif"`
}

func DetailUserFunc(user models.Users) DetailUserFormatter {

	DetailUserFormatter := DetailUserFormatter{}
	DetailUserFormatter.ID = user.ID
	DetailUserFormatter.Name = user.Name
	DetailUserFormatter.Initial = user.Initial
	DetailUserFormatter.Email = user.Email
	DetailUserFormatter.ProfilePhoto = user.ProfilePhoto
	DetailUserFormatter.IsVerif = user.IsVerif

	return DetailUserFormatter
}
