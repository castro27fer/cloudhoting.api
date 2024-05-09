package validators

import (
	"fmt"
	"net/http"

	translations "github.com/ebarquero85/link-backend/src/translations"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func Init_Request_validation() *CustomValidator {

	v := validator.New()

	return &CustomValidator{validator: v} //descomentariar para activar las validaciones
}

func Request[T comparable](data *T, c echo.Context) (err error) {

	if err = c.Bind(data); err != nil {

		fmt.Print("error en bind \n")
		return err
	}

	if err = c.Validate(data); err != nil {
		fmt.Print("error en validate")
		return err
	}

	return nil

}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {

		var errors2 []types.Error_Request

		trans2 := translations.Get_translator()

		message, found_message := trans2.T("bad_request")
		if found_message != nil {
			message = "There are incomplete fields"
		}

		for _, err := range err.(validator.ValidationErrors) {

			name := err.Field()
			fieldName, found_field := trans2.T(name)
			if found_field != nil {
				fieldName = name
			}

			text, found := trans2.T(err.Tag(), fieldName, err.Param())
			if found != nil {
				text = err.Translate(trans2)
			}

			errors2 = append(errors2, types.Error_Request{
				Name:    name,
				Message: text,
			})

		}

		return echo.NewHTTPError(http.StatusBadRequest, struct {
			Message     string                `json:"message"`
			Validations []types.Error_Request `json:"validations"`
		}{
			Message:     message,
			Validations: errors2,
		})
	}

	return nil
}
