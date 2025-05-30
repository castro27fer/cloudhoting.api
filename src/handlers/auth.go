package handlers

import (
	"math/rand"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"time"

	"github.com/ebarquero85/link-backend/src/config"
	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/email"
	"github.com/ebarquero85/link-backend/src/messages"
	models "github.com/ebarquero85/link-backend/src/models/auth"
	services "github.com/ebarquero85/link-backend/src/services"
	translation "github.com/ebarquero85/link-backend/src/translations"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/ebarquero85/link-backend/src/validators"
	"github.com/labstack/echo/v4"
)

type Data struct {
	Token       string               `json:"token"`
	Language    string               `json:"language"`
	Permissions []*models.Permission `json:"permissions"`
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
	data := new(types.AuthRequest)

	if err = validators.Request(data, c); err != nil {
		return err
	}

	profile := services.ProfilesServices.Get_rol_default()

	user := services.User{
		UserModel: models.UserModel{
			Name:        data.Name,
			LastName:    data.LastName,
			Email:       data.Email,
			Password:    data.Password,
			CountryID:   data.CountryId,
			Confirmed:   false,
			Language:    config.LANGUAGE,
			ProfileType: profile,
		},
	}

	// Create User
	if err = user.Create(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//generate token of activation
	activation_account_token := services.JWTService.ActiveAccount.CreateToken(&user)

	//send email of activation
	email_error := services.EmailService.SendEmail(mail.Address{Name: user.FullName(), Address: user.Email}, "account activation", activation_account_token)
	if email_error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, email_error.Error())
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
	data := new(types.CodeVerify)

	if err := validators.Request(data, c); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var codes []models.CodeVerifyModel

	//get all codes send with the email
	db.Databases.DBPostgresql.Instance.Where("email=?", data.Email).Find(&codes)

	if len(codes) > 0 {

		//update statu code a Cancel
		for _, code := range codes {

			code.Status = "Cancel"
			code.Update()

		}

	}

	//generate code of six digit
	code := rand.Intn(999999-100000) + 100000

	//create new code
	codeVerify := models.CodeVerifyModel{
		Email:  data.Email,
		Code:   strconv.Itoa(code),
		Status: "NotVerify",
	}

	if err = codeVerify.Create(); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	var expired_code string = os.Getenv("CODE_VERIFY_EXPIRATION")

	// Convert string to time.Duration
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
	trans := translation.Get_translator()

	if err = email.SendActivationEmail(mail.Address{
		Name:    data.Names,
		Address: data.Email,
	}, codeVerify.Code); err != nil {

		// error_internal, _ := trans.T("internal_server_error")

		// return c.JSON(http.StatusInternalServerError, validators.ValidationError{
		// 	Status:      http.StatusInternalServerError,
		// 	Message:     error_internal,
		// 	Validations: []types.Error_Request{},
		// })

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

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

	// 	LoginRequest := new(types.LoginRequest)
	// 	account := new(models.AccountModel)

	// 	if err = validators.Request(LoginRequest, c); err != nil {
	// 		return c.JSON(http.StatusBadRequest, err)
	// 	}

	// 	// Find User
	// 	if err = db.Databases.DBPostgresql.Instance.Where("Email = ?", LoginRequest.Email).First(account).Error; err != nil {
	// 		fmt.Println("error en el email: ", err)
	// 		return c.JSON(http.StatusOK, types.JsonResponse[interface{}]{
	// 			Status:  messages.WARNING,
	// 			Message: messages.GetMessageTranslation("CREDENTIALS_INVALID"),
	// 			Data:    nil,
	// 		})
	// 	}

	// 	// Verify Password
	// 	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(LoginRequest.Password)); err != nil {

	// 		fmt.Println("error en la contraseña: ", err)
	// 		return c.JSON(http.StatusBadRequest, &validators.ValidationError{
	// 			Status:      http.StatusBadRequest,
	// 			Message:     messages.GetMessageTranslation("CREDENTIALS_INVALID"),
	// 			Validations: []types.Error_Request{},
	// 		})
	// 	}

	// 	// Generate token
	// 	token := GenerateJWT(account)

	// 	// Save Login
	// 	login := models.LoginModel{
	// 		AccountId: account.ID,
	// 		Token:     token,
	// 		IP:        LoginRequest.IP,
	// 	}

	// 	_ = login.Create() // No need check if error

	// 	return c.JSON(http.StatusOK, types.JsonResponse[Data]{
	// 		Status:  messages.SUCCESS,
	// 		Message: "",
	// 		Data: Data{
	// 			Token:       token,
	// 			Language:    account.Language,
	// 			Permissions: account.Permissions,
	// 		},
	// 	})

	return nil
}
