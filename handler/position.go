package handler

import (
	"net/http"
	"pasarwarga/helper"
	"pasarwarga/models"
	Position "pasarwarga/position"

	"github.com/gin-gonic/gin"
)

type PositionHandler struct {
	PositionService Position.Service
}

func NewPositionHandler(PositionService Position.Service) *PositionHandler {
	return &PositionHandler{PositionService}
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

	currentUser := c.MustGet("CurrentUser").(models.Users)
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

func (h *PositionHandler) ListPosition(c *gin.Context) {

	ListPosition, err := h.PositionService.ListPosition()

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Article", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Article Data", http.StatusOK, "success", ListPosition)
	c.JSON(http.StatusOK, response)

}

//buat detail

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
	response := helper.APIResponse("Detail Company Data", http.StatusOK, "success", Position.FormatDetailPosition(FindPosition))
	c.JSON(http.StatusOK, response)

}
