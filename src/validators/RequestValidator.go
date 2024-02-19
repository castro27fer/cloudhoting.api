package validators

import (
	"fmt"
	"net/http"

	translations "github.com/ebarquero85/link-backend/src/translations"
	en_translations "github.com/ebarquero85/link-backend/src/translations/en"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func Init_Request_validation() *CustomValidator {

	v := validator.New()

	en_translations.RegisterDefaultTranslations(v, translations.Get_translator())

	return &CustomValidator{validator: v} //descomentariar para activar las validaciones
}

func Request[T comparable](data *T, c echo.Context) (err error) {

	if err = c.Bind(data); err != nil {
		return err
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	return nil

}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {

		var errors2 []types.Error_Request

		trans2 := translations.Get_translator()
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
			fmt.Println("error", err.Translate(trans2))

		}

		return echo.NewHTTPError(http.StatusBadRequest, errors2)
	}

	return nil
}
