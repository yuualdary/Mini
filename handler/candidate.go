package handler

import (
	"net/http"
	"pasarwarga/Candidate"
	"pasarwarga/helper"
	"pasarwarga/models"

	"github.com/gin-gonic/gin"
)

type CandidateHandler struct {
	CandidateService Candidate.Service
}

func NewCandidateHandler(CandidateService Candidate.Service) *CandidateHandler {
	return &CandidateHandler{CandidateService}
}
func (h *CandidateHandler) CreateCandidate(c *gin.Context) {
	var input Candidate.CreateCandidateInput

	err := c.ShouldBindJSON(input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Candidate Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(models.Users)
	input.User = currentUser

	NewCandidate, err := h.CandidateService.CreateCandidate(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Candidate", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Candidate Data", http.StatusOK, "success", NewCandidate)
	c.JSON(http.StatusOK, response)

}

func (h *CandidateHandler) ListCandidateToPosition(c *gin.Context) {

	var input Candidate.DetailCandidateInput

	err := c.ShouldBindUri(input)

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
	input.User = currentUser

	ListAllCandidate, err := h.CandidateService.ListCandidate(input)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("List Candidate", http.StatusOK, "success", ListAllCandidate)
	c.JSON(http.StatusOK, response)

}
func (h *CandidateHandler) UpdateCandidate(c *gin.Context) {

	var input Candidate.DetailCandidateInput

	err := c.ShouldBindUri(input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputdata Candidate.CreateCandidateInput

	err = c.ShouldBindJSON(input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Fail Get Data From Candidate Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(models.Users)
	inputdata.User = currentUser

	NewCandidate, err := h.CandidateService.UpdateCandidateStatus(input, inputdata)

	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Save Candidate", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Candidate Data", http.StatusOK, "success", NewCandidate)
	c.JSON(http.StatusOK, response)

}