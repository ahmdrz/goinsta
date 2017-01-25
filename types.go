package goinsta

import (
	response "github.com/ahmdrz/goinsta/response"
	"net/http/cookiejar"
)

type Informations struct {
	Username   string
	Password   string
	DeviceID   string
	UUID       string
	UsernameId string
	RankToken  string
	Token      string
}

type Instagram struct {
	cookie       string
	cookiejar    *cookiejar.Jar

	IsLoggedIn   bool
	Informations Informations
	LoggedInUser response.User
}
