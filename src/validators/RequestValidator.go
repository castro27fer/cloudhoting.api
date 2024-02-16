package validators

import (
	"github.com/labstack/echo/v4"
)

func Request[T comparable](data *T, c echo.Context) (err error) {

	if err = c.Bind(data); err != nil {
		return err
	}

	if err = c.Validate(data); err != nil {
		return err
	}

	return nil

}
