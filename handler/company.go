package handler

import (
	"fmt"
	"net/http"
	"pasarwarga/Company"
	"pasarwarga/helper"
	"pasarwarga/models"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	CompanyService Company.Service
}

func NewCompanyHandler(CompanyService Company.Service) *CompanyHandler {
	return &CompanyHandler{CompanyService}
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var input Company.CreateCompanyInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Company Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(models.Users)
	input.User = currentUser
	fmt.Println(input)

	NewCategory, err := h.CompanyService.CreateCompany(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Company", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", NewCategory)
	c.JSON(http.StatusOK, response)
}
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {

	var inputid Company.CompanyFindIDInput
	err := c.ShouldBindUri(&inputid)

	if err != nil {
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input Company.CreateCompanyInput

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Company Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println("test")
	currentUser := c.MustGet("CurrentUser").(models.Users)
	input.User = currentUser
	NewCategory, err := h.CompanyService.UpdateCompany(input, inputid)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Company", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", NewCategory)
	c.JSON(http.StatusOK, response)

}

func (h *CompanyHandler) DetailCompany(c *gin.Context) {

	var input Company.CompanyFindIDInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FindDetail, err := h.CompanyService.DetailCompany(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Company Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//fmt.Println(lokasi)
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", Company.FormatCompany(FindDetail))
	c.JSON(http.StatusOK, response)
}
func (h *CompanyHandler) ListCompany(c *gin.Context) {

	input:= c.Query("input")
	inputjobtag := c.Query("jobtag")
	inputprovince := c.Query("province")
	inputcity := c.Query("city")
	// input := err
//	fmt.Println(input)
	ListAllCompany, err := h.CompanyService.ListCompany(input,inputjobtag,inputprovince,inputcity)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	response := helper.APIResponse("List Company", http.StatusOK, "success", Company.FormatListCompany(ListAllCompany))
	c.JSON(http.StatusOK, response)

}
