package handlers

import (
	"net/http"
	"strings"
	"time"

	db "github.com/ebarquero85/link-backend/src/database"
	"github.com/ebarquero85/link-backend/src/messages"
	"github.com/ebarquero85/link-backend/src/models"
	"github.com/ebarquero85/link-backend/src/types"
	"github.com/ebarquero85/link-backend/src/utils"
	"github.com/ebarquero85/link-backend/src/validators"
	"github.com/labstack/echo/v4"
)

// @Summary Create Bookmark
// @Description
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param Body body types.BookmarkRequest true " "
// @Success 200 {object} map[string]interface{}
// @Router /bookmark [post]
// @Security Bearer
func HandlePostBookmark(c echo.Context) (err error) {

	bookmarkRequest := new(types.BookmarkRequest)

	if err = validators.Request(bookmarkRequest, c); err != nil {
		return err
	}

	bookmark_id, title_id, icon_id := getAllIds(c, bookmarkRequest)

	// getNoteId (Create Note)
	note_id := getNoteId(c, bookmarkRequest)

	// Create Final
	main := models.MainModel{
		UserId:       c.Get("UserId").(int),
		CreatedAt:    time.Now(),
		CollectionId: bookmarkRequest.CollectionId,
		BookmarkId:   bookmark_id,
		TitleId:      title_id,
		IconId:       icon_id, //sql.NullInt32{Int32: 0, Valid: false},
		NoteId:       note_id, //sql.NullInt32{},
	}

	if err = main.Create(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.JsonResponse[interface{}]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("BOOKMARK_CREATED"),
		Data:    nil,
	})

}

// @Summary Delete Bookmark
// @Description
// @Tags Bookmarks
// @Accept */*
// @Produce json
// @Param id path int true "Bookmark Id" required
// @Success 200 {object} map[string]interface{}
// @Router /bookmark/{id} [delete]
// @Security Bearer
func HandleDeleteBookmark(c echo.Context) (err error) {

	id := utils.GetParam(c, "id")

	user_id := c.Get("UserId").(int)

	db.Databases.DBPostgresql.Instance.Where("id = ? AND user_id = ?", id, user_id).Delete(&models.MainModel{})

	return c.JSON(http.StatusOK, types.JsonResponse[int]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("BOOKMARK_DELETED"),
		Data:    id,
	})

}

// @Summary Update Bookmark
// @Description
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param id path int true "Bookmark Id" required
// @Param cuerpo body types.BookmarkRequest true " "
// @Success 200 {object} map[string]interface{}
// @Router /bookmark/{id} [put]
// @Security Bearer
func HandleUpdateBookmark(c echo.Context) (err error) {

	id := utils.GetParam(c, "id")

	bookmarkRequest := new(types.BookmarkRequest)

	if err = validators.Request(bookmarkRequest, c); err != nil {
		return err
	}

	bookmark_id, title_id, icon_id := getAllIds(c, bookmarkRequest)

	user_id := c.Get("UserId").(int)

	main := models.MainModel{
		//UserId:       user_id,
		//CreatedAt:    time.Now(),
		CollectionId: bookmarkRequest.CollectionId,
		BookmarkId:   bookmark_id,
		TitleId:      title_id,
		IconId:       icon_id,
		//NoteId:     note_id,
	}

	if err = main.Update(id, user_id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Actualizar su nota

	// Primero buscamos el ID de la Nota.
	if err := db.Databases.DBPostgresql.Instance.First(&main, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Creamos la instancia con los datos actualizados.
	note := models.NoteModel{
		Note: bookmarkRequest.Note,
	}

	// Se actualiza.
	if err := note.Update(main.NoteId); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, types.JsonResponse[interface{}]{
		Status:  messages.SUCCESS,
		Message: messages.GetMessageTranslation("BOOKMARK_UPDATED"),
		Data:    nil,
	})

}

// Find or Create Bookmark
func getBookmarkId(c echo.Context, bookmarkRequest *types.BookmarkRequest) int {

	hash_bookmark := utils.CreateMd5Hash(bookmarkRequest.Url)

	bookmark := models.BookmarkModel{
		Hash:      hash_bookmark,
		CreatedAt: time.Now(),
		Url:       bookmarkRequest.Url,
	}

	result := db.Databases.DBPostgresql.Instance.FirstOrCreate(&bookmark, models.BookmarkModel{Hash: hash_bookmark})
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return bookmark.BookmarkId

}

// Find or Create Title
func getTitleId(c echo.Context, bookmarkRequest *types.BookmarkRequest) int {

	hash_title := utils.CreateMd5Hash(bookmarkRequest.Title)

	title := models.TitleModel{
		Hash:      hash_title,
		CreatedAt: time.Now(),
		Title:     strings.TrimSpace(bookmarkRequest.Title),
	}

	result := db.Databases.DBPostgresql.Instance.FirstOrCreate(&title, models.TitleModel{Hash: hash_title})
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return title.TitleId

}

// Find or Create Icon
func getIconId(c echo.Context, bookmarkRequest *types.BookmarkRequest) int {

	hash_icon := utils.CreateMd5Hash(bookmarkRequest.Icon)

	if hash_icon == "" {
		return 1
	}

	icon := models.IconModel{
		Hash: hash_icon,
		Icon: bookmarkRequest.Icon,
	}

	result := db.Databases.DBPostgresql.Instance.FirstOrCreate(&icon, models.IconModel{Hash: hash_icon})
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return icon.IconId

}

// Create Note.
func getNoteId(c echo.Context, bookmarkRequest *types.BookmarkRequest) int {

	nota_text := ""
	if bookmarkRequest.Note != "" {
		nota_text = bookmarkRequest.Note
	}

	note := models.NoteModel{
		Note: nota_text,
	}

	if err := note.Create(); err != nil {
		panic(err.Error())
	}

	return note.NoteId

}

func getAllIds(c echo.Context, bookmarkRequest *types.BookmarkRequest) (int, int, int) {

	// getBookmarkId
	bookmark_id := getBookmarkId(c, bookmarkRequest)

	// getTitleId
	title_id := getTitleId(c, bookmarkRequest)

	// getIconId
	icon_id := getIconId(c, bookmarkRequest)

	return bookmark_id, title_id, icon_id

}
