package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/messages"
	"github.com/ebarquero85/link-backend/src/models"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/ebarquero85/link-backend/src/validators"
	"github.com/labstack/echo/v4"
)

// @Summary Get Collections
// @Description
// @Tags Collections
// @Accept */*
// @Produce text/plain
// @Success 200
// @Router /collections [get]
// @Security Bearer
func HandleGetCollections(c echo.Context) (err error) {
	// Obtener el valor de UserId de la interfaz

	UserId := c.Get("UserId").(int)

	return c.JSON(http.StatusOK, types.JsonResponse[[]models.CollectionModel]{
		Status:  messages.SUCCESS,
		Message: "",
		Data:    GetCollections(UserId),
	})

}

// @Summary Create Collection
// @Description
// @Tags Collections
// @Accept json
// @Produce json
// @Param Body body types.CollectionRequest true " "
// @Success 200 {object} map[string]interface{}
// @Router /collection [post]
// @Security Bearer
func HandlePostCollection(c echo.Context) (err error) {

	collectionRequest := new(types.CollectionRequest)

	if err = validators.Request(collectionRequest, c); err != nil {
		return err
	}

	collection := models.CollectionModel{
		UserId: c.Get("UserId").(int),
		Name:   collectionRequest.Name,
	}

	// Create Collection
	if err = collection.Create(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.JsonResponse[models.CollectionModel]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("COLLECTION_CREATED"),
		Data:    collection,
	})

}

// @Summary Delete Collection
// @Description
// @Tags Collections
// @Accept */*
// @Produce json
// @Param id path int true "Collection Id" required
// @Success 200 {object} map[string]interface{}
// @Router /collection/{id} [delete]
// @Security Bearer
func HandleDeleteCollection(c echo.Context) (err error) {

	var (
		id            string
		collection_id int
	)

	id = c.Param("id")

	if collection_id, err = strconv.Atoi(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println(collection_id)

	return c.JSON(http.StatusOK, types.JsonResponse[int]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("COLLECTION_DELETED"),
		Data:    collection_id,
	})

}

func GetCollections(UserId int) []models.CollectionModel {

	var collections []models.CollectionModel

	db.Databases.DBPostgresql.Instance.Find(&collections, models.CollectionModel{UserId: UserId})

	return collections

}
