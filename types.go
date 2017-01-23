package goinsta

import (
	response "github.com/ahmdrz/goinsta/response"
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
	lastJson []byte

	IsLoggedIn   bool
	Informations Informations
	LoggedInUser response.User
}
