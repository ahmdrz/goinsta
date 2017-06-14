// goinsta project goinsta.go
package goinsta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"net/http/cookiejar"

	"github.com/ahmdrz/goinsta/response"
)

// GetSessions return current instagram session and cookies
// Maybe need for webpages that use this API
func (insta *Instagram) GetSessions(url *url.URL) []*http.Cookie {
	return insta.cookiejar.Cookies(url)
}

// SetCookies can enable us to set cookie, it'll be help for webpage that use this API without Login-again.
func (insta *Instagram) SetCookies(url *url.URL, cookies []*http.Cookie) error {
	if insta.cookiejar == nil {
		var err error
		insta.cookiejar, err = cookiejar.New(nil) //newJar()
		if err != nil {
			return err
		}
	}
	insta.cookiejar.SetCookies(url, cookies)
	return nil
}

// Const values ,
// GOINSTA Default variables contains API url , user agent and etc...
// GOINSTA_IG_SIG_KEY is Instagram sign key, It's important
// Filter_<name>
const (
	Filter_Walden           = 20
	Filter_Crema            = 616
	Filter_Reyes            = 614
	Filter_Moon             = 111
	Filter_Ashby            = 116
	Filter_Maven            = 118
	Filter_Brannan          = 22
	Filter_Hefe             = 21
	Filter_Valencia         = 25
	Filter_Clarendon        = 112
	Filter_Helena           = 117
	Filter_Brooklyn         = 115
	Filter_Dogpatch         = 105
	Filter_Ludwig           = 603
	Filter_Stinson          = 109
	Filter_Inkwell          = 10
	Filter_Rise             = 23
	Filter_Perpetua         = 608
	Filter_Juno             = 613
	Filter_Charmes          = 108
	Filter_Ginza            = 107
	Filter_Hudson           = 26
	Filter_Normat           = 0
	Filter_Slumber          = 605
	Filter_Lark             = 615
	Filter_Skyline          = 113
	Filter_Kelvin           = 16
	Filter_1977             = 14
	Filter_Lo_Fi            = 2
	Filter_Aden             = 612
	Filter_Amaro            = 24
	Filter_Sutro            = 18
	Filter_Vasper           = 106
	Filter_Nashville        = 15
	Filter_X_Pro_II         = 1
	Filter_Mayfair          = 17
	Filter_Toaster          = 19
	Filter_Earlybird        = 3
	Filter_Willow           = 28
	Filter_Sierra           = 27
	Filter_Gingham          = 114
	GOINSTA_API_URL         = "https://i.instagram.com/api/v1/"
	GOINSTA_USER_AGENT      = "Instagram 10.15.0 Android (18/4.3; 320dpi; 720x1280; Xiaomi; HM 1SW; armani; qcom; en_US)"
	GOINSTA_IG_SIG_KEY      = "b03e0daaf2ab17cda2a569cace938d639d1288a1197f9ecf97efd0a4ec0874d7"
	GOINSTA_EXPERIMENTS     = "ig_android_sms_consent_in_reg,ig_android_flexible_sampling_universe,ig_android_background_conf_resend_fix,ig_restore_focus_on_reg_textbox_universe,ig_android_analytics_data_loss,ig_android_gmail_oauth_in_reg,ig_android_phoneid_sync_interval,ig_android_stay_at_one_tap_on_error,ig_android_link_to_access_if_email_taken_in_reg,ig_android_non_fb_sso,ig_android_family_apps_user_values_provider_universe,ig_android_reg_inline_errors,ig_android_run_fb_reauth_on_background,ig_fbns_push,ig_android_reg_omnibox,ig_android_show_password_in_reg_universe,ig_android_background_phone_confirmation_v2,ig_fbns_blocked,ig_android_access_redesign,ig_android_please_create_username_universe,ig_android_gmail_oauth_in_access,ig_android_reg_whiteout_redesign_v3"
	GOINSTA_SIG_KEY_VERSION = "4"
)

// GOINSTA_DEVICE_SETTINGS variable is a simulate of an android device
var GOINSTA_DEVICE_SETTINGS = map[string]interface{}{
	"manufacturer":    "Xiaomi",
	"model":           "HM 1SW",
	"android_version": 18,
	"android_release": "4.3",
}

// NewViaProxy All requests will use proxy server (example http://<ip>:<port>)
func NewViaProxy(username, password, proxy string) *Instagram {
	insta := New(username, password)
	insta.proxy = proxy
	return insta
}

// New try to fill Instagram struct
// New does not try to login , it will only fill
// Instagram struct
func New(username, password string) *Instagram {
	information := Informations{
		DeviceID: generateDeviceID(generateMD5Hash(username + password)),
		Username: username,
		Password: password,
		UUID:     generateUUID(true),
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
	insta.cookiejar, _ = cookiejar.New(nil) //newJar()

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
		"phone_id":            generateUUID(true),
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
	insta.GetV2Inbox()
	insta.GetRecentActivity()
	insta.GetReelsTrayFeed()

	return nil
}

// Logout of Instagram
func (insta *Instagram) Logout() error {
	_, err := insta.sendSimpleRequest("accounts/logout/")
	insta.cookiejar = nil
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
			"maxid":          maxID,
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
		},
	})
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)

	return resp, err
}

// MediaLikers return likers of a media , input is mediaid of a media
func (insta *Instagram) MediaLikers(mediaID string) (response.MediaLikersResponse, error) {
	body, err := insta.sendSimpleRequest("media/%s/likers/?", mediaID)
	if err != nil {
		return response.MediaLikersResponse{}, err
	}
	resp := response.MediaLikersResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// SyncFeatures simulates Instagram app behavior
func (insta *Instagram) SyncFeatures() error {
	data, err := insta.prepareData(map[string]interface{}{
		"id":          insta.LoggedInUser.ID,
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

// AutoCompleteUserList simulates Instagram app behavior
func (insta *Instagram) AutoCompleteUserList() error {
	_, err := insta.sendRequest(&reqOptions{
		Endpoint:     "friendships/autocomplete_user_list/",
		IgnoreStatus: true,
		Query: map[string]string{
			"version": "2",
		},
	})
	return err
}

// MegaphoneLog simulates Instagram app behavior
func (insta *Instagram) MegaphoneLog() error {
	data, err := insta.prepareData(map[string]interface{}{
		"id":        insta.LoggedInUser.ID,
		"type":      "feed_aysf",
		"action":    "seen",
		"reason":    "",
		"device_id": insta.Informations.DeviceID,
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
	result := response.StatusResponse{}
	data, err := insta.prepareData(map[string]interface{}{
		"id":         insta.LoggedInUser.ID,
		"experiment": "ig_android_profile_contextual_feed",
	})
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "qe/expose/",
		PostData: generateSignature(data),
	})
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &result)

	return err
}

// MediaInfo return media information
func (insta *Instagram) MediaInfo(mediaID string) (response.MediaInfoResponse, error) {
	result := response.MediaInfoResponse{}
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/info/", mediaID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// SetPublicAccount Sets account to public
func (insta *Instagram) SetPublicAccount() (response.ProfileDataResponse, error) {
	result := response.ProfileDataResponse{}
	data, err := insta.prepareData()
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "accounts/set_public/",
		PostData: generateSignature(data),
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// SetPrivateAccount Sets account to private
func (insta *Instagram) SetPrivateAccount() (response.ProfileDataResponse, error) {
	result := response.ProfileDataResponse{}
	data, err := insta.prepareData()
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "accounts/set_private/",
		PostData: generateSignature(data),
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// GetProfileData return current user information
func (insta *Instagram) GetProfileData() (response.ProfileDataResponse, error) {
	result := response.ProfileDataResponse{}
	data, err := insta.prepareData()
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "accounts/current_user/",
		PostData: generateSignature(data),
		Query: map[string]string{
			"edit": "true",
		},
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// RemoveProfilePicture will remove current logged in user profile picture
func (insta *Instagram) RemoveProfilePicture() (response.ProfileDataResponse, error) {
	result := response.ProfileDataResponse{}
	data, err := insta.prepareData()
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "accounts/remove_profile_picture/",
		PostData: generateSignature(data),
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// GetuserID return information of a user by user ID
func (insta *Instagram) GetUserByID(userID int64) (response.GetUsernameResponse, error) {
	result := response.GetUsernameResponse{}
	data, err := insta.prepareData()
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("users/%d/info/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// GetUsername return information of a user by username
func (insta *Instagram) GetUserByUsername(username string) (response.GetUsernameResponse, error) {
	body, err := insta.sendSimpleRequest("users/%s/usernameinfo/", username)
	if err != nil {
		return response.GetUsernameResponse{}, err
	}

	resp := response.GetUsernameResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// SearchLocation return search location by lat & lng & search query in instagram
func (insta *Instagram) SearchLocation(lat, lng, search string) (response.SearchLocationResponse, error) {
	if lat == "" || lng == "" {
		return response.SearchLocationResponse{}, fmt.Errorf("lat & lng must not be empty")
	}

	query := map[string]string{
		"rank_token":     insta.Informations.RankToken,
		"latitude":       lat,
		"longitude":      lng,
		"ranked_content": "true",
	}

	if search != "" {
		query["search_query"] = search
	} else {
		query["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	}
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "location_search/",
		Query:    query,
	})

	if err != nil {
		return response.SearchLocationResponse{}, err
	}

	resp := response.SearchLocationResponse{}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// GetLocationFeed return location feed data by locationID in Instagram
func (insta *Instagram) GetLocationFeed(locationID int64, maxID string) (response.LocationFeedResponse, error) {
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("feed/location/%d/", locationID),
		Query: map[string]string{
			"max_id": maxID,
		},
	})
	if err != nil {
		return response.LocationFeedResponse{}, err
	}

	resp := response.LocationFeedResponse{}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// GetTagRelated can get related tags by tags in instagram
func (insta *Instagram) GetTagRelated(tag string) (response.TagRelatedResponse, error) {
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("tags/%s/related", tag),
		Query: map[string]string{
			"visited":       fmt.Sprintf(`[{"id":"%s","type":"hashtag"}]`, tag),
			"related_types": `["hashtag"]`,
		},
	})

	if err != nil {
		return response.TagRelatedResponse{}, err
	}
	resp := response.TagRelatedResponse{}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// TagFeed search by tags in instagram
func (insta *Instagram) TagFeed(tag string) (response.TagFeedsResponse, error) {
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("feed/tag/%s/", tag),
		Query: map[string]string{
			"rank_token":     insta.Informations.RankToken,
			"ranked_content": "true",
		},
	})
	if err != nil {
		return response.TagFeedsResponse{}, err
	}

	resp := response.TagFeedsResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UploadPhoto can upload your photo with any quality , better to use 87
func (insta *Instagram) UploadPhoto(photo_path string, photo_caption string, upload_id int64, quality int, filter_type int) (response.UploadPhotoResponse, error) {
	photo_name := fmt.Sprintf("pending_media_%d.jpg", upload_id)

	//multipart request body
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	f, err := os.Open(photo_path)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	defer f.Close()

	w.WriteField("upload_id", strconv.FormatInt(upload_id, 10))
	w.WriteField("_uuid", insta.Informations.UUID)
	w.WriteField("_csrftoken", insta.Informations.Token)
	w.WriteField("image_compression", `{"lib_name":"jt","lib_version":"1.3.0","quality":"`+strconv.Itoa(quality)+`"}`)

	fw, err := w.CreateFormFile("photo", photo_name)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return response.UploadPhotoResponse{}, err
	}
	if err := w.Close(); err != nil {
		return response.UploadPhotoResponse{}, err
	}

	//making post request
	req, err := http.NewRequest("POST", GOINSTA_API_URL+"upload/photo/", &b)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	req.Header.Set("X-IG-Capabilities", "3Q4=")
	req.Header.Set("X-IG-Connection-Type", "WIFI") // cool header :smile:
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-type", w.FormDataContentType())
	req.Header.Set("Connection", "close")
	req.Header.Set("User-Agent", GOINSTA_USER_AGENT)

	client := &http.Client{
		Jar: insta.cookiejar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}

	if resp.StatusCode != 200 {
		return response.UploadPhotoResponse{}, fmt.Errorf("invalid status code" + resp.Status)
	}

	upresponse := response.UploadResponse{}
	err = json.Unmarshal(body, &upresponse)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}

	if upresponse.Status == "ok" {
		w, h, err := getImageDimension(photo_path)
		if err != nil {
			return response.UploadPhotoResponse{}, err
		}

		config := map[string]interface{}{
			"media_folder": "Instagram",
			"source_type":  4,
			"caption":      photo_caption,
			"upload_id":    strconv.FormatInt(upload_id, 10),
			"device":       GOINSTA_DEVICE_SETTINGS,
			"edits": map[string]interface{}{
				"crop_original_size": []int{w * 1.0, h * 1.0},
				"crop_center":        []float32{0.0, 0.0},
				"crop_zoom":          1.0,
				"filter_type":        filter_type,
			},
			"extra": map[string]interface{}{
				"source_width":  w,
				"source_height": h,
			},
		}
		data, err := insta.prepareData(config)
		if err != nil {
			return response.UploadPhotoResponse{}, err
		}

		body, err = insta.sendRequest(&reqOptions{
			Endpoint: "media/configure/?",
			PostData: generateSignature(data),
		})
		if err != nil {
			return response.UploadPhotoResponse{}, err
		}

		uploadresponse := response.UploadPhotoResponse{}
		err = json.Unmarshal(body, &uploadresponse)

		return uploadresponse, err
	} else {
		return response.UploadPhotoResponse{}, fmt.Errorf(upresponse.Status)
	}
}

// NewUploadID return unix nano time
func (insta *Instagram) NewUploadID() int64 {
	return time.Now().UnixNano()
}

// Follow one of instagram users with userID , you can find userID in GetUsername
func (insta *Instagram) Follow(userID int64) (response.FollowResponse, error) {
	resp := response.FollowResponse{}
	data, err := insta.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return resp, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/create/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UnFollow one of instagram users with userID , you can find userID in GetUsername
func (insta *Instagram) UnFollow(userID int64) (response.UnFollowResponse, error) {
	resp := response.UnFollowResponse{}
	data, err := insta.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return resp, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/destroy/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)

	return resp, err
}

func (insta *Instagram) Block(userID int64) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/block/%d/", userID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) UnBlock(userID int64) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/unblock/%d/", userID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) Like(mediaID string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/like/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) UnLike(mediaID string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/unlike/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) EditMedia(mediaID string, caption string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"caption_text": caption,
	})
	if err != nil {
		return []byte{}, err
	}

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

func (insta *Instagram) GetV2Inbox() (response.DirectListResponse, error) {
	result := response.DirectListResponse{}
	body, err := insta.sendSimpleRequest("direct_v2/inbox/?")
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
}

func (insta *Instagram) ChangePassword(newpassword string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"old_password":  insta.Informations.Password,
		"new_password1": newpassword,
		"new_password2": newpassword,
	})
	if err != nil {
		return []byte{}, err
	}
	bytes, err := insta.sendRequest(&reqOptions{
		Endpoint: "accounts/change_password/",
		PostData: generateSignature(data),
	})
	if err == nil {
		insta.Informations.Password = newpassword
	}
	return bytes, err
}

func (insta *Instagram) Timeline(maxID string) ([]byte, error) {
	return insta.sendRequest(&reqOptions{
		Endpoint: "feed/timeline/",
		Query: map[string]string{
			"max_id":         maxID,
			"rank_token":     insta.Informations.RankToken,
			"ranked_content": "true",
		},
	})
}

// getImageDimension return image dimension , types is .jpg and .png
func getImageDimension(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}

func (insta *Instagram) SelfUserFollowers(maxID string) (response.UsersResponse, error) {
	return insta.UserFollowers(insta.LoggedInUser.ID, maxID)
}

func (insta *Instagram) SelfUserFollowing(maxID string) (response.UsersResponse, error) {
	return insta.UserFollowing(insta.LoggedInUser.ID, maxID)
}

func (insta *Instagram) SelfTotalUserFollowers() (response.UsersResponse, error) {
	resp := response.UsersResponse{}
	for {
		temp_resp, err := insta.SelfUserFollowers(resp.NextMaxID)
		if err != nil {
			return response.UsersResponse{}, err
		}
		resp.Users = append(resp.Users, temp_resp.Users...)
		resp.PageSize += temp_resp.PageSize
		if !temp_resp.BigList {
			return resp, nil
		}
		resp.NextMaxID = temp_resp.NextMaxID
		resp.Status = temp_resp.Status
	}
}

func (insta *Instagram) SelfTotalUserFollowing() (response.UsersResponse, error) {
	resp := response.UsersResponse{}
	for {
		temp_resp, err := insta.SelfUserFollowing(resp.NextMaxID)
		if err != nil {
			return response.UsersResponse{}, err
		}
		resp.Users = append(resp.Users, temp_resp.Users...)
		resp.PageSize += temp_resp.PageSize
		if !temp_resp.BigList {
			return resp, nil
		}
		resp.NextMaxID = temp_resp.NextMaxID
		resp.Status = temp_resp.Status
	}
}

func (insta *Instagram) GetRecentActivity() ([]byte, error) {
	return insta.sendSimpleRequest("news/inbox/?")
}

func (insta *Instagram) GetFollowingRecentActivity() (response.FollowingRecentActivityResponse, error) {
	result := response.FollowingRecentActivityResponse{}
	bytes, err := insta.sendSimpleRequest("news/?")
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (insta *Instagram) SearchUsername(query string) (response.SearchUserResponse, error) {
	result := response.SearchUserResponse{}
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "users/search/",
		Query: map[string]string{
			"ig_sig_key_version": GOINSTA_SIG_KEY_VERSION,
			"is_typeahead":       "true",
			"query":              query,
			"rank_token":         insta.Informations.RankToken,
		},
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

func (insta *Instagram) SearchTags(query string) ([]byte, error) {
	return insta.sendRequest(&reqOptions{
		Endpoint: "tags/search/",
		Query: map[string]string{
			"is_typeahead": "true",
			"rank_token":   insta.Informations.RankToken,
		},
	})
}

func (insta *Instagram) SearchFacebookUsers(query string) ([]byte, error) {
	return insta.sendRequest(&reqOptions{
		Endpoint: "fbsearch/topsearch/",
		Query: map[string]string{
			"query":      query,
			"rank_token": insta.Informations.RankToken,
		},
	})
}

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
		Jar: insta.cookiejar,
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
func (insta *Instagram) GetUserStories(userID int64) (response.TrayUserResponse, error) {
	result := response.TrayUserResponse{}
	if userID == 0 {
		return result, nil
	}

	bytes, err := insta.sendSimpleRequest("feed/user/%d/reel_media/", userID)
	if err != nil {
		return result, err
	}

	json.Unmarshal([]byte(bytes), &result)

	return result, nil
}

func (insta *Instagram) UserFriendShip(userID int64) (response.UserFriendShipResponse, error) {
	result := response.UserFriendShipResponse{}
	data, err := insta.prepareData(map[string]interface{}{
		"user_id": userID,
	})

	if err != nil {
		return result, err
	}

	bytes, err := insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/show/%d/", userID),
		PostData: generateSignature(data),
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}
	return result, err
}

func (insta *Instagram) GetPopularFeed() (response.GetPopularFeedResponse, error) {
	result := response.GetPopularFeedResponse{}
	bytes, err := insta.sendRequest(&reqOptions{
		Endpoint: "feed/popular/",
		Query: map[string]string{
			"people_teaser_supported": "1",
			"rank_token":              insta.Informations.RankToken,
			"ranked_content":          "true",
		},
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}
	return result, err
}

func (insta *Instagram) prepareData(otherData ...map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.LoggedInUser.ID,
		"_csrftoken": insta.Informations.Token,
	}
	if len(otherData) > 0 {
		for i := range otherData {
			for key, value := range otherData[i] {
				data[key] = value
			}
		}
	}
	bytes, err := json.Marshal(data)
	return string(bytes), err
}
