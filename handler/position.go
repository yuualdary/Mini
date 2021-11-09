package handler

import (
	"net/http"
	"pasarwarga/Company"
	"pasarwarga/Position"
	"pasarwarga/category"
	"pasarwarga/helper"
	"pasarwarga/location"
	"pasarwarga/models"

	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	PositionService Position.Service
	CategoryService category.Service
	LocationService location.Service

}

func NewPositionHandler(PositionService Position.Service,CategoryService category.Service,	LocationService location.Service) *PositionHandler {
	return &PositionHandler{PositionService,CategoryService,LocationService}
}

func (h *PositionHandler) CreatePosition(c *gin.Context) {

	var input Position.CreatePositionInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentOwner").(models.Users)//get owner
	input.Users = currentUser
	NewPosition, err := h.PositionService.CreatePosition(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Position", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Position Data", http.StatusOK, "success", NewPosition)
	c.JSON(http.StatusOK, response)
}

func (h *PositionHandler) UpdatePosition(c *gin.Context) {

	var input Position.DetailPositionInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputdata Position.CreatePositionInput

	err = c.ShouldBindJSON(&inputdata)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(models.Users)//get owner
	input.Users = currentUser
	NewPosition, err := h.PositionService.UpdatePosition(input, inputdata)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", NewPosition)
	c.JSON(http.StatusOK, response)

	
}

func (h *PositionHandler) CreatePositionTag(c *gin.Context) {

	var input Position.DetailPositionInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputdata Position.CreateTagPosition

	err = c.ShouldBindJSON(&inputdata)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(models.Users)//get owner
	input.Users = currentUser
	NewPosition, err := h.PositionService.CreateTagPosition(input, inputdata)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Position Data", http.StatusOK, "success", NewPosition)
	c.JSON(http.StatusOK, response)

}

func (h *PositionHandler) ListPosition(c *gin.Context) {


	input:= c.Query("input")
	inputjobtag := c.Query("jobtag")
	inputprovince := c.Query("province")
	inputcity := c.Query("city")
	ListPosition, err := h.PositionService.ListPosition(input,inputjobtag,inputprovince,inputcity)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Article", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Article Data", http.StatusOK, "success", Position.FormatListCandidate(ListPosition))
	c.JSON(http.StatusOK, response)

}

//buat detail




func (h *PositionHandler) ListCompanyPosition(c *gin.Context) {

	var input Company.CompanyFindIDInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FindPositionInCompany, err := h.PositionService.ListCompanyPosition(input.ID)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Position Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	// GetLocations, err := h.LocationService.LocationList()

	// if err != nil {
	// 	errors := helper.FormatValidationError(err)

	// 	ErrorMessage := gin.H{
	// 		"error": errors,
	// 	}

	// 	response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", Position.FormatCompanyListGtine(FindPositionInCompany,GetLocations))
	 response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", FindPositionInCompany)

	//better taruh di service/repo?
	c.JSON(http.StatusOK, response)

}


func (h *PositionHandler) DetailPosition(c *gin.Context) {

	var input Position.DetailPositionInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FindPosition, err := h.PositionService.DetailPosition(input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	GetCategory, err := h.CategoryService.ListCategory()

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success",Position.FormatCompanyRoutine(FindPosition,GetCategory))
	//go routine ada tinggal
	//cocok digunakan untuk DB (?)
	c.JSON(http.StatusOK, response)

}

func (h *PositionHandler) DeletePosition(c *gin.Context) {

	var input Position.DetailPositionInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(models.Users)
	input.Users = currentUser
	err = h.PositionService.DeletePosition(input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", "Delete Data")
	c.JSON(http.StatusOK, response)

}
