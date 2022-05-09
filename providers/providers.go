package providers

import "net/url"

type Adobe struct {
	EmailAdobe    string `json:"emailAdobe" validate:"email"`
	ClientIDAdobe string `json:"clientIDAdobe" validate:"required"`
}

type Facebook struct {
	EmailFacebook  string `json:"emailFacebook" validate:"email"`
	APIKeyFacebook string `json:"apiKeyFacebook" validate:"required"`
}

type Google struct {
	EmailGoogle    string `json:"emailGoogle" validate:"email"`
	ClientIDGoogle string `json:"clientIDGoogle" validate:"required"`
}

type IntelliBrands struct {
	EmailIntelliBrands    string `json:"emailIntelliBrands" validate:"email"`
	ClientIDIntelliBrands string `json:"clientIDIntelliBrands" validate:"required"`
}

type Microsoft struct {
	EmailMicrosoft    string `json:"emailMicrosoft" validate:"email"`
	ClientIDMicrosoft string `json:"ClientIDMicrosoft" validate:"required"`
}

type MySQL struct {
	HostMySQL     string `json:"hostMySQL" validate:"required"`
	DBNameMySQL   string `json:"dbNameMySQL" validate:"required"`
	PasswordMySQL string `json:"passwordMySQL"`
	UserMySQL     string `json:"userMySQL" validate:"required"`
}

type SQLite struct {
	DBNameSQLite string `json:"dbNameSQLite" validate:"required"`
}

type Twilio struct {
	EmailTwilio    string `json:"emailTwilio" validate:"email"`
	APIKeySendGrid string `json:"apiKeySendGrid" validate:"required"`
}

type Twitter struct {
	EmailTwitter  string `json:"emailTwitter" validate:"email"`
	APIKeyTwitter string `json:"apiKeyTwitter" validate:"required"`
}

type Vodafone struct {
	Number   uint64  `json:"number" validate:"required"`
	Password string  `json:"password"` // todo: size, regex
	Url      url.URL `json:"url"`
}
