package handler

import (
	"fmt"
	"net/http"
	"pasarwarga/File"
	"pasarwarga/helper"
	"pasarwarga/models"

	"github.com/gin-gonic/gin"
)


type FilePdfHandler struct {
	PdfService File.Service
}

func NewFilePdfHandler(PdfService File.Service) *FilePdfHandler{
	return &FilePdfHandler{PdfService}
}

func (h *FilePdfHandler) CreateFilePDF(c *gin.Context){

	
	currentUser := c.MustGet("CurrentUser").(models.Users)
	userID := currentUser.ID
	
	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	file.Filename = userID+".pdf"//change name

	path := fmt.Sprintf("userpdf/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	NewFile, err := h.PdfService.CreateService(userID,path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewFile)
	c.JSON(http.StatusOK, response)
}



func (h *FilePdfHandler) UpdateFile(c *gin.Context){

	var inputid File.DetailFile

	err := c.ShouldBindUri(&inputid)

	if err != nil {
		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(models.Users)
	userID := currentUser.ID
	
	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	file.Filename = userID+".pdf"//change name
	//todo
	//kalau bisa pikirin kalau ditengah" gagal

	path := fmt.Sprintf("userpdf/%s", file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	NewFile, err := h.PdfService.UpdateFile(inputid,path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewFile)
	c.JSON(http.StatusOK, response)
}


func (h *FilePdfHandler) DetailFile(c *gin.Context){

	var inputid File.DetailFile

	err := c.ShouldBindUri(&inputid)

	if err != nil {
		ErrorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Fail Get Data", http.StatusBadRequest, "errors", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(models.Users)
	inputid.User = currentUser

	NewFile, err := h.PdfService.DetailFile(inputid)
	if err != nil {

		ErrorMessage := gin.H{
			"errors": err.Error(),
		}

		response := helper.APIResponse("Fail Get Data From Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Detail User Data", http.StatusOK, "success", NewFile)
	c.JSON(http.StatusOK, response)
}