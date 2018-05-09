package goinsta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	neturl "net/url"
	"strconv"
	"time"
)

// New creates Instagram structure
func New(username, password string) *Instagram {
	inst := &Instagram{
		user: username,
		pass: password,
		dID: generateDeviceID(
			generateMD5Hash(username + password),
		),
		uuid: generateUUID(true),
		pid:  generateUUID(true),
		c:    &http.Client{},
	}

	inst.Users = newUsers(inst)
	// not needed
	// this object is created after login
	// inst.Account = NewAccount(inst)

	return inst
}

func NewWithProxy(user, pass, url string) (*Instagram, error) {
	inst := New(user, pass)
	uri, err := neturl.Parse(url)
	_ = uri
	if err == nil {
		// TODO
		//inst.c.Transport = proxhttp.ProxyURL(uri)
	}
	return inst, err
}

// ChangeTo logouts from the current account and login into another
func (inst *Instagram) ChangeTo(user, pass string) (err error) {
	inst.Logout()
	inst = New(user, pass)
	return inst.Login()
}

// Export ...
// TODO: Import and export (in other good readable format)
func (inst *Instagram) Export(path string) error {
	bytes, err := json.Marshal(
		map[string]interface{}{
			"uuid":       inst.uuid,
			"rank_token": inst.rankToken,
			"token":      inst.token,
			"phone_id":   inst.pid,
			"device_id":  inst.dID,
			"client":     inst.c,
		})
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, 0755)
}

// Login performs instagram login.
//
// Password will be deleted after login
func (inst *Instagram) Login() error {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	inst.c.Jar = jar

	body, err := inst.sendRequest(
		&reqOptions{
			Endpoint: urlFetchHeaders,
			Query: map[string]string{
				"challenge_type": "signup",
				"guid":           generateUUID(false),
			},
		},
	)
	if err != nil {
		return fmt.Errorf("login failed for %s: %s", inst.user, err.Error())
	}

	result, err := json.Marshal(
		map[string]interface{}{
			"guid":                inst.uuid,
			"login_attempt_count": 0,
			"_csrftoken":          inst.token,
			"device_id":           inst.dID,
			"phone_id":            inst.pid,
			"username":            inst.user,
			"password":            inst.pass,
		},
	)
	if err == nil {
		body, err = inst.sendRequest(
			&reqOptions{
				Endpoint: urlLogin,
				Query:    generateSignature(b2s(result)),
				IsPost:   true,
			},
		)
		if err != nil {
			goto end
		}
		inst.pass = ""

		// getting account data
		res := accountResp{}

		err = json.Unmarshal(body, &res)
		if err != nil {
			ierr := instaError{}
			err = json.Unmarshal(body, &ierr)
			if err != nil {
				err = errToInstagram(ierr)
			}
			return err
		}
		inst.Account = &res.Account
		inst.Account.inst = inst

		inst.rankToken = strconv.FormatInt(inst.Account.ID, 10) + "_" + inst.uuid
		inst.logged = true

		inst.syncFeatures()
		// inst.Timeline("")
		// inst.GetRankedRecipients()
		// inst.GetRecentRecipients()
		inst.megaphoneLog()
		// inst.GetV2Inbox()
		// inst.GetRecentActivity()
		// inst.GetReelsTrayFeed()
	}

end:
	return err
}

// Logout closes current session
func (inst *Instagram) Logout() error {
	_, err := inst.sendSimpleRequest("accounts/logout/")
	inst.logged = false
	inst.c.Jar = nil
	inst.c = nil
	return err
}

func (inst *Instagram) syncFeatures() error {
	data, err := inst.prepareData(
		map[string]interface{}{
			"id":          inst.Account.ID,
			"experiments": goInstaExperiments,
		},
	)
	if err != nil {
		return err
	}

	_, err = inst.sendRequest(
		&reqOptions{
			Endpoint: "qe/sync/",
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err != nil {
		return err
	}

	_, err = inst.sendRequest(
		&reqOptions{
			Endpoint: urlAutoComplete,
			Query: map[string]string{
				"version": "2",
			},
		},
	)
	return err
}

func (inst *Instagram) megaphoneLog() error {
	data, err := inst.prepareData(
		map[string]interface{}{
			"id":        inst.Account.ID,
			"type":      "feed_aysf",
			"action":    "seen",
			"reason":    "",
			"device_id": inst.dID,
			"uuid":      generateMD5Hash(string(time.Now().Unix())),
		},
	)
	if err != nil {
		return err
	}
	_, err = inst.sendRequest(
		&reqOptions{
			Endpoint: "megaphone/log/",
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

func (inst *Instagram) expose() error {
	data, err := inst.prepareData(
		map[string]interface{}{
			"id":         inst.Account.ID,
			"experiment": "ig_android_profile_contextual_feed",
		},
	)
	if err != nil {
		return err
	}

	_, err = inst.sendRequest(
		&reqOptions{
			Endpoint: "qe/expose/",
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)

	return err
}
