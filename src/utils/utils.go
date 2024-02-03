package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CreateMd5Hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetParam(c echo.Context, name string) int {

	param := c.Param(name)

	value, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}

	return value

}
