package goinsta

import (
	"strconv"
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
	LoggedInUser UserInfo
}

/////////////////////
// Instagram types //
/////////////////////
type UserInfo struct {
	Username          string `json:"username"`
	ProfilePictureId  string `json:"profile_pic_id"`
	ProfilePictureURL string `json:"profile_pic_url"`
	FullName          string `json:"full_name"`
	PK                int64  `json:"pk"`
	IsVerified        bool   `json:"is_verified"`
	IsPrivate         bool   `json:"is_private"`
}

func (user UserInfo) StringID() {
	return strconv.FormatInt(user.PK, 10)
}
