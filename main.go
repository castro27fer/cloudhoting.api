package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/ebarquero85/link-backend/docs/links" // Necesario para que funcione Swagger
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"

	"github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/handlers"
	"github.com/ebarquero85/link-backend/src/middlewares"
	en_translations "github.com/ebarquero85/link-backend/src/translations/en"
	"github.com/ebarquero85/link-backend/src/types"
)

type CustomValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

// use a single instance , it caches struct info
var (
	uni *ut.UniversalTranslator
	// validate *validator.Validate
)

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {

		var errors2 []types.Error_Request

		for _, err := range err.(validator.ValidationErrors) {

			errors2 = append(errors2, types.Error_Request{Name: err.Field(), Message: err.Translate(cv.trans)})

			fmt.Println("namespace", err.Namespace())
			fmt.Println("field", err.Field())
			fmt.Println("structNamespace", err.StructNamespace())
			fmt.Println("structField", err.StructField())
			fmt.Println("Tag", err.Tag())
			fmt.Println("ActualTag", err.ActualTag())
			fmt.Println("Kind", err.Kind())
			fmt.Println("type", err.Type())
			fmt.Println("value", err.Value())
			fmt.Println("param", err.Param())
			fmt.Println("error", err.Translate(cv.trans))
		}

		// Convertir la instancia a una cadena JSON
		// jsonString, err := json.Marshal(errors2)
		// if err != nil {
		// 	fmt.Println("Error al convertir a JSON:", err)
		// 	return err
		// }

		//return errors.New(string(jsonString))
		return echo.NewHTTPError(http.StatusBadRequest, errors2)
	}

	return nil
}

func init() {

	// Load .env configuration
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	database.Connect("postgres")

	// database.Databases.DBPostgresql.Instance.Migrator().AutoMigrate(&models.UserModel{}, &models.AccountModel{})
}

// @title Links App API
// @version 1.0

// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	// Echo instance
	e := echo.New()

	en := en.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	// Se agrega el paquete Validator a Echo
	v := validator.New()
	e.Validator = &CustomValidator{validator: v, trans: trans} //descomentariar para activar las validaciones

	en_translations.RegisterDefaultTranslations(v, trans)

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.ValidateTokenMiddleware)
	// e.Use(middlewares.ErrorsLogMiddleware)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler) // http://localhost:3000/swagger/index.html

	// Auth
	e.POST("/auth/register", handlers.HandlePostRegister)
	e.POST("/auth/login", handlers.HandlePostLogin)

	// Collections
	e.GET("/collections", handlers.HandleGetCollections)
	e.POST("/collection", handlers.HandlePostCollection)
	e.DELETE("/collection/:id", handlers.HandleDeleteCollection)

	// Categories
	e.POST("/category", handlers.HandlePostCategory)

	// Bookmarks
	e.POST("/bookmark", handlers.HandlePostBookmark)
	e.DELETE("/bookmark/:id", handlers.HandleDeleteBookmark)
	e.PUT("/bookmark/:id", handlers.HandleUpdateBookmark)

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_PORT")))

}
