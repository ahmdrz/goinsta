package goinst

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/cookiejar"
	"strconv"
	"time"
)

// New creates Instagram structure
func New(username, password string) (*Instagram, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	inst := &Instagram{
		user: username,
		pass: password,
		dID:  generateDeviceID(generateMD5Hash(username + password)),
		uuid: generateUUID(true),
		pid:  generateUUID(true),
		c: &http.Client{
			Jar: jar,
		},
	}

	inst.FriendShip = &FriendShip{
		inst: ist,
	}

	inst.Users = &Users{
		inst: ist,
	}

	return inst, err
}

func NewWithProxy(user, pass, url string) (*Instagram, error) {
	inst, err := New(user, pass)
	if err == nil {
		uri, err := url.Parse(url)
		if err == nil {
			inst.c.Transport = http.ProxyURL(uri)
		}
	}
	return inst, err
}

// ChangeTo logouts from the current account and login into another
func (inst *Instagram) ChangeTo(user, pass string) (err error) {
	inst.Logout()
	inst, err = inst.New(user, pass)
	if err == nil {
		err = inst.Login()
	}
	return
}

func (inst *Instagram) Export(path string) error {
	bytes, err := json.Marshal(map[string]interface{}{
		"uuid":         inst.uuid,
		"rank_token":   inst.rankToken,
		"token":        inst.token,
		"phone_id":     inst.phoneID,
		"device_id":    inst.deviceID,
		"proxy":        inst.proxy,
		"is_logged_in": inst.isLoggedIn,
		"cookie_jar":   inst.cookiejar,
	})
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, bytes, 0755)
}

func (inst *Instagram) Login() error {
	body, err := inst.sendRequest(&reqOptions{
		Endpoint: "si/fetch_headers/",
		Query: map[string]string{
			"challenge_type": "signup",
			"guid":           generateUUID(false),
		},
	})
	if err != nil {
		return fmt.Errorf("login failed for %s error %s", inst.username, err.Error())
	}

	result, _ := json.Marshal(map[string]interface{}{
		"guid":                inst.uuid,
		"login_attempt_count": 0,
		"_csrftoken":          inst.token,
		"device_id":           inst.deviceID,
		"phone_id":            inst.phoneID,
		"username":            inst.username,
		"password":            inst.password,
	})

	body, err = inst.sendRequest(&reqOptions{
		Endpoint:  "accounts/login/",
		QueryData: generateSignature(string(result)),
		IsPost:    true,
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

	inst.CurrentUser.UserResponse = Result.LoggedInUser
	inst.rankToken = strconv.FormatInt(Result.LoggedInUser.ID, 10) + "_" + inst.uuid
	inst.isLoggedIn = true

	inst.SyncFeatures()
	inst.FriendShip.AutoCompleteUserList()
	// inst.Timeline("")
	// inst.GetRankedRecipients()
	// inst.GetRecentRecipients()
	inst.MegaphoneLog()
	// inst.GetV2Inbox()
	// inst.GetRecentActivity()
	// inst.GetReelsTrayFeed()

	return nil
}

// Logout closes current session
func (inst *Instagram) Logout() error {
	_, err := inst.sendSimpleRequest("accounts/logout/")
	inst.c.Jar = nil
	inst.c = nil
	return err
}

// SyncFeatures simulates Instagram app behavior
func (inst *Instagram) SyncFeatures() error {
	data, err := inst.prepareData(
		map[string]interface{}{
			"id":          inst.CurrentUser.ID,
			"experiments": GOINSTA_EXPERIMENTS,
		},
	)
	if err != nil {
		return err
	}

	_, err = inst.sendRequest(&reqOptions{
		Endpoint:  "qe/sync/",
		QueryData: generateSignature(data),
		IsPost:    true,
	})
	return err
}

// MegaphoneLog simulates Instagram app behavior
func (inst *Instagram) MegaphoneLog() error {
	data, err := inst.prepareData(
		map[string]interface{}{
			"id":        inst.CurrentUser.ID,
			"type":      "feed_aysf",
			"action":    "seen",
			"reason":    "",
			"device_id": inst.deviceID,
			"uuid":      generateMD5Hash(string(time.Now().Unix())),
		},
	)
	if err != nil {
		return err
	}
	_, err = inst.sendRequest(&reqOptions{
		Endpoint:  "megaphone/log/",
		QueryData: generateSignature(data),
		IsPost:    true,
	})
	return err
}

// Expose , expose instgram
// return error if status was not 'ok' or runtime error
func (inst *Instagram) Expose() error {
	data, err := inst.prepareData(map[string]interface{}{
		"id":         inst.CurrentUser.ID,
		"experiment": "ig_android_profile_contextual_feed",
	})
	if err != nil {
		return err
	}

	_, err = inst.sendRequest(&reqOptions{
		Endpoint:  "qe/expose/",
		QueryData: generateSignature(data),
		IsPost:    true,
	})

	return err
}
