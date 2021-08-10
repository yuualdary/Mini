package handler

import (
	"net/http"
	"pasarwarga/article"
	"pasarwarga/helper"

	"github.com/gin-gonic/gin"
)


type ArticleHandler struct {
	ArticleService article.Service
}

func NewArticleHandler(ArticleService article.Service)*ArticleHandler{
	return &ArticleHandler{ArticleService}
}

func(h *ArticleHandler)CreateArticle(c *gin.Context){
	var input article.CreateArticleInput

	err := c.ShouldBindJSON(&input)
	
	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage :=gin.H{
			"error" : errors,
		}

		response := helper.APIResponse("Fail Get Data From Article Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	NewArticle, err := h.ArticleService.CreateArticle(input)
	
	if err != nil{

		ErrorMessage := gin.H{
			"errors":err.Error(),
		}
		response := helper.APIResponse("Fail Save Article", http.StatusBadRequest,"errors",ErrorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewArticle)
	c.JSON(http.StatusOK,response)

}


func(h *ArticleHandler)UpdateArticle(c *gin.Context){
	
	var inputid article.ArticleDetailInput	
	err := c.ShouldBindUri(&inputid)

	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	var inputdata article.CreateArticleInput

	err = c.ShouldBindJSON(&inputdata)
	
	if err != nil {
		errors := helper.FormatValidationError(err)

		ErrorMessage :=gin.H{
			"error" : errors,
		}

		response := helper.APIResponse("Fail Get Data From Article Input", http.StatusBadRequest, "error", ErrorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	NewArticle, err := h.ArticleService.UpdateArticle(inputid,inputdata)
	
	if err != nil{

		ErrorMessage := gin.H{
			"errors":err.Error(),
		}
		response := helper.APIResponse("Fail Save Article", http.StatusBadRequest,"errors",ErrorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	response := helper.APIResponse("Detail Category Data", http.StatusOK, "success", NewArticle)
	c.JSON(http.StatusOK,response)

}
func(h *ArticleHandler)DetailArticle(c *gin.Context){

	var input article.ArticleDetailInput	
	err := c.ShouldBindUri(&input)

	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	FindDetail,err:=h.ArticleService.DetailArticle(input)

	if err != nil{

		ErrorMessage := gin.H{
			"errors":err.Error(),
		}
		response := helper.APIResponse("Fail Get Article", http.StatusBadRequest,"errors",ErrorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	response := helper.APIResponse("Detail Article Data", http.StatusOK, "success", FindDetail)
	c.JSON(http.StatusOK,response)
	


}



func(h *ArticleHandler)ListArticle(c *gin.Context){



	ListAllArticle,err := h.ArticleService.ListArticle()

	if err != nil{

		ErrorMessage := gin.H{
			"errors":err.Error(),
		}
		response := helper.APIResponse("Fail Get Article", http.StatusBadRequest,"errors",ErrorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	response := helper.APIResponse("Detail Article Data", http.StatusOK, "success",ListAllArticle)
	c.JSON(http.StatusOK,response)
	


}

func(h *ArticleHandler)DeleteArticle(c *gin.Context){

	var input article.ArticleDetailInput	
	err := c.ShouldBindUri(&input)

	if err != nil{
		response := helper.APIResponse("Fail Get Bind Data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err =h.ArticleService.DeleteArticle(input)

	if err != nil{

		ErrorMessage := gin.H{
			"errors":err.Error(),
		}
		response := helper.APIResponse("Fail Get Article", http.StatusBadRequest,"errors",ErrorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	response := helper.APIResponse("Detail Article Data", http.StatusOK, "success", "")
	c.JSON(http.StatusOK,response)
	


}