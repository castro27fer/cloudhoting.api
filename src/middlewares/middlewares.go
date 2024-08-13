package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/handlers"
	"github.com/ebarquero85/link-backend/src/messages"
	models "github.com/ebarquero85/link-backend/src/models/auth"
	translate "github.com/ebarquero85/link-backend/src/translations"
	translation "github.com/ebarquero85/link-backend/src/translations"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/ebarquero85/link-backend/src/validators"
	"github.com/labstack/echo/v4"
)

// func Auth_validate(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		data := new(types.AuthRequest)
// 		if err := c.Bind(data); err != nil {
// 			return err // Manejar el error adecuadamente según tus necesidades
// 		}

// 		if err := c.st(data); err != nil {
// // 			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 		}

// 		// Almacenar los datos en el Context
// 		c.Set("data", data)

// 		return next(c)
// 	}

// }

func VerifyCode(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		trans := translation.Get_translator()
		message_error, _ := trans.T("code_verify_invalid")

		data := new(types.AuthRequest)

		//valid request
		if err := validators.Request(data, c); err != nil {
			// return err
			fmt.Print("Error en el middleware: ", err)

			return c.JSON(http.StatusBadRequest, err)

		}

		code := c.Request().Header.Get("Code-Verify")
		if code == "" {

			fmt.Print("code empty: ", code)

			return c.JSON(http.StatusNotAcceptable, types.JsonResponse[string]{
				Status:  strconv.Itoa(http.StatusNotAcceptable),
				Message: message_error,
			})
		}

		var codeVerify models.CodeVerifyModel
		result := db.Databases.DBPostgresql.Instance.Where("email=? AND code=? AND status=?", data.Email, code, "NotVerify").First(&codeVerify)

		fmt.Print("Code Verify: ", codeVerify, result)

		if result.Error != nil {
			return c.JSON(http.StatusNotAcceptable, types.JsonResponse[string]{
				Status:  strconv.Itoa(http.StatusNotAcceptable),
				Message: message_error,
			})
		}

		codeVerify.Status = "Verified"
		codeVerify.Update()

		// save data in context
		c.Set("AuthRequest", data)

		return next(c)
	}
}

// change translator, get language of the headers of request
func LanguageUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		header := c.Request().Header
		translate.Change_translator(header.Get("Accept-Language"))

		return next(c)
	}
}

// Middleware to validate the token
func ValidateTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		path := c.Request().URL.Path

		if strings.HasPrefix(path, "/swagger") || strings.HasPrefix(path, "/auth") {
			return next(c)
		}

		if valid := handlers.VerifyJWT(c); valid {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Token Invalid")

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
