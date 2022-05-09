package model

type SMS struct {
	To   string `json:"to" validate:"required"`
	Body string `json:"body" validate:"required"`
}
