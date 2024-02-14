package handlers

import (
	"fmt"
	"net/http"

	"github.com/ebarquero85/link-backend/src/config"
	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/messages"
	"github.com/ebarquero85/link-backend/src/models"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/ebarquero85/link-backend/src/utils"
	"github.com/ebarquero85/link-backend/src/validators"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	Token             string                   `json:"token"`
	Languaje          string                   `json:"language"`
	CollectionDefault int                      `json:"collection_default"`
	Collections       []models.CollectionModel `json:"collections"`
}

// @Summary Register
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body types.AuthRequest true "Request"
// @Success 200 {object} map[string]interface{}
// @Router /auth/register [post]
func HandlePostRegister(c echo.Context) (err error) {

	AuthRequest := new(types.AuthRequest)

	if err = validators.Request(AuthRequest, c); err != nil {
		return err
	}

	password := ""

	if password, err = utils.GeneratePasswordHash(AuthRequest.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user := models.UserModel{
		Name:      AuthRequest.Name,
		LastName:  AuthRequest.LastName,
		Email:     AuthRequest.Email,
		Password:  password,
		Confirmed: true, //config.DEFAULT_CONFIRMED,
		Language:  config.LANGUAGE,
	}

	// var user2 models.UserModel
	// Create User
	if err = user.Create(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(user)

	//create account

	account := models.AccountModel{
		UserId:   uint(user.UserId),
		Email:    AuthRequest.Email,
		Password: password,
	}

	if err = account.Create(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Create first Collection
	// collection := models.CollectionModel{
	// 	UserId: user.UserId,
	// 	Name:   messages.GetMessageTranslation("FIRST_COLLECTION"),
	// }

	// if err = collection.Create(); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	// Associate first collection to created user
	// db.Databases.DBPostgresql.Instance.Model(user).Update("collection_default", collection.CollectionId)

	// Generate token
	// token := "" //GenerateJWT(&user)

	return c.JSON(http.StatusOK, types.JsonResponse[string]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("USER_REGISTERED"),
		// Data:    token,
	})

}

// @Summary Login
// @Description
// @Tags Auth
// @Accept json
// @Produce json
// @Param Body body types.AuthRequest true " "
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func HandlePostLogin(c echo.Context) (err error) {

	AuthRequest := new(types.AuthRequest)
	user := new(models.UserModel)

	if err = validators.Request(AuthRequest, c); err != nil {
		return err
	}

	// Find User
	if err = db.Databases.DBPostgresql.Instance.Where("Email = ?", AuthRequest.Email).First(user).Error; err != nil {
		return c.JSON(http.StatusOK, types.JsonResponse[interface{}]{
			Status:  messages.WARNING,
			Message: messages.GetMessageTranslation("CREDENTIALS_INVALID"),
			Data:    nil,
		})
	}

	// Verify Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(AuthRequest.Password)); err != nil {
		return c.JSON(http.StatusOK, types.JsonResponse[interface{}]{
			Status:  messages.WARNING,
			Message: messages.GetMessageTranslation("CREDENTIALS_INVALID"),
			Data:    nil,
		})
	}

	// Generate token
	token := GenerateJWT(user)

	// Save Login
	login := models.LoginModel{
		UserId: user.UserId,
	}

	_ = login.Create() // No necestary check if error

	return c.JSON(http.StatusOK, types.JsonResponse[Data]{
		Status:  messages.SUCCESS,
		Message: "",
		Data: Data{
			Token:             token,
			Languaje:          user.Language,
			CollectionDefault: user.CollectionDefault,
			Collections:       GetCollections(user.UserId),
		},
	})

}
