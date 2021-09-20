package handler

import (
	"net/http"
	"pasarwarga/category"
	"pasarwarga/helper"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService category.Service
}

func NewCategoryHandler(CategoryService category.Service) *CategoryHandler {
	return &CategoryHandler{CategoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {

	var input category.CategoryInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Category Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	NewCategory, err := h.CategoryService.CreateCategory(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Category", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewCategory)
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) DetailCategory(c *gin.Context) {

	var input category.CategoryIDInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FindDetail, err := h.CategoryService.DetailCategory(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Category", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", FindDetail)
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {

	var input category.CategoryIDInput
	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.CategoryService.DeleteCategory(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Category", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete Category Data", http.StatusOK, "success", "")
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) ListCategory(c *gin.Context) {

	ListAllCategory, err := h.CategoryService.ListCategory()

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Category", http.StatusOK, "success", ListAllCategory)
	c.JSON(http.StatusOK, response)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {

	var inputid category.CategoryIDInput
	err := c.ShouldBindUri(&inputid)

	if err != nil {
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputdata category.CategoryInput
	err = c.ShouldBindJSON(&inputdata)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Category Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	NewCategory, err := h.CategoryService.UpdateCategory(inputid, inputdata)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Update Category", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewCategory)
	c.JSON(http.StatusOK, response)
}
