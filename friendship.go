package goinsta

import (
	"encoding/json"
	"fmt"
)

func (f *InstagramFriendShip) GetOne(userID int64) (output FriendShipsShowResponse, err error) {
	body, err := f.instagram.sendSimpleRequest(fmt.Sprintf("friendships/show/%d", userID))
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &output)
	return
}

func (f *InstagramFriendShip) GetMulti(userIDs ...int64) (output FriendShipsShowResponse, err error) {
	list := ""
	for i, id := range userIDs {
		list += fmt.Sprintf("%d", id)
		if i != len(userIDs)-1 {
			list += ","
		}
	}

	data, err := f.instagram.prepareData(map[string]interface{}{
		"user_ids": list,
	})

	fmt.Println(data)

	body, err := f.instagram.sendRequest(&reqOptions{
		Endpoint:     "friendships/show_many/",
		PostData:     generateSignature(data),
		IgnoreStatus: true,
	})
	if err != nil {
		return output, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &output)

	return
}

func (f *InstagramFriendShip) Pending() (output FriendShipsPendingResponse, err error) {
	body, err := f.instagram.sendSimpleRequest("friendships/pending/")
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &output)
	return
}

func (f *InstagramFriendShip) Approve(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := f.instagram.prepareData(map[string]interface{}{
		"user_id": fmt.Sprintf("%d", userID),
	})
	if err != nil {
		return
	}

	body, err := f.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/approve/%d/", userID),
		PostData: generateSignature(data),
	})

	var response struct {
		FriendshipStatus FriendShipsShowResponse `json:"friendship_status"`
		Status           string                  `json:"status"`
	}

	err = json.Unmarshal(body, &output)

	output = response.FriendshipStatus
	return
}

func (f *InstagramFriendShip) Reject(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := f.instagram.prepareData(map[string]interface{}{
		"user_id": fmt.Sprintf("%d", userID),
	})
	if err != nil {
		return
	}

	body, err := f.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/ignore/%d/", userID),
		PostData: generateSignature(data),
	})

	var response struct {
		FriendshipStatus FriendShipsShowResponse `json:"friendship_status"`
		Status           string                  `json:"status"`
	}

	err = json.Unmarshal(body, &output)

	output = response.FriendshipStatus
	return
}

func (f *InstagramFriendShip) RemoveFollower(userID int64) (output FriendShipsShowResponse, err error) {
	data, err := f.instagram.prepareData(map[string]interface{}{
		"user_id": fmt.Sprintf("%d", userID),
	})
	if err != nil {
		return
	}

	body, err := f.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/remove_follower/%d/", userID),
		PostData: generateSignature(data),
	})

	var response struct {
		FriendshipStatus FriendShipsShowResponse `json:"friendship_status"`
		Status           string                  `json:"status"`
	}

	err = json.Unmarshal(body, &output)

	output = response.FriendshipStatus
	return
}
