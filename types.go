package goinsta

import (
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

	IsLoggedIn   bool
	Informations Informations
	LoggedInUser response.User
}
