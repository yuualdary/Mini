package handler

import (
	"net/http"
	"pasarwarga/helper"
	"pasarwarga/location"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	LocationService location.Service
}

func NewLocationHandler(LocationService location.Service) *LocationHandler {
	return &LocationHandler{LocationService}
}

func (h *LocationHandler) CreateLocation(c *gin.Context) {

	var input location.CreateLocationInput

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

	NewLocation, err := h.LocationService.CreateLocation(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Location", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewLocation)
	c.JSON(http.StatusOK, response)

}

func (h *LocationHandler) ListLocation(c *gin.Context) {

	List, err := h.LocationService.LocationList()

	if err != nil {
		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Category", http.StatusOK, "success", List)
	c.JSON(http.StatusOK, response)
}
func (h *LocationHandler) UpdateLocation(c *gin.Context) {

	var input location.DetailLocationInput

	err := c.ShouldBindUri(input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputdata location.CreateLocationInput

	err = c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Location Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	NewLocation, err := h.LocationService.UpdateLocation(input, inputdata)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Location", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewLocation)
	c.JSON(http.StatusOK, response)

}
