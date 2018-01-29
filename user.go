package goinsta

import (
	"encoding/json"
	"fmt"
)

func (u *InstagramUsers) ByUsername(username string) (user UserResponse, err error) {
	body, err := u.instagram.sendSimpleRequest("users/%s/usernameinfo/", username)
	if err != nil {
		return
	}

	var userSearchResponse struct {
		User   UserResponse `json:"user"`
		Status string       `json:"status"`
	}
	err = json.Unmarshal(body, &userSearchResponse)

	if userSearchResponse.Status != "ok" {
		err = fmt.Errorf("bad status, %s", userSearchResponse.Status)
		return
	}

	user = userSearchResponse.User
	return
}

func (u *InstagramUsers) ByID(userID int64) (user UserResponse, err error) {
	data, err := u.instagram.prepareData()
	if err != nil {
		return
	}

	body, err := u.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("users/%d/info/", userID),
		PostData: generateSignature(data),
	})

	var userSearchResponse struct {
		User   UserResponse `json:"user"`
		Status string       `json:"status"`
	}
	err = json.Unmarshal(body, &userSearchResponse)

	if userSearchResponse.Status != "ok" {
		err = fmt.Errorf("bad status, %s", userSearchResponse.Status)
		return
	}

	user = userSearchResponse.User
	return
}
