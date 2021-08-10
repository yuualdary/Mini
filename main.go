package main

import (
	"pasarwarga/article"
	"pasarwarga/category"
	"pasarwarga/config"
	"pasarwarga/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	config.ConnectDatabase()

	CategoryRepository := category.NewRepository(config.DB)
	ArticleRepository := article.NewRepository(config.DB)

	CategoryService := category.NewService(CategoryRepository, ArticleRepository)
	ArticleService := article.NewService(ArticleRepository)

	CategoryHandler := handler.NewCategoryHandler(CategoryService)
	ArticleHandler := handler.NewArticleHandler(ArticleService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/category/", CategoryHandler.CreateCategory)
		v1.GET("/category/:id", CategoryHandler.DetailCategory)
		v1.GET("/category/", CategoryHandler.ListCategory)
		v1.PUT("/category/:id", CategoryHandler.UpdateCategory)
		v1.DELETE("/category/:id", CategoryHandler.DeleteCategory)
		v1.POST("/article", ArticleHandler.CreateArticle)
		v1.GET("/article/:id", ArticleHandler.DetailArticle)
		v1.GET("/article", ArticleHandler.ListArticle)
		v1.PUT("/article/:id", ArticleHandler.UpdateArticle)
		v1.DELETE("/article/:id", ArticleHandler.DeleteArticle)

	}

	router.Run(":8000")

}
