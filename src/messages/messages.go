package messages

import "github.com/ebarquero85/link-backend/src/config"

const (
	SUCCESS = "success"
	WARNING = "warning"
	ERROR   = "error"
)

var messages = map[string]map[string]string{
	"en": {
		"COLLECTION_CREATED":  "Created Collection",
		"USER_REGISTERED":     "Registered User",
		"FIRST_COLLECTION":    "First Collection",
		"CREDENTIALS_INVALID": "Invalid Credentials",
		"BOOKMARK_CREATED":    "Created Bookmark",
		"BOOKMARK_DELETED":    "Deleted Bookmark",
		"BOOKMARK_UPDATED":    "Updated Bookmark",
		"CATEGORY_CREATED":    "Created Category",
		"COLLECTION_DELETED":  "Deleted Collection",
	},
	"es": {
		"COLLECTION_CREATED":  "Colección Creada",
		"USER_REGISTERED":     "Usuario Registrado",
		"FIRST_COLLECTION":    "Primera Collección",
		"CREDENTIALS_INVALID": "Credenciales Inválidas",
		"BOOKMARK_CREATED":    "Marcador Creado",
		"BOOKMARK_DELETED":    "Marcador Borrado",
		"BOOKMARK_UPDATED":    "Marcador Actualizado",
		"CATEGORY_CREATED":    "Categoría Creada",
		"COLLECTION_DELETED":  "Colección Eliminada",
	},
}

func GetMessageTranslation(key string) string {
	var language map[string]string
	var ok bool
	var mensaje string

	if language, ok = messages[config.LANGUAGE]; !ok {
		language = messages["en"]
	}

	if mensaje, ok = language[key]; !ok {
		return "Translation: No key message found"
	}

	return mensaje
}
