package middleware

import (
	"net/http"
	"pasarwarga/Company"
	"pasarwarga/Users"
	"pasarwarga/auth"
	"pasarwarga/helper"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(AuthService auth.Service, UserService Users.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		AuthHeader := c.GetHeader("Authorization")

		if !strings.Contains(AuthHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//bearer tokentokentoken split bearer dengan spasi

		TokenString := ""

		ArrayToken := strings.Split(AuthHeader, " ")

		if len(ArrayToken) == 2 {
			TokenString = ArrayToken[1]
			//ambil array token
		}
		//check apakah token valid
		Token, err := auth.NewTokenService().ValidateToken(TokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//ngeclaim token
		Claim, IsTokenOk := Token.Claims.(jwt.MapClaims)

		if !IsTokenOk || !Token.Valid {

			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		UserID := (Claim["user_id"].(string)) //map -> string

		User, err := UserService.GetUserById(UserID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("CurrentUser", User)

	}

}

// func AuthCompanyMiddleware(CompanyService Company.Service) gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		session := sessions.Default(c)

// 		UserIDSession := session.Get("UserID")

// 		if UserIDSession == nil {
// 			c.Redirect(http.StatusNotFound, "/login")
// 			return
// 		}
// 	}
// }

func AuthCompanyMiddleware(AuthService auth.Service, UserService Users.Service, CompanyService Company.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		AuthHeader := c.GetHeader("Authorization")

		if !strings.Contains(AuthHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//bearer tokentokentoken split bearer dengan spasi

		TokenString := ""

		ArrayToken := strings.Split(AuthHeader, " ")

		if len(ArrayToken) == 2 {
			TokenString = ArrayToken[1]
			//ambil array token
		}
		//check apakah token valid
		Token, err := auth.NewTokenService().ValidateToken(TokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//ngeclaim token
		Claim, IsTokenOk := Token.Claims.(jwt.MapClaims)

		if !IsTokenOk || !Token.Valid {

			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		UserID := (Claim["user_id"].(string)) //map -> string

		GetOwner, err := CompanyService.CompanyOwner(UserID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//User = GetOwner
		c.Set("CurrentOwner", GetOwner)
	}
}


func AuthMiddlewareUnregistered(AuthService auth.Service, UserService Users.Service) gin.HandlerFunc {

	return func(c *gin.Context) {

		AuthHeader := c.GetHeader("Authorization")

		if !strings.Contains(AuthHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		//bearer tokentokentoken split bearer dengan spasi

		TokenString := ""

		ArrayToken := strings.Split(AuthHeader, " ")

		if len(ArrayToken) == 2 {
			TokenString = ArrayToken[1]
			//ambil array token
		}
		//check apakah token valid
		Token, err := auth.NewTokenService().ValidateToken(TokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		//ngeclaim token
		Claim, IsTokenOk := Token.Claims.(jwt.MapClaims)

		if !IsTokenOk || !Token.Valid {

			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return

		}

		UserID := (Claim["user_id"].(string)) //map -> string

		User, err := UserService.GetUserUnregisteredById(UserID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("CurrentUser", User)

	}

}