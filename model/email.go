package model

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Email struct {
	From    *mail.Email
	HTML    string
	Subject string
	Text    string
	To      *mail.Email
}
