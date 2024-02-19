package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/ebarquero85/link-backend/docs/links" // Necesario para que funcione Swagger

	"github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/handlers"
	"github.com/ebarquero85/link-backend/src/middlewares"

	translation "github.com/ebarquero85/link-backend/src/translations"
	requestValidation "github.com/ebarquero85/link-backend/src/validators"
)

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

	translation.Init_translate_default()
	translation.Load_languages()

	e.Validator = requestValidation.Init_Request_validation()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.ValidateTokenMiddleware)
	e.Use(middlewares.LanguageUser)
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
