package model

import (
	"strings"
	"time"
)

const (
	RoleAdmin     = "admin"
	RoleClient    = "client"
	RoleModerator = "moderator"
	RoleService   = "service"
)

type User struct {
	Model
	Address       string         `json:"address,omitempty" validate:"max=128"`
	AuthID        string         `json:"authID,omitempty"`
	AuthProvider  string         `json:"authProvider,omitempty"`
	Dob           time.Time      `json:"dob,omitempty"` // todo: date validation
	FirstName     string         `json:"firstName" validate:"min=1,max=64"`
	LastName      string         `json:"lastName" validate:"min=2,max=64"`
	Email         string         `json:"email,omitempty" gorm:"uniqueIndex,size:255" validate:"email,max=64"` // todo: required if no mobile
	History       []Login        `json:"history,omitempty" gorm:"foreignKey:CreatedByID"`
	IdentityNo    string         `json:"identityNo,omitempty"` // todo: gorm:"unique"
	LastSeenAt    time.Time      `json:"lastSeenAt,omitempty"`
	Mobile        string         `json:"mobile" gorm:"unique" validate:"numeric"` // todo: required if no email
	Notifications []Notification `json:"notifications,omitempty"`
	Password      string         `json:"password,omitempty"`
	Photo         string         `json:"photo,omitempty"`
	Recovery      string         `json:"recovery,omitempty"` // todo validate:"numeric,len=6"
	Scopes        string         `json:"scopes"`
}

type UserDTO struct {
	ID           uint      `json:"id"`
	Address      string    `json:"address,omitempty"`
	AuthID       string    `json:"authID,omitempty"`
	AuthProvider string    `json:"authProvider,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	Dob          time.Time `json:"dob,omitempty"`
	Email        string    `json:"email,omitempty"`
	FirstName    string    `json:"firstName"`
	IdentityNo   string    `json:"identityNo,omitempty"`
	LastName     string    `json:"lastName"`
	LastSeenAt   time.Time `json:"lastSeenAt"`
	Mobile       string    `json:"mobile"`
	Photo        string    `json:"photo,omitempty"`
	Scopes       string    `json:"scopes"`
}

func (u *User) AuthType() string {
	if len(u.AuthProvider) > 0 {
		return u.AuthProvider
	}

	return "credentials"
}

func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) Is(role string) bool {
	return strings.Contains(u.Scopes, role)
}

func (u *User) IsBasic() bool {
	return !strings.Contains(u.Scopes, RoleAdmin) && !strings.Contains(u.Scopes, RoleModerator)
}
