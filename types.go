package goinsta

import (
	"net/http"
	"net/http/cookiejar"

	response "response"
)

type Informations struct {
	Username  string
	Password  string
	DeviceID  string
	UUID      string
	RankToken string
	Token     string
	PhoneID   string
}

type Instagram struct {
	Cookiejar *cookiejar.Jar
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
