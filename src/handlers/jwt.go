package handlers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ebarquero85/link-backend/src/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTClaim struct {
	UserId int `json:"UserId"`
	jwt.RegisteredClaims
}

func GenerateJWT(User *models.UserModel) string {

	TOKEN_EXPIRATION_HOURS, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HOURS"))
	JWT_KEY := os.Getenv("JWT_KEY")

	duracion := time.Duration(TOKEN_EXPIRATION_HOURS) * time.Hour

	expirationTime := time.Now().Add(duracion)

	claims := &JWTClaim{
		UserId: User.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := []byte(JWT_KEY)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return ""
	}

	return tokenString

}

func VerifyJWT(c echo.Context) bool {

	JWT_KEY := os.Getenv("JWT_KEY")
	reqToken := GetTokenFromHeader(c)

	if reqToken == "" {
		return false
	}

	token, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_KEY), nil
	})

	if payload, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// Check if token expired
		if float64(time.Now().Unix()) > payload["exp"].(float64) {
			return false
		}

		UserId := int(payload["UserId"].(float64))

		c.Set("UserId", UserId)

		return true // Valid

	} else {

		return false // Invalid
	}

}

func GetTokenFromHeader(c echo.Context) string {

	reqToken := c.Request().Header.Get("Authorization")

	if reqToken == "" {
		return ""
	}

	splitToken := strings.Split(reqToken, "Bearer ")

	if len(splitToken) < 2 {
		return ""
	}

	reqToken = splitToken[1]

	return reqToken

}
