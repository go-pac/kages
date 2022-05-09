package model

import (
	"github.com/go-pac/kages/providers"
)

const (
	DefaultDBMySQL = "false"
	EnvPort        = "PORT"
)

type ConfigAPI struct {
	ConfigAuth
	ConfigDB
	ConfigEmail
	Port uint16
}

type ConfigAuth struct {
	providers.Adobe
	providers.Facebook
	providers.Twitter
	providers.Microsoft
	providers.Google
}

type ConfigDB struct {
	providers.MySQL
	providers.SQLite
}

type ConfigEmail struct {
	providers.Twilio // todo: use providers.Google
}
