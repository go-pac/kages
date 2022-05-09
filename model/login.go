package model

type Login struct {
	Model
	Client string `json:"client" validate:"required"` // todo: size
	IP     string `json:"ip" validate:"required"`     // todo: regex
	Type   string `json:"type" validate:"required"`   // todo: enum
}
