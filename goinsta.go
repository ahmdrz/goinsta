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
<<<<<<< HEAD
	information := Informations{
		DeviceID: generateDeviceID(generateMD5Hash(username + password)),
		Username: username,
		Password: password,
		UUID:     generateUUID(true),
		PhoneID:  generateUUID(true),
	}
	return &Instagram{
		InstaType: InstaType{
			Informations: information,
		},
	}
}

// Login to Instagram.
// return error if can't send request to instagram server
func (insta *Instagram) Login() error {
	insta.Cookiejar, _ = cookiejar.New(nil) //newJar()

	body, err := insta.sendRequest(&reqOptions{
		Endpoint:   "si/fetch_headers/",
		IsLoggedIn: true,
		Query: map[string]string{
			"challenge_type": "signup",
			"guid":           generateUUID(false),
		},
	})
	if err != nil {
		return fmt.Errorf("login failed for %s error %s", insta.Informations.Username, err.Error())
	}

	result, _ := json.Marshal(map[string]interface{}{
		"guid":                insta.Informations.UUID,
		"login_attempt_count": 0,
		"_csrftoken":          insta.Informations.Token,
		"device_id":           insta.Informations.DeviceID,
		"phone_id":            insta.Informations.PhoneID,
		"username":            insta.Informations.Username,
		"password":            insta.Informations.Password,
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
		LoggedInUser response.User `json:"logged_in_user"`
		Status       string        `json:"status"`
	}

	err = json.Unmarshal(body, &Result)
	if err != nil {
		return err
	}

	insta.LoggedInUser = Result.LoggedInUser
	insta.Informations.RankToken = strconv.FormatInt(Result.LoggedInUser.ID, 10) + "_" + insta.Informations.UUID
	insta.IsLoggedIn = true

	insta.SyncFeatures()
	insta.AutoCompleteUserList()
	insta.GetRankedRecipients()
	insta.Timeline("")
	insta.GetRankedRecipients()
	insta.GetRecentRecipients()
	insta.MegaphoneLog()

	return nil
}

// Logout of Instagram
func (insta *Instagram) Logout() error {
	_, err := insta.sendSimpleRequest("accounts/logout/")
	insta.Cookiejar = nil
	return err
}

// UserFollowing return followings of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowing(userID int64, maxID string) (response.UsersResponse, error) {
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/%d/following/", userID),
		Query: map[string]string{
			"max_id":             maxID,
			"ig_sig_key_version": GOINSTA_SIG_KEY_VERSION,
			"rank_token":         insta.Informations.RankToken,
		},
	})
	if err != nil {
		return response.UsersResponse{}, err
	}

	resp := response.UsersResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UserFollowers return followers of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowers(userID int64, maxID string) (response.UsersResponse, error) {
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/%d/followers/", userID),
		Query: map[string]string{
			"max_id":             maxID,
			"ig_sig_key_version": GOINSTA_SIG_KEY_VERSION,
			"rank_token":         insta.Informations.RankToken,
		},
	})
	if err != nil {
		return response.UsersResponse{}, err
	}

	resp := response.UsersResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// LatestFeed - Get the latest page of your own Instagram feed.
func (insta *Instagram) LatestFeed() (response.UserFeedResponse, error) {
	return insta.UserFeed(insta.LoggedInUser.ID, "", "")
}

// LatestUserFeed - Get the latest Instagram feed for the given user id
func (insta *Instagram) LatestUserFeed(userID int64) (response.UserFeedResponse, error) {
	return insta.UserFeed(userID, "", "")
}

// UserFeed - Returns the Instagram feed for the given user id.
// You can use maxID and minTimestamp for pagination, otherwise leave them empty to get the latest page only.
func (insta *Instagram) UserFeed(userID int64, maxID, minTimestamp string) (response.UserFeedResponse, error) {
	resp := response.UserFeedResponse{}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("feed/user/%d/", userID),
		Query: map[string]string{
			"max_id":         maxID,
			"rank_token":     insta.Informations.RankToken,
			"min_timestamp":  minTimestamp,
			"ranked_content": "true",
		},
	})
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UserTaggedFeed - Returns the feed for medua a given user is tagged in
func (insta *Instagram) UserTaggedFeed(userID, maxID int64, minTimestamp string) (response.UserTaggedFeedResponse, error) {
	resp := response.UserTaggedFeedResponse{}
	maxid := ""
	if maxID != 0 {
		maxid = string(maxID)
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("usertags/%d/feed/", userID),
		Query: map[string]string{
			"max_id":         maxid,
			"rank_token":     insta.Informations.RankToken,
			"min_timestamp":  minTimestamp,
			"ranked_content": "true",
		},
	})
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)

	return resp, err
}

// MediaComments - Returns comments of a media, input is mediaid of a media
// You can use maxID for pagination, otherwise leave it empty to get the latest page only.
func (insta *Instagram) MediaComments(mediaID string, maxID string) (response.MediaCommentsResponse, error) {
	resp := response.MediaCommentsResponse{}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/comments", mediaID),
		Query: map[string]string{
			"max_id": maxID,
=======
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
>>>>>>> v2
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

<<<<<<< HEAD
	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/edit_media/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) DeleteMedia(mediaID string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/delete/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) RemoveSelfTag(mediaID string) ([]byte, error) {
	data, err := insta.prepareData()
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/remove/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) Comment(mediaID, text string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"comment_text": text,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/comment/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) DeleteComment(mediaID, commentID string) ([]byte, error) {
	data, err := insta.prepareData()
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/comment/%s/delete/", mediaID, commentID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) GetRecentRecipients() ([]byte, error) {
	return insta.sendSimpleRequest("direct_share/recent_recipients/")
}

func (insta *Instagram) GetV2Inbox(cursor string) (response.DirectListResponse, error) {
	result := response.DirectListResponse{}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "direct_v2/inbox/",
		Query: map[string]string{
			"cursor": cursor,
		},
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetDirectPendingRequests() (response.DirectPendingRequests, error) {
	result := response.DirectPendingRequests{}
	body, err := insta.sendSimpleRequest("direct_v2/pending_inbox/?")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetRankedRecipients() (response.DirectRankedRecipients, error) {
	result := response.DirectRankedRecipients{}
	body, err := insta.sendSimpleRequest("direct_v2/ranked_recipients/?")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetDirectThread(threadid string) (response.DirectThread, error) {
	result := response.DirectThread{}
	body, err := insta.sendSimpleRequest("direct_v2/threads/%s/", threadid)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) Explore() (response.ExploreResponse, error) {
	result := response.ExploreResponse{}
	body, err := insta.sendSimpleRequest("discover/explore/")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
=======
end:
	return err
>>>>>>> v2
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
<<<<<<< HEAD
	})
}

// DirectMessage sends direct message to recipient.
// Recipient must be user id.
func (insta *Instagram) DirectMessage(recipient string, message string) (response.DirectMessageResponse, error) {
	result := response.DirectMessageResponse{}
	recipients, err := json.Marshal([][]string{{recipient}})
	if err != nil {
		return result, err
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary(insta.Informations.UUID)
	w.WriteField("recipient_users", string(recipients))
	w.WriteField("client_context", insta.Informations.UUID)
	w.WriteField("thread_ids", `["0"]`)
	w.WriteField("text", message)
	w.Close()

	req, err := http.NewRequest("POST", GOINSTA_API_URL+"direct_v2/threads/broadcast/text/", &b)
	if err != nil {
		return result, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-en")
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", GOINSTA_USER_AGENT)

	client := &http.Client{
		Jar: insta.Cookiejar,
	}

	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return result, fmt.Errorf(string(body))
	}

	json.Unmarshal(body, &result)
	return result, nil
}

// GetTrayFeeds - Get all available Instagram stories of your friends
func (insta *Instagram) GetReelsTrayFeed() (response.TrayResponse, error) {
	bytes, err := insta.sendSimpleRequest("feed/reels_tray/")
	if err != nil {
		return response.TrayResponse{}, err
	}

	result := response.TrayResponse{}
	json.Unmarshal([]byte(bytes), &result)

	return result, nil
}

// GetUserStories - Get all available Instagram stories for the given user id
func (insta *Instagram) GetUserStories(userID int64) (response.StoryResponse, error) {
	result := response.StoryResponse{}
	if userID == 0 {
		return result, nil
	}

	bytes, err := insta.sendSimpleRequest("feed/user/%d/reel_media/", userID)
	if err != nil {
		return result, err
=======
	)
	if err != nil {
		return err
>>>>>>> v2
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
