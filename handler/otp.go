package handler

import (
	"fmt"
	"net/http"
	"pasarwarga/Otp"
	"pasarwarga/helper"
	"pasarwarga/models"

	"github.com/gin-gonic/gin"
)

type OtpHandler struct {
	service Otp.Service
}

func NewOtpHandler(service Otp.Service) *OtpHandler {
	return &OtpHandler{service}
}
func (h *OtpHandler) ResendOTP(c *gin.Context) {

	CurrentUser := c.MustGet("CurrentUser").(models.Users)
	UserID := CurrentUser.ID
	fmt.Println(UserID)
	NewOtp, err := h.service.ResendOTP(UserID)

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
func (h *OtpHandler) CheckOtp(c *gin.Context) {

	var input Otp.OtpInput
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
	NewOtp, err := h.service.CheckOTP(input)

	if err != nil {
		ErrorMessage := gin.H{
			"error": err.Error(),
		}
		response := helper.APIResponse("Failed", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// formatter := Users.DetailUserFunc(GetDetail)
	response := helper.APIResponse("success Login", http.StatusOK, "success", NewOtp)
	c.JSON(http.StatusOK, response)

}
