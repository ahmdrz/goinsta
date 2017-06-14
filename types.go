package goinsta

import (
	"net/http"
	"net/http/cookiejar"

	response "github.com/ahmdrz/goinsta/response"
)

type Informations struct {
	Username  string
	Password  string
	DeviceID  string
	UUID      string
	RankToken string
	Token     string
}

type Instagram struct {
	cookiejar *cookiejar.Jar
	InstaType
}

type InstaType struct {
	IsLoggedIn   bool
	Informations Informations
	LoggedInUser response.User

	proxy string
}

type BackupType struct {
	Cookies []http.Cookie
	InstaType
}
