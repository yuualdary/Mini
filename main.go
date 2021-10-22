package main

import (
	"pasarwarga/Candidate"
	"pasarwarga/Company"
	"pasarwarga/File"
	"pasarwarga/Otp"
	"pasarwarga/Position"
	"pasarwarga/Users"
	"pasarwarga/article"
	"pasarwarga/auth"
	"pasarwarga/category"
	"pasarwarga/config"
	"pasarwarga/handler"
	"pasarwarga/location"
	"pasarwarga/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())
	config.ConnectDatabase()
	//	var redis *redis.Client
	CategoryRepository := category.NewRepository(config.DB)
	ArticleRepository := article.NewRepository(config.DB)
	UserRepository := Users.NewRepository(config.DB)
	CompanyRepository := Company.NewRepository(config.DB)
	OtpRepository := Otp.NewRepository(config.DB)
	LocationRepository := location.NewRepository(config.DB)
	PositionRepository := Position.NewRepository(config.DB)
	CandidateRepository := Candidate.NewRepository(config.DB)
	FilePdfRepository := File.NewRepository(config.DB)

	CategoryService := category.NewService(CategoryRepository, ArticleRepository)
	ArticleService := article.NewService(ArticleRepository)
	UsersService := Users.NewService(UserRepository)
	OtpService := Otp.NewService(OtpRepository, UserRepository)
	CompanyService := Company.NewService(CompanyRepository, UserRepository)
	FilePdfService := File.NewService(FilePdfRepository,UserRepository )
	LocationService := location.NewService(LocationRepository)

	PositionService := Position.NewService(PositionRepository, CompanyRepository)
	CandidateService := Candidate.NewService(CandidateRepository, UserRepository, PositionRepository, CompanyRepository)
	AuthService := auth.NewService()
	//AuthToken := auth.NewTokenService()
	//Servers := auth.NewAuthService(redis)

	CategoryHandler := handler.NewCategoryHandler(CategoryService)
	ArticleHandler := handler.NewArticleHandler(ArticleService)
	UsersHandler := handler.NewUserHandler(UsersService, AuthService)
	CompanyHandler := handler.NewCompanyHandler(CompanyService)
	OtpHandler := handler.NewOtpHandler(OtpService)
	LocationHandler := handler.NewLocationHandler(LocationService)
	PositionHandler := handler.NewPositionHandler(PositionService,CategoryService,LocationService)
	CandidateHandler := handler.NewCandidateHandler(CandidateService,CompanyService)
	FilePdfHandler := handler.NewFilePdfHandler(FilePdfService)


	router.Static("/images", "./images")

	v1 := router.Group("/api/v1")
	{
		v1.POST("/category/", CategoryHandler.CreateCategory)
		v1.GET("/category/:id", CategoryHandler.DetailCategory)
		v1.GET("/category/", CategoryHandler.ListCategory)
		v1.GET("/categorypositiontag/", CategoryHandler.ListPositionTag)
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


		v1.POST("/otp/otpcheck", middleware.AuthMiddlewareUnregistered(AuthService, UsersService), OtpHandler.CheckOtp)
		v1.POST("/otp/resendotp", middleware.AuthMiddlewareUnregistered(AuthService, UsersService), OtpHandler.ResendOTP)
		
		v1.POST("/company", middleware.AuthMiddleware(AuthService, UsersService), CompanyHandler.CreateCompany)
		v1.POST("/company/:id", CompanyHandler.DetailCompany)
		v1.PUT("/company/:id", middleware.AuthMiddleware(AuthService, UsersService), CompanyHandler.UpdateCompany)
		v1.GET("/company", CompanyHandler.ListCompany)
		v1.GET("/company/position/:id", PositionHandler.ListCompanyPosition)

		v1.POST("/location", LocationHandler.CreateLocation)
		v1.GET("/location", LocationHandler.ListLocation)

		v1.GET("/position", PositionHandler.ListPosition)
		v1.POST("/position/:id", PositionHandler.DetailPosition)
		v1.POST("/position", middleware.AuthMiddleware(AuthService, UsersService), PositionHandler.CreatePosition)
		v1.POST("/positioncategory/:id", middleware.AuthMiddleware(AuthService, UsersService),PositionHandler.CreatePositionTag)
		v1.GET("/candidate/:id", middleware.AuthCompanyMiddleware(AuthService, UsersService, CompanyService), CandidateHandler.ListCandidateToPosition)

		v1.POST("/candidate/", middleware.AuthMiddleware(AuthService, UsersService), CandidateHandler.CreateCandidate)
		v1.GET("/users/application", middleware.AuthMiddleware(AuthService, UsersService), CandidateHandler.ListUserApplication)
		v1.PUT("/candidate/:id", middleware.AuthCompanyMiddleware(AuthService, UsersService, CompanyService), CandidateHandler.UpdateCandidate)

		v1.POST("/filepdf", middleware.AuthMiddleware(AuthService, UsersService), FilePdfHandler.CreateFilePDF)
		v1.GET("/filepdf/:id", middleware.AuthMiddleware(AuthService, UsersService), FilePdfHandler.DetailFile)




	}

	router.Run(":8000")

}
