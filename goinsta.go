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

// Instagram represent the main API handler
//
// Profiles: Represents instragram's user profile.
// Account:  Represents instagram's personal account.
// Search:   Represents instagram's search.
// Timeline: Represents instagram's timeline.
// Activity: Represents instagram's user activity.
// Inbox:    Represents instagram's messages.
//
// See Scheme section in README.md for more information.
//
// We recommend to use Export and Import functions after first Login.
//
// Also you can use SetProxy and UnsetProxy to set and unset proxy.
// Golang also provides the option to set a proxy using HTTP_PROXY env var.
type Instagram struct {
	user string
	pass string
	// device id
	dID string
	// uuid
	uuid string
	// rankToken
	rankToken string
	// token
	token string
	// phone id
	pid string

	// Instagram objects

	// Profiles is the user interaction
	Profiles *Profiles
	// Account stores all personal data of the user and his/her options.
	Account *Account
	// Search performs searching of multiple things (users, locations...)
	Search *Search
	// Timeline allows to receive timeline media.
	Timeline *Timeline
	// Activity are instagram notifications.
	Activity *Activity
	// Inbox are instagram message/chat system.
	Inbox *Inbox

	c *http.Client
}

// SetDeviceID sets device id
func (i *Instagram) SetDeviceID(id string) {
	i.dID = id
}

// SetUUID sets uuid
func (i *Instagram) SetUUID(uuid string) {
	i.uuid = uuid
}

// SetPhoneID sets phone id
func (i *Instagram) SetPhoneID(id string) {
	i.pid = id
}

// New creates Instagram structure
func New(username, password string) *Instagram {
	inst := &Instagram{
		user: username,
		pass: password,
		dID: generateDeviceID(
			generateMD5Hash(username + password),
		),
		uuid: generateUUID(), // both uuid must be differents
		pid:  generateUUID(),
		c: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}
	inst.init()

	return inst
}

func (inst *Instagram) init() {
	inst.Profiles = newProfiles(inst)
	inst.Activity = newActivity(inst)
	inst.Timeline = newTimeline(inst)
	inst.Search = newSearch(inst)
	inst.Inbox = newInbox(inst)
}

// SetProxy sets proxy for connection.
func (inst *Instagram) SetProxy(url string) error {
	uri, err := neturl.Parse(url)
	if err == nil {
		inst.c.Transport = &http.Transport{
			Proxy: http.ProxyURL(uri),
		}
	}
	return err
}

// UnsetProxy unsets proxy for connection.
func (inst *Instagram) UnsetProxy() {
	inst.c.Transport = nil
}

// Export exports *Instagram object options
func (inst *Instagram) Export(path string) error {
	url, err := neturl.Parse(goInstaAPIUrl)
	if err != nil {
		return err
	}

	config := ConfigFile{
		User:      inst.user,
		DeviceID:  inst.dID,
		UUID:      inst.uuid,
		RankToken: inst.rankToken,
		Token:     inst.token,
		PhoneID:   inst.pid,
		Cookies:   inst.c.Jar.Cookies(url),
	}
	bytes, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, bytes, 0644)
}

// Export exports *Instagram object into a json string
func (inst *Instagram) GetJson () (string, error) {
	url, err := neturl.Parse(goInstaAPIUrl)

	if err != nil {
		return "", err
	}

	config := ConfigFile{
		ID:        inst.Account.ID,
		User:      inst.user,
		DeviceID:  inst.dID,
		UUID:      inst.uuid,
		RankToken: inst.rankToken,
		Token:     inst.token,
		PhoneID:   inst.pid,
		Cookies:   inst.c.Jar.Cookies(url),
	}

	bytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}

// FromJson imports instagram configuration from a json string
//
// This function does not set proxy automatically. Use SetProxy after this call.
func FromJson(jsonSource string) (*Instagram, error){
	url, err := neturl.Parse(goInstaAPIUrl)
	if err != nil {
		return nil, err
	}

	config := ConfigFile{}

	err = json.Unmarshal([]byte(jsonSource), &config)

	if err != nil {
		return nil, err
	}

	inst := &Instagram{
		user:      config.User,
		dID:       config.DeviceID,
		uuid:      config.UUID,
		rankToken: config.RankToken,
		token:     config.Token,
		pid:       config.PhoneID,
		c: &http.Client {
			Transport: &http.Transport {
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}

	inst.c.Jar, err = cookiejar.New(nil)
	if err != nil {
		return inst, err
	}

	inst.c.Jar.SetCookies(url, config.Cookies)

	inst.init()
	inst.Account = &Account{inst: inst, ID: config.ID}
	inst.Account.Sync()

	return inst, nil
}

// Import imports instagram configuration
//
// This function does not set proxy automatically. Use SetProxy after this call.
func Import(path string) (*Instagram, error) {
	url, err := neturl.Parse(goInstaAPIUrl)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := ConfigFile{}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	inst := &Instagram{
		user:      config.User,
		dID:       config.DeviceID,
		uuid:      config.UUID,
		rankToken: config.RankToken,
		token:     config.Token,
		pid:       config.PhoneID,
		c: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}
	inst.c.Jar, err = cookiejar.New(nil)
	if err != nil {
		return inst, err
	}
	inst.c.Jar.SetCookies(url, config.Cookies)

	inst.init()
	inst.Account = &Account{inst: inst}
	inst.Account.Sync()

	return inst, nil
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
				"guid":           inst.uuid,
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
				err = instaToErr(ierr)
			}
			goto end
		}
		inst.Account = &res.Account
		inst.Account.inst = inst

		inst.rankToken = strconv.FormatInt(inst.Account.ID, 10) + "_" + inst.uuid

		inst.syncFeatures()
		inst.megaphoneLog()
	}

end:
	return err
}

// Logout closes current session
func (inst *Instagram) Logout() error {
	_, err := inst.sendSimpleRequest(urlLogout)
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
			Endpoint: urlSync,
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
			Endpoint: urlMegaphoneLog,
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
			Endpoint: urlExpose,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)

	return err
}

// AcquireFeed returns initilised FeedMedia
//
// Use FeedMedia.Sync() to update FeedMedia information. Do not forget to set id (you can use FeedMedia.SetID)
func (inst *Instagram) AcquireFeed() *FeedMedia {
	return &FeedMedia{inst: inst}
}
