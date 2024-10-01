package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	UserId int `json:"UserId"`
	jwt.RegisteredClaims
}

type JWT struct {
	ExpirationTime int
	SecretKey      string
}

func (jwtService *JWT) CreateToken(User *User) string {

	time_expired := time.Duration(jwtService.ExpirationTime) * time.Hour

	expirationTime := time.Now().Add(time_expired)

	claims := &JWTClaim{
		UserId: int(User.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtKey := []byte(jwtService.SecretKey)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return ""
	}

	return tokenString

}

func (JWTService *JWT) VerifyToken(token string) (bool, error) {

	tokenSuccess, error_parse := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(JWTService.SecretKey), nil
	})

	if error_parse != nil {
		return false, error_parse
	}

	var payload jwt.MapClaims
	var ok bool

	if payload, ok = tokenSuccess.Claims.(jwt.MapClaims); !ok || !tokenSuccess.Valid {
		return false, nil // Invalid
	}

	// Check if token expired
	if float64(time.Now().Unix()) > payload["exp"].(float64) {
		return false, nil
	}

	return true, nil
}

type JWTS struct {
	ActiveAccount JWT
	ResetPassword JWT
	Session       JWT
}

var JWTService JWTS = JWTS{}

func (jwtService *JWTS) Start() error {

	var errToken error = nil

	//Token of Session
	if jwtService.Session.SecretKey = os.Getenv("JWT_SESSION_KEY"); jwtService.Session.SecretKey != "" {
		return errors.New("JWT_SESSION_KEY not found")
	}

	if jwtService.Session.ExpirationTime, errToken = strconv.Atoi(os.Getenv("JWT_SESSION_EXPIRATION")); errToken != nil {
		return errToken
	}

	//Token of reset password
	if jwtService.ResetPassword.SecretKey = os.Getenv("JWT_RESET_PASSWORD_KEY"); jwtService.ResetPassword.SecretKey != "" {
		return errors.New("JWT_RESET_PASSWORD_KEY not found")
	}

	if jwtService.ResetPassword.ExpirationTime, errToken = strconv.Atoi(os.Getenv("JWT_SESSION_EXPIRATION")); errToken != nil {
		return errToken
	}

	//Token of activation account
	if jwtService.ActiveAccount.SecretKey = os.Getenv("JWT_RESET_PASSWORD_KEY"); jwtService.ActiveAccount.SecretKey != "" {
		return errors.New("JWT_RESET_PASSWORD_KEY not found")
	}

	if jwtService.ActiveAccount.ExpirationTime, errToken = strconv.Atoi(os.Getenv("JWT_SESSION_EXPIRATION")); errToken != nil {
		return errToken
	}

	return nil
}

// func GetTokenFromHeader(c echo.Context) string {

// 	reqToken := c.Request().Header.Get("Authorization")

// 	if reqToken == "" {
// 		return ""
// 	}

// 	splitToken := strings.Split(reqToken, "Bearer ")

// 	if len(splitToken) < 2 {
// 		return ""
// 	}

// 	reqToken = splitToken[1]

// 	return reqToken

// }
