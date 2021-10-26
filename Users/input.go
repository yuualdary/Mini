package Users

type RegisterInput struct {
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	BOD      string `json:"bod" binding:"required"`
	Password string    `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordconfirmation" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// type OtpInput struct {
// 	Otp  int `json:"otp" binding:"required"`
// 	User models.Users
// }
