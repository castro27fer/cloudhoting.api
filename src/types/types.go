package types

type Error_Request struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type JsonResponse[T interface{}] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type AuthRequest struct {
	Email    string `json:"email" validate:"required,email,min=8,max=50" minLength:"8" maxLength:"50" example:"example@mail.com"`
	Password string `json:"password" validate:"required,min=6,max=24" minLength:"6" maxLength:"24" example:"nSjYMS9wEz"`
	Name     string `json:"name" validate:"required,min=3,max=40" example:"Jennifer"`
	LastName string `json:"lastName" validate:"required,min=3,max=40" example:"Zeledon"`
}

type CollectionRequest struct {
	Name string `json:"name" validate:"required,max=50" maxLength:"50" example:"Social Networks"`
}

type CategoryRequest struct {
	Name         string `json:"name" validate:"required,max=50" maxLength:"50" example:"Youtube Videos"`
	Color        string `json:"color" validate:"required,max=3" maxLength:"3" example:"blu"`
	CollectionId int    `json:"collection_id" validate:"required" example:"1"`
}

type BookmarkRequest struct {
	CollectionId int    `json:"collection_id" validate:"required" example:"10"`
	Title        string `json:"title" validate:"required,max=512" maxLength:"512" example:"The Go Programming Language"`
	Url          string `json:"url" validate:"required,max=2048" maxLength:"2048" example:"https://go.dev/"`
	Note         string `json:"note" validate:"max=1024" maxLength:"1024" example:"This is a good site for learning"`
	Icon         string `json:"icon" validate:"max=10000" maxLength:"10000" example:"xxxxx"`
}
