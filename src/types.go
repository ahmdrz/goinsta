package goinsta

import (
	response "github.com/ahmdrz/goinsta/src/response"
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
	IsLoggedIn   bool
	Informations Informations
	LoggedInUser response.User
}
