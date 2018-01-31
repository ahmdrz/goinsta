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

// -----------------------------------------
// -----------------------------------------
// ----------- CURRENT USER METHODS --------
// -----------------------------------------
// -----------------------------------------

func (u *CurrentUser) Followers(options ...string) (FollowerAndFollowingResponse, error) {
	return u.instagram.FriendShip.Followers(u.ID, options...)
}

func (u *CurrentUser) Following(options ...string) (FollowerAndFollowingResponse, error) {
	return u.instagram.FriendShip.Following(u.ID, options...)
}

func (u *CurrentUser) Follow(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := u.instagram.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return
	}

	body, err := u.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/create/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return
	}

	var resp struct {
		FriendShipResponse FriendShipsShowResponse `json:"friendship_status"`
		Status             string                  `json:"status"`
	}
	err = json.Unmarshal(body, &resp)

	if resp.Status != "ok" {
		err = fmt.Errorf("bad status, %s", resp.Status)
		return
	}

	output = resp.FriendShipResponse

	return
}

func (u *CurrentUser) UnFollow(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := u.instagram.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return
	}

	body, err := u.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/destroy/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return
	}

	var resp struct {
		FriendShipResponse FriendShipsShowResponse `json:"friendship_status"`
		Status             string                  `json:"status"`
	}
	err = json.Unmarshal(body, &resp)

	if resp.Status != "ok" {
		err = fmt.Errorf("bad status, %s", resp.Status)
		return
	}

	output = resp.FriendShipResponse

	return
}

func (u *CurrentUser) Block(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := u.instagram.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return
	}

	body, err := u.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/block/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return
	}

	var resp struct {
		FriendShipResponse FriendShipsShowResponse `json:"friendship_status"`
		Status             string                  `json:"status"`
	}
	err = json.Unmarshal(body, &resp)

	if resp.Status != "ok" {
		err = fmt.Errorf("bad status, %s", resp.Status)
		return
	}

	output = resp.FriendShipResponse
	return
}

func (u *CurrentUser) UnBlock(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := u.instagram.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return
	}

	body, err := u.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/unblock/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return
	}

	var resp struct {
		FriendShipResponse FriendShipsShowResponse `json:"friendship_status"`
		Status             string                  `json:"status"`
	}
	err = json.Unmarshal(body, &resp)

	if resp.Status != "ok" {
		err = fmt.Errorf("bad status, %s", resp.Status)
		return
	}

	output = resp.FriendShipResponse
	return
}
