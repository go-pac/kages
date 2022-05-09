package model

type Notification struct {
	Model
	Description string `json:"description" validate:"min=3,max=256"`
	Name        string `json:"name" validate:"min=3,max=64"`
	Seen        bool   `json:"seen"`
	UserID      uint   `json:"userID" validate:"required"`
}

type NotificationDTO struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	CreatedByID uint   `json:"createdByID"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Seen        bool   `json:"seen"`
	UserID      uint   `json:"userID"`
}
