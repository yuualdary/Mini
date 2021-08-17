package main

import (
	"pasarwarga/Otp"
	"pasarwarga/Users"
	"pasarwarga/article"
	"pasarwarga/auth"
	"pasarwarga/category"
	"pasarwarga/company"
	"pasarwarga/config"
	"pasarwarga/handler"
	"pasarwarga/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	config.ConnectDatabase()

	CategoryRepository := category.NewRepository(config.DB)
	ArticleRepository := article.NewRepository(config.DB)
	UserRepository := Users.NewRepository(config.DB)
	CompanyRepository := company.NewRepository(config.DB)
	OtpRepository := Otp.NewRepository(config.DB)

	CategoryService := category.NewService(CategoryRepository, ArticleRepository)
	ArticleService := article.NewService(ArticleRepository)
	UsersService := Users.NewService(UserRepository)
	OtpService := Otp.NewService(OtpRepository, UserRepository)
	CompanyService := company.NewService(CompanyRepository, UserRepository)
	AuthService := auth.NewService()

	CategoryHandler := handler.NewCategoryHandler(CategoryService)
	ArticleHandler := handler.NewArticleHandler(ArticleService)
	UsersHandler := handler.NewUserHandler(UsersService, AuthService)
	CompanyHandler := handler.NewCompanyHandler(CompanyService)
	OtpHandler := handler.NewOtpHandler(OtpService)

	router.Static("/images", "./images")

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
		v1.POST("/users/register", UsersHandler.RegisterUser)
		v1.POST("/users/login", UsersHandler.LoginUser)
		v1.POST("/users/uploadavatar", middleware.AuthMiddleware(AuthService, UsersService), UsersHandler.SaveAvatar)
		v1.POST("/otp/otpcheck", middleware.AuthMiddleware(AuthService, UsersService), OtpHandler.CheckOtp)
		v1.POST("/otp/resendotp", middleware.AuthMiddleware(AuthService, UsersService), OtpHandler.ResendOTP)
		v1.POST("/company", middleware.AuthMiddleware(AuthService, UsersService), CompanyHandler.CreateCompany)
		v1.POST("/company/:id", CompanyHandler.DetailCompany)
		v1.PUT("/company/:id", middleware.AuthMiddleware(AuthService, UsersService), CompanyHandler.UpdateCompany)
		v1.GET("/company/:id", CompanyHandler.DetailCompany)
		v1.GET("/company", CompanyHandler.ListCompany)

	}

	router.Run(":8000")

}
