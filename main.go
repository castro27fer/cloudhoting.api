package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/ebarquero85/link-backend/docs/links" // Necesario para que funcione Swagger

	"github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/database/migration"
	"github.com/ebarquero85/link-backend/src/middlewares"
	"github.com/ebarquero85/link-backend/src/routes"
	"github.com/ebarquero85/link-backend/src/services"

	translation "github.com/ebarquero85/link-backend/src/translations"
	requestValidation "github.com/ebarquero85/link-backend/src/validators"
)

func init() {

	// Load .env configuration
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	base := database.Connect("postgres")
	if err := migration.Init(base, false); err != nil {
		panic(err.Error())
	}

	//instance rolAndPermissions service

	// database.Databases.DBPostgresql.Instance.Migrator().AutoMigrate(&auth.UserModel{}, &auth.AccountModel{}, &auth.CodeVerifyModel{}, &auth.LoginModel{})
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

	services.ProfilesServices.Start()
	services.EmailService.Start()

	e.Validator = requestValidation.Init_Request_validation()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middlewares.ValidateTokenMiddleware)
	e.Use(middlewares.LanguageUser)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler) // http://localhost:3000/swagger/index.html

	routes.Router(e)

	e.Logger.Fatal(e.Start(os.Getenv("LISTEN_PORT")))

}
