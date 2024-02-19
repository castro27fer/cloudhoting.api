package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ebarquero85/link-backend/src/handlers"
	"github.com/ebarquero85/link-backend/src/messages"
	translate "github.com/ebarquero85/link-backend/src/translations"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/labstack/echo/v4"
)

// change translator, get language of the headers of request
func LanguageUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		header := c.Request().Header
		translate.Change_translator(header.Get("Accept-Language"))

		return next(c)
	}
}

// Middleware para validar el token
func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		path := c.Request().URL.Path

		if strings.HasPrefix(path, "/swagger") || strings.HasPrefix(path, "/auth") {
			return next(c)
		}

		if valido := handlers.VerifyJWT(c); valido {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Token INVALIDO")

	}
}

func ErrorsLogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err != nil {

			// Open/create file
			file, err2 := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err2 != nil {
				log.Fatal("Error al abrir el archivo de registro:", err2)
			}
			defer file.Close()

			// Configure and save message in file
			log.SetOutput(file)
			log.Println(fmt.Sprintf("Error: %s | UserId: %s", err.Error(), strconv.Itoa(c.Get("UserId").(int))))

			// return next(c)

			// Check if is Http Error
			if ecttp, ok := err.(*echo.HTTPError); ok {

				if os.Getenv("ENV") == "develop" {

					return c.JSON(ecttp.Code, types.JsonResponse[string]{
						Status:  messages.ERROR,
						Message: err.Error(),
					})

				}

				// Puedes personalizar la respuesta de error según tus necesidades
				return c.JSON(http.StatusInternalServerError, types.JsonResponse[string]{
					Status:  messages.ERROR,
					Message: "Something went wrong!! :(",
				})

			}

		}

		return nil
	}
}
