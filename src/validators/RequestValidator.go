package validators

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Request[T comparable](data *T, c echo.Context) (err error) {

	if err = c.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(data); err != nil {

		fmt.Printf("%v", err)
		return err
	}

	return nil

}
