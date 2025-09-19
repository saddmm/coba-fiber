package dto

type CreatePostDto struct {
	UserId  uint   `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}


