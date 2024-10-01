package routes

import (
	"github.com/ebarquero85/link-backend/src/handlers"

	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {

	// Auth
	e.POST("/auth/register", handlers.HandlePostRegister)

	e.POST("/auth/login", handlers.HandlePostLogin)
	e.POST("/auth/codeVerify", handlers.HandleCodeVerify)
}
