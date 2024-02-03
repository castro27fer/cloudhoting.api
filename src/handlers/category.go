package handlers

import (
	"net/http"

	"github.com/ebarquero85/link-backend/src/messages"
	"github.com/ebarquero85/link-backend/src/models"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/ebarquero85/link-backend/src/validators"
	"github.com/labstack/echo/v4"
)

// @Summary Create Category
// @Description
// @Tags Categories
// @Accept json
// @Produce json
// @Param Body body types.CategoryRequest true " "
// @Success 200 {object} map[string]interface{}
// @Router /category [post]
// @Security Bearer
func HandlePostCategory(c echo.Context) (err error) {

	categoryRequest := new(types.CategoryRequest)

	if err = validators.Request(categoryRequest, c); err != nil {
		return err
	}

	category := models.CategoryModel{
		UserId:       c.Get("UserId").(int),
		CollectionId: categoryRequest.CollectionId,
		Name:         categoryRequest.Name,
		Color:        categoryRequest.Color,
	}

	// Create Category
	if err = category.Create(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.JsonResponse[models.CategoryModel]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("CATEGORY_CREATED"),
		Data:    category,
	})

}
