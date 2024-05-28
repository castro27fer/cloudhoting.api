package handlers

import (
	"fmt"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/ebarquero85/link-backend/src/config"
	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/email"
	"github.com/ebarquero85/link-backend/src/messages"
	"github.com/ebarquero85/link-backend/src/models"
	translation "github.com/ebarquero85/link-backend/src/translations"
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

	//get request
	AuthRequest := c.Get("AuthRequest").(*types.AuthRequest)

	password := ""
	//hash password
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

	trans := translation.Get_translator()
	text, _ := trans.T("user_register")

	return c.JSON(http.StatusOK, types.JsonResponse[string]{
		Status:  messages.SUCCESS,
		Message: text,
		// Data:    token,
	})
}

func HandleCodeVerify(c echo.Context) (err error) {

	//get request
	AuthRequest := new(types.CodeVerify)

	fmt.Print(AuthRequest)
	if err := validators.Request(AuthRequest, c); err != nil {
		// return err
		fmt.Print("Errors of validate: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	var codes []models.CodeVerifyModel

	//get all codes send with the email
	db.Databases.DBPostgresql.Instance.Where("email=?", AuthRequest.Email).Find(&codes)

	if len(codes) > 0 {

		//update statu code a Cancel
		for _, code := range codes {

			code.Status = "Cancel"
			code.Update()

		}

	}

	//generate code of six digito
	code := utils.GenerateRandomNumber()

	//create new code
	codeVerify := models.CodeVerifyModel{
		Email:  AuthRequest.Email,
		Code:   strconv.Itoa(code),
		Status: "NotVerify",
	}

	if err = codeVerify.Create(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var expired_code string = os.Getenv("CODE_VERIFY_EXPIRATION")

	// Convertir el string a time.Duration
	duration, err := time.ParseDuration(expired_code)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	//expire code
	time.AfterFunc(duration, func() {

		codeVerify.Status = "Expired"
		codeVerify.Update()
	})

	//send email with code activate

	fmt.Printf("Nombres: %v\n", AuthRequest.Names)
	if err = email.SendActivationEmail(mail.Address{
		Name:    AuthRequest.Names,
		Address: AuthRequest.Email,
	}, codeVerify.Code); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	trans := translation.Get_translator()
	text, _ := trans.T("send_code")

	return c.JSON(http.StatusOK, types.JsonResponse[string]{
		Status:  messages.SUCCESS,
		Message: text,
		// Data:    token,
	})

}

func GenerateRandomNumber() {
	panic("unimplemented")
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

	AuthRequest := new(types.LoginRequest)
	user := new(models.UserModel)

	if err = validators.Request(AuthRequest, c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Find User
	if err = db.Databases.DBPostgresql.Instance.Where("Email = ?", AuthRequest.Email).First(user).Error; err != nil {
		fmt.Println("error en el email: ", err)
		return c.JSON(http.StatusOK, types.JsonResponse[interface{}]{
			Status:  messages.WARNING,
			Message: messages.GetMessageTranslation("CREDENTIALS_INVALID"),
			Data:    nil,
		})
	}

	// Verify Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(AuthRequest.Password)); err != nil {

		fmt.Println("error en la contraseña: ", err)
		return c.JSON(http.StatusBadRequest, &validators.ValidationError{
			Status:      http.StatusBadRequest,
			Message:     messages.GetMessageTranslation("CREDENTIALS_INVALID"),
			Validations: []types.Error_Request{},
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
