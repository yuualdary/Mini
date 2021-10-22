package handler

import (
	"fmt"
	"net/http"
	"pasarwarga/Users"
	"pasarwarga/auth"
	"pasarwarga/helper"
	"pasarwarga/models"

	//"github.com/go-redis/redis/v7"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService Users.Service
	// AuthService auth.TokenInterface
	// Servers     auth.AuthInterface
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

	response := helper.APIResponse("Account Has Been Registered, You Can Check Your Email For Verification", http.StatusOK, "success", NewUser)
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

	path := fmt.Sprintf("images/%s-%s", UserID, file.Filename)
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

// func (h *UserHandler) LoginUser(c *gin.Context) {
// 	var input Users.LoginInput

// 	err := c.ShouldBindJSON(&input)

// 	if err != nil {
// 		errors := helper.FormatValidationError(err)

// 		ErrorMessage := gin.H{
// 			"error": errors,
// 		}
// 		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return

// 	}

// 	NewLogin, err := h.UserService.LoginUser(input)

// 	if err != nil {
// 		ErrorMessage := gin.H{
// 			"error": err.Error(),
// 		}
// 		response := helper.APIResponse("Login Failed", http.StatusBadRequest, "error", ErrorMessage)
// 		c.JSON(http.StatusUnprocessableEntity, response)
// 		return
// 	}

// 	ts, err := h.AuthService.GenerateToken(NewLogin.ID)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}
// 	fmt.Println(NewLogin)
// 	//save token to redis
// 	saveErr := h.Servers.CreateAuth(NewLogin.ID, ts)
// 	if saveErr != nil {
// 		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
// 	}
// 	tokens := map[string]string{
// 		"access_token":  ts.AccessToken,
// 		"refresh_token": ts.RefreshToken,
// 	}
// 	c.JSON(http.StatusOK, tokens)
// }

// func (h *UserHandler) Logout(c *gin.Context) {
// 	//If metadata is passed and the tokens valid, delete them from the redis store
// 	metadata, _ := h.AuthService.ExtractTokenMetadata(c.Request)
// 	if metadata != nil {
// 		deleteErr := h.Servers.DeleteTokens(metadata)
// 		if deleteErr != nil {
// 			c.JSON(http.StatusBadRequest, deleteErr.Error())
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusOK, "Successfully logged out")
// }
