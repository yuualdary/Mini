package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type TokenManager struct{}

func NewTokenService() *TokenManager {
	return &TokenManager{}
}

type TokenInterface interface {
	GenerateToken(userID string) (*TokenDetails, error)
	ValidateToken(EncodeToken string) (*jwt.Token, error)
	ExtractTokenMetadata(*http.Request) (*AccessDetails, error)
}

//Token implements the TokenInterface
var _ TokenInterface = &TokenManager{}

func (s *TokenManager) GenerateToken(userID string) (*TokenDetails, error) {
	//masukkin ke token detail
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix() //expires after 30 min
	td.TokenUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userID
	var err error

	claim := jwt.MapClaims{}
	claim["access_uuid"] = td.AccessToken
	claim["user_id"] = userID //value dari user
	claim["exp"] = td.AtExpires
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	//token valid jika, dibuat dengan secret key
	td.AccessToken, err = token.SignedString(SECRET_KEY)

	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userID

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	td.RefreshToken, err = rtoken.SignedString(SECRET_KEY)
	if err != nil {
		return nil, err
	}
	return td, nil

	//	return SignedToken, nil
}

func (s *TokenManager) ValidateToken(EncodeToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(EncodeToken, func(token *jwt.Token) (interface{}, error) { //func ny bawaan
		//jadi fungsi func mengecek apakah token yang dibuat sesuai dengan secret_key yand kita buat
		_, ok := token.Method.(*jwt.SigningMethodHMAC) //tipenya HMAC karena diatas pake HS256

		if !ok {

			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil

}

func (t *TokenManager) ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	acc, err := Extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

//get the token from the request body
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Extract(token *jwt.Token) (*AccessDetails, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userId, userOk := claims["user_id"].(string)
		userName, userNameOk := claims["user_name"].(string)
		if ok == false || userOk == false || userNameOk == false {
			return nil, errors.New("unauthorized")
		} else {
			return &AccessDetails{
				TokenUuid: accessUuid,
				UserId:    userId,
				UserName:  userName,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	acc, err := Extract(token)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
