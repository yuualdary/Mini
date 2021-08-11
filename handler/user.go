package handler

import (
	"fmt"
	"net/http"
	"pasarwarga/Users"
	"pasarwarga/auth"
	"pasarwarga/helper"
	"pasarwarga/models"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService Users.Service
	AuthService auth.Service
}

func NewUserHandler(UserService Users.Service, AuthService auth.Service) *UserHandler {
	return &UserHandler{UserService, AuthService}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {

	var input Users.RegisterInput

	err := c.ShouldBindJSON(&input)
	//form validation

	// DateValidator:=helper.DateValidator(input.BOD)

	// if DateValidator != ""{

	// 	response := helper.APIResponse("Fail Register Data", http.StatusBadRequest,"errors",DateValidator)
	// 	c.JSON(http.StatusUnprocessableEntity,response)
	// 	return
	// }

	if err != nil {

		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Fail Register Data", http.StatusBadRequest, "errors", ErrorMessage)
		// fmt.Println(response)

		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//

	NewUser, err := h.UserService.RegisterUser(input)

	if err != nil {
		// fmt.Println(err)

		ErrorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Fail Register Data", http.StatusBadRequest, "errors", ErrorMessage)
		fmt.Println(response)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Account Has Been Registered", http.StatusOK, "success", NewUser)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) SaveAvatar(c *gin.Context) {

	file, err := c.FormFile("file")

	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload an image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	CurrentUser := c.MustGet("CurrentUser").(models.Users)
	UserID := CurrentUser.ID

	path := fmt.Sprintf("images/%d-%s", UserID, file.Filename)
	err = c.SaveUploadedFile(file, path)

	if err != nil {

		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload an image", http.StatusBadRequest, "error", data)
		fmt.Println(response)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.UserService.SaveAvatar(UserID, path)

	if err != nil {

		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("Failed to upload an image", http.StatusBadRequest, "error", data)
		fmt.Println(response)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar Successfully Uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	var input Users.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	NewLogin, err := h.UserService.LoginUser(input)

	if err != nil {
		ErrorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.AuthService.GenerateToken(NewLogin.ID)

	if err != nil {
		ErrorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := Users.FormatUser(NewLogin, token)
	response := helper.APIResponse("success Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *UserHandler) ResendOTP(c *gin.Context) {

	CurrentUser := c.MustGet("CurrentUser").(models.Users)
	UserID := CurrentUser.ID
	fmt.Println(UserID)
	NewOtp, err := h.UserService.ResendOTP(UserID)

	if err != nil {
		ErrorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("success Request Otp", http.StatusOK, "success", NewOtp)
	c.JSON(http.StatusOK, response)

}

func (h *UserHandler) CheckOtp(c *gin.Context) {

	var input Users.OtpInput
	err := c.ShouldBindJSON(&input)
	//err adalah value dari input ->mendapatkan semua isi dan yang tidak diisi, jika ada tidak diisi masuk ke validasi (err)

	if err != nil {

		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}
		response := helper.APIResponse("Fail Add Data", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	CurrentUser := c.MustGet("CurrentUser").(models.Users)
	input.User.ID = CurrentUser.ID
	//fmt.Println(UserID)
	NewOtp, err := h.UserService.CheckOTP(input)

	if err != nil {
		ErrorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	GetDetail, err := h.UserService.GetUserById(NewOtp.UsersID)

	if err != nil {
		ErrorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := Users.DetailUserFunc(GetDetail)
	response := helper.APIResponse("success Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
