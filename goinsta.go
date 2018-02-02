package goinsta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/cookiejar"
	"strconv"
	"time"
)

// New returns pointer to Instagram structure.
func New(username, password string) *Instagram {
	instagram := &Instagram{
		username: username,
		password: password,
		deviceID: generateDeviceID(generateMD5Hash(username + password)),
		uuid:     generateUUID(true),
		phoneID:  generateUUID(true),
	}

	instagram.FriendShip = &FriendShip{
		instagram: instagram,
	}

	instagram.Users = &Users{
		instagram: instagram,
	}

	return instagram
}

func (i *Instagram) SetUser(user string) {
	i.username = user
}

func (i *Instagram) SetPass(pass string) {
	i.password = pass
}

func (insta *Instagram) Export(path string) error {
	bytes, err := json.Marshal(map[string]interface{}{
		"uuid":         insta.uuid,
		"rank_token":   insta.rankToken,
		"token":        insta.token,
		"phone_id":     insta.phoneID,
		"device_id":    insta.deviceID,
		"proxy":        insta.proxy,
		"is_logged_in": insta.isLoggedIn,
		"cookie_jar":   insta.cookiejar,
	})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0755)
}

func (insta *Instagram) Login() error {
	insta.cookiejar, _ = cookiejar.New(nil)

	body, err := insta.sendRequest(&reqOptions{
		Endpoint:   "si/fetch_headers/",
		IsLoggedIn: true,
		Query: map[string]string{
			"challenge_type": "signup",
			"guid":           generateUUID(false),
		},
	})
	if err != nil {
		return fmt.Errorf("login failed for %s error %s", insta.username, err.Error())
	}

	result, _ := json.Marshal(map[string]interface{}{
		"guid":                insta.uuid,
		"login_attempt_count": 0,
		"_csrftoken":          insta.token,
		"device_id":           insta.deviceID,
		"phone_id":            insta.phoneID,
		"username":            insta.username,
		"password":            insta.password,
	})

	body, err = insta.sendRequest(&reqOptions{
		Endpoint:   "accounts/login/",
		PostData:   generateSignature(string(result)),
		IsLoggedIn: true,
	})
	if err != nil {
		return err
	}

	var Result struct {
		LoggedInUser UserResponse `json:"logged_in_user"`
		Status       string       `json:"status"`
	}

	err = json.Unmarshal(body, &Result)
	if err != nil {
		return err
	}

	insta.CurrentUser.UserResponse = Result.LoggedInUser
	insta.rankToken = strconv.FormatInt(Result.LoggedInUser.ID, 10) + "_" + insta.uuid
	insta.isLoggedIn = true

	insta.SyncFeatures()
	insta.FriendShip.AutoCompleteUserList()
	// insta.Timeline("")
	// insta.GetRankedRecipients()
	// insta.GetRecentRecipients()
	insta.MegaphoneLog()
	// insta.GetV2Inbox()
	// insta.GetRecentActivity()
	// insta.GetReelsTrayFeed()

	return nil
}

// Logout of Instagram
func (insta *Instagram) Logout() error {
	_, err := insta.sendSimpleRequest("accounts/logout/")
	insta.cookiejar = nil
	return err
}

// SyncFeatures simulates Instagram app behavior
func (insta *Instagram) SyncFeatures() error {
	data, err := insta.prepareData(map[string]interface{}{
		"id":          insta.CurrentUser.ID,
		"experiments": GOINSTA_EXPERIMENTS,
	})
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(&reqOptions{
		Endpoint: "qe/sync/",
		PostData: generateSignature(data),
	})
	return err
}

// MegaphoneLog simulates Instagram app behavior
func (insta *Instagram) MegaphoneLog() error {
	data, err := insta.prepareData(map[string]interface{}{
		"id":        insta.CurrentUser.ID,
		"type":      "feed_aysf",
		"action":    "seen",
		"reason":    "",
		"device_id": insta.deviceID,
		"uuid":      generateMD5Hash(string(time.Now().Unix())),
	})
	if err != nil {
		return err
	}
	_, err = insta.sendRequest(&reqOptions{
		Endpoint: "megaphone/log/",
		PostData: generateSignature(data),
	})
	return err
}

// Expose , expose instagram
// return error if status was not 'ok' or runtime error
func (insta *Instagram) Expose() error {
	data, err := insta.prepareData(map[string]interface{}{
		"id":         insta.CurrentUser.ID,
		"experiment": "ig_android_profile_contextual_feed",
	})
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(&reqOptions{
		Endpoint: "qe/expose/",
		PostData: generateSignature(data),
	})

	return err
}
