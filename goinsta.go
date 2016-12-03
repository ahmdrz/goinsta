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
	"strings"
	"time"

	response "github.com/ahmdrz/goinsta/response"
)

// GetLastJson return latest json response from instagram
func (insta *Instagram) GetLastJson() string {
	return lastJson
}

// GetSessions return current instagram session and cookies
// Maybe need for webpages that use this API
func (insta *Instagram) GetSessions() map[string][]*http.Cookie {
	return cookiejar.cookies
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
	GOINSTA_USER_AGENT      = "Instagram 9.7.0 Android (18/4.3; 320dpi; 720x1280; Xiaomi; HM 1SW; armani; qcom; en_US)"
	GOINSTA_IG_SIG_KEY      = "2f6dcdf76deb0d3fd008886d032162a79b88052b5f50538c1ee93c4fe7d02e60"
	GOINSTA_EXPERIMENTS     = "ig_android_reg_login_btn_active_state,ig_android_ci_opt_in_at_reg,ig_android_one_click_in_old_flow,ig_android_merge_fb_and_ci_friends_page,ig_android_non_fb_sso,ig_android_mandatory_full_name,ig_android_reg_enable_login_password_btn,ig_android_reg_phone_email_active_state,ig_android_analytics_data_loss,ig_fbns_blocked,ig_android_contact_point_triage,ig_android_reg_next_btn_active_state,ig_android_prefill_phone_number,ig_android_show_fb_social_context_in_nux,ig_android_one_tap_login_upsell,ig_fbns_push,ig_android_phoneid_sync_interval"
	GOINSTA_SIG_KEY_VERSION = "4"
)

// GOINSTA_DEVICE_SETTINGS variable is a simulate of an android device
var GOINSTA_DEVICE_SETTINGS = map[string]interface{}{
	"manufacturer":    "Xiaomi",
	"model":           "HM 1SW",
	"android_version": 18,
	"android_release": "4.3",
}

// New try to fill Instagram struct
// New does not try to login , it will only fill
// Instagram struct
func New(username, password string) *Instagram {
	cookiejar = newJar()
	return &Instagram{
		Informations: Informations{
			DeviceID: generateDeviceID(generateMD5Hash(username + password)),
			Username: username,
			Password: password,
			UUID:     generateUUID(true),
		},
	}
}

// Login to Instagram.
// return error if can't send request to instagram server
func (insta *Instagram) Login() error {
	err := insta.sendRequest("si/fetch_headers/?challenge_type=signup&guid="+generateUUID(false), "", true)
	if err != nil {
		return fmt.Errorf("Login failed for %s error %s :", insta.Informations.Username, err.Error())
	}

	data := cookie[strings.Index(cookie, "csrftoken=")+10:]
	data = data[:strings.Index(data, ";")]

	result, err := json.Marshal(map[string]interface{}{
		"guid":                insta.Informations.UUID,
		"login_attempt_count": 0,
		"_csrftoken":          data,
		"device_id":           insta.Informations.DeviceID,
		"phone_id":            generateUUID(true),
		"username":            insta.Informations.Username,
		"password":            insta.Informations.Password,
	})
	if err != nil {
		return err
	}

	err = insta.sendRequest("accounts/login/", generateSignature(string(result)), true)
	if err != nil {
		return err
	}

	var Result struct {
		LoggedInUser response.User `json:"logged_in_user"`
		Status       string        `json:"status"`
	}

	err = json.Unmarshal([]byte(lastJson), &Result)
	if err != nil {
		return err
	}

	insta.Informations.Token = data
	insta.Informations.UsernameId = strconv.FormatInt(Result.LoggedInUser.PK, 10)
	insta.Informations.RankToken = insta.Informations.UsernameId + "_" + insta.Informations.UUID
	insta.IsLoggedIn = true
	insta.LoggedInUser = Result.LoggedInUser

	return nil
}

// Logout of Instagram
func (insta *Instagram) Logout() error {
	err := insta.sendRequest("accounts/logout/", "", false)
	cookiejar = nil
	return err
}

// UserFollowings return followings of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowing(userid, maxid string) (response.UsersReponse, error) {
	err := insta.sendRequest("friendships/"+userid+"/following/?max_id="+maxid+"&ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.UsersReponse{}, err
	}

	resp := response.UsersReponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.UsersReponse{}, err
	}

	return resp, nil
}

// UserFollowers return followers of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowers(userid, maxid string) (response.UsersReponse, error) {
	err := insta.sendRequest("friendships/"+userid+"/followers/?max_id="+maxid+"&ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.UsersReponse{}, err
	}

	resp := response.UsersReponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.UsersReponse{}, err
	}

	return resp, nil
}

// FirstUserFeed latest users feed
func (insta *Instagram) FirstUserFeed(userid string) (response.UserFeedResponse, error) {
	err := insta.sendRequest("feed/user/"+userid+"/?rank_token="+insta.Informations.RankToken+"&maxid=&min_timestamp=&ranked_content=true", "", false)
	if err != nil {
		return response.UserFeedResponse{}, err
	}
	resp := response.UserFeedResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.UserFeedResponse{}, err
	}

	return resp, nil
}

// UserFeed has tree mode ,
// If input was one string that we call maxid , mode is pagination
// If input was two string can pagination by timestamp and maxid
// If input was empty default value will select.
func (insta *Instagram) UserFeed(strings ...string) (response.FeedsResponse, error) {

	if len(strings) == 2 { // maxid and timestamp
		err := insta.sendRequest("feed/user/"+insta.Informations.UsernameId+"/?rank_token="+insta.Informations.RankToken+"&maxid="+strings[0]+"&min_timestamp="+strings[1]+"&ranked_content=true", "", false)
		if err != nil {
			return response.FeedsResponse{}, err
		}
	} else if len(strings) == 1 { // only maxid
		err := insta.sendRequest("feed/user/"+insta.Informations.UsernameId+"/?rank_token="+insta.Informations.RankToken+"&maxid="+strings[0]+"&ranked_content=true", "", false)
		if err != nil {
			return response.FeedsResponse{}, err
		}
	} else if len(strings) == 0 { // nothing (current user)
		err := insta.sendRequest("feed/user/"+insta.Informations.UsernameId+"/?rank_token="+insta.Informations.RankToken+"&ranked_content=true", "", false)
		if err != nil {
			return response.FeedsResponse{}, err
		}
	} else {
		return response.FeedsResponse{}, fmt.Errorf("invalid input")

	}

	resp := response.FeedsResponse{}
	err := json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.FeedsResponse{}, err
	}

	return resp, nil
}

// MediaLikers return likers of a media , input is mediaid of a media
func (insta *Instagram) MediaLikers(mediaId string) (response.MediaLikersResponse, error) {
	err := insta.sendRequest("media/"+mediaId+"/likers/?", "", false)
	if err != nil {
		return response.MediaLikersResponse{}, err
	}
	resp := response.MediaLikersResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.MediaLikersResponse{}, err
	}

	return resp, nil
}

// Expose , expose instagram
// return error if status was not 'ok' or runtime error
func (insta *Instagram) Expose() error {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"id":         insta.Informations.UsernameId,
		"experiment": "ig_android_profile_contextual_feed",
	})
	if err != nil {
		return err
	}

	err = insta.sendRequest("qe/expose/", generateSignature(string(bytes)), false)
	if err != nil {
		return err
	}

	resp := response.StatusResponse{}

	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return err
	}

	return nil
}

// MediaInfo return media information
func (insta *Instagram) MediaInfo(mediaId string) (response.FeedsResponse, error) {

	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"media_id":   mediaId,
	})
	if err != nil {
		return response.FeedsResponse{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/info/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.FeedsResponse{}, err
	}

	resp := response.FeedsResponse{}

	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.FeedsResponse{}, err
	}

	return resp, nil
}

// SetPublicAccount Sets account to public
func (insta *Instagram) SetPublicAccount() (response.ProfileDataResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	err = insta.sendRequest("accounts/set_public/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	return resp, nil
}

// SetPrivateAccount Sets account to private
func (insta *Instagram) SetPrivateAccount() (response.ProfileDataResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	err = insta.sendRequest("accounts/set_private/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	return resp, nil
}

// GetProfileData return current user information
func (insta *Instagram) GetProfileData() (response.ProfileDataResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	err = insta.sendRequest("accounts/current_user/?edit=true", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	return resp, nil
}

// RemoveProfilePicture will remove current logged in user profile picture
func (insta *Instagram) RemoveProfilePicture() (response.ProfileDataResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	err = insta.sendRequest("accounts/remove_profile_picture/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	return resp, nil
}

// GetUsername return information of a user by username
func (insta *Instagram) GetUsername(username string) (response.GetUsernameResponse, error) {
	err := insta.sendRequest("users/"+username+"/usernameinfo/", "", false)
	if err != nil {
		return response.GetUsernameResponse{}, err
	}

	resp := response.GetUsernameResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.GetUsernameResponse{}, err
	}

	return resp, nil
}

// TagFeed search by tags in instagram
func (insta *Instagram) TagFeed(tag string) (response.TagFeedsResponse, error) {
	err := insta.sendRequest("feed/tag/"+tag+"/?rank_token="+insta.Informations.RankToken+"&ranked_content=true", "", false)
	if err != nil {
		return response.TagFeedsResponse{}, err
	}

	resp := response.TagFeedsResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.TagFeedsResponse{}, err
	}

	return resp, nil
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
	w.WriteField("image_compression", "{\"lib_name\":\"jt\",\"lib_version\":\"1.3.0\",\"quality\":\""+strconv.Itoa(quality)+"\"}")

	fw, err := w.CreateFormFile("photo", photo_name)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return response.UploadPhotoResponse{}, err
	}
	w.Close()

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

	tempjar := newJar()
	for key, value := range cookiejar.cookies { // make a copy of session
		tempjar.cookies[key] = value
	}

	client := &http.Client{
		Jar: tempjar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	defer resp.Body.Close()

	lastResponse = resp
	cookie = resp.Header.Get("Set-Cookie")

	body, _ := ioutil.ReadAll(resp.Body)

	lastJson = string(body)

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

		var config map[string]interface{} = make(map[string]interface{})
		config["_csrftoken"] = insta.Informations.Token
		config["media_folder"] = "Instagram"
		config["source_type"] = 4
		config["_uid"] = insta.Informations.UsernameId
		config["_uuid"] = insta.Informations.UUID
		config["caption"] = photo_caption
		config["upload_id"] = strconv.FormatInt(upload_id, 10)
		config["device"] = GOINSTA_DEVICE_SETTINGS
		config["edits"] = map[string]interface{}{
			"crop_original_size": []int{w * 1.0, h * 1.0},
			"crop_center":        []float32{0.0, 0.0},
			"crop_zoom":          1.0,
			"filter_type":        filter_type,
		}
		config["extra"] = map[string]interface{}{
			"source_width":  w,
			"source_height": h,
		}

		bytes, err := json.Marshal(config)
		if err != nil {
			return response.UploadPhotoResponse{}, err
		}
		err = insta.sendRequest("media/configure/?", generateSignature(string(bytes)), false)
		if err != nil {
			return response.UploadPhotoResponse{}, err
		}

		uploadresponse := response.UploadPhotoResponse{}
		err = json.Unmarshal([]byte(lastJson), &uploadresponse)
		if err != nil {
			return response.UploadPhotoResponse{}, err
		}

		return uploadresponse, nil
	} else {
		return response.UploadPhotoResponse{}, fmt.Errorf(upresponse.Status)
	}
}

// NewUploadID return unix nano time
func (insta *Instagram) NewUploadID() int64 {
	return time.Now().UnixNano()
}

// Follow one of instagram users with userid , you can find userid in GetUsername
func (insta *Instagram) Follow(userid string) (response.FollowResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})
	if err != nil {
		return response.FollowResponse{}, err
	}

	err = insta.sendRequest("friendships/create/"+userid+"/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.FollowResponse{}, err
	}

	resp := response.FollowResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.FollowResponse{}, err
	}

	return resp, nil
}

// UnFollow one of instagram users with userid , you can find userid in GetUsername
func (insta *Instagram) UnFollow(userid string) (response.UnFollowResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})
	if err != nil {
		return response.UnFollowResponse{}, err
	}

	err = insta.sendRequest("friendships/destroy/"+userid+"/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.UnFollowResponse{}, err
	}

	resp := response.UnFollowResponse{}
	err = json.Unmarshal([]byte(lastJson), &resp)
	if err != nil {
		return response.UnFollowResponse{}, err
	}

	return resp, nil
}

func (insta *Instagram) Block(userid string) ([]byte, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("friendships/block/"+userid+"/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) UnBlock(userid string) ([]byte, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("friendships/unblock/"+userid+"/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) Like(mediaId string) ([]byte, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"media_id":   mediaId,
	})
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/like/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) UnLike(mediaId string) ([]byte, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"media_id":   mediaId,
	})
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/unlike", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) EditMedia(mediaId string, caption string) ([]byte, error) {
	var Data struct {
		UUID        string `json:"_uuid"`
		UID         string `json:"_uid"`
		CaptionText string `json:"caption_text"`
		CSRFToken   string `json:"_csrftoken"`
	}

	Data.UUID = insta.Informations.UUID
	Data.UID = insta.Informations.UsernameId
	Data.CaptionText = caption
	Data.CSRFToken = insta.Informations.Token

	bytes, err := json.Marshal(Data)
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/edit_media/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) DeleteMedia(mediaId string) ([]byte, error) {
	var Data struct {
		UUID      string `json:"_uuid"`
		UID       string `json:"_uid"`
		MediaID   string `json:"media_id"`
		CSRFToken string `json:"_csrftoken"`
	}

	Data.UUID = insta.Informations.UUID
	Data.UID = insta.Informations.UsernameId
	Data.MediaID = mediaId
	Data.CSRFToken = insta.Informations.Token

	bytes, err := json.Marshal(Data)
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/delete/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) RemoveSelfTag(mediaId string) ([]byte, error) {
	var Data struct {
		UUID      string `json:"_uuid"`
		UID       string `json:"_uid"`
		CSRFToken string `json:"_csrftoken"`
	}

	Data.UUID = insta.Informations.UUID
	Data.UID = insta.Informations.UsernameId
	Data.CSRFToken = insta.Informations.Token

	bytes, err := json.Marshal(Data)
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/remove/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) Comment(mediaId, text string) ([]byte, error) {
	var Data struct {
		UUID        string `json:"_uuid"`
		UID         string `json:"_uid"`
		CSRFToken   string `json:"_csrftoken"`
		CommentText string `json:"comment_text"`
	}

	Data.UUID = insta.Informations.UUID
	Data.UID = insta.Informations.UsernameId
	Data.CSRFToken = insta.Informations.Token
	Data.CommentText = text

	bytes, err := json.Marshal(Data)
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/comment/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) DeleteComment(mediaId, commentId string) ([]byte, error) {
	var Data struct {
		UUID      string `json:"_uuid"`
		UID       string `json:"_uid"`
		CSRFToken string `json:"_csrftoken"`
	}

	Data.UUID = insta.Informations.UUID
	Data.UID = insta.Informations.UsernameId
	Data.CSRFToken = insta.Informations.Token

	bytes, err := json.Marshal(Data)
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("media/"+mediaId+"/comment/"+commentId+"/delete/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) GetRecentRecipients() ([]byte, error) {
	err := insta.sendRequest("direct_share/recent_recipients/", "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) GetV2Inbox() ([]byte, error) {
	err := insta.sendRequest("direct_v2/inbox/?", "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) GetDirectPendingRequests() (response.DirectPendingRequests, error) {
	err := insta.sendRequest("direct_v2/pending_inbox/?", "", false)
	if err != nil {
		return response.DirectPendingRequests{}, err
	}

	result := response.DirectPendingRequests{}
	json.Unmarshal([]byte(insta.GetLastJson()), &result)
	return result, nil
}

func (insta *Instagram) GetRankedRecipients() (response.DirectRankedRecipients, error) {
	err := insta.sendRequest("direct_v2/ranked_recipients/?", "", false)
	if err != nil {
		return response.DirectRankedRecipients{}, err
	}

	result := response.DirectRankedRecipients{}
	json.Unmarshal([]byte(insta.GetLastJson()), &result)
	return result, nil
}

func (insta *Instagram) GetDirectThread(threadid string) (response.DirectThread, error) {
	err := insta.sendRequest("direct_v2/threads/"+threadid+"/", "", false)
	if err != nil {
		return response.DirectThread{}, err
	}

	result := response.DirectThread{}
	json.Unmarshal([]byte(insta.GetLastJson()), &result)
	return result, nil
}

func (insta *Instagram) Explore() ([]byte, error) {
	err := insta.sendRequest("discover/explore/", "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) ChangePassword(newpassword string) ([]byte, error) {
	var Data struct {
		UUID         string `json:"_uuid"`
		UID          string `json:"_uid"`
		CSRFToken    string `json:"_csrftoken"`
		OldPassword  string `json:"old_password"`
		NewPassword1 string `json:"new_password1"`
		NewPassword2 string `json:"new_password2"`
	}

	Data.UUID = insta.Informations.UUID
	Data.UID = insta.Informations.UsernameId
	Data.CSRFToken = insta.Informations.Token
	Data.OldPassword = insta.Informations.Password
	Data.NewPassword1 = newpassword
	Data.NewPassword2 = newpassword

	bytes, err := json.Marshal(Data)
	if err != nil {
		return []byte{}, err
	}

	err = insta.sendRequest("accounts/change_password/", generateSignature(string(bytes)), false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) Timeline(maxid ...string) ([]byte, error) {
	nextmaxid := ""

	if len(maxid) == 0 {
		nextmaxid = ""
	} else if len(maxid) == 1 {
		nextmaxid = "&max_id=" + maxid[0]
	} else {
		return []byte{}, fmt.Errorf("Incorrect input")
	}

	err := insta.sendRequest("feed/timeline/?rank_token="+insta.Informations.RankToken+"&ranked_content=true"+nextmaxid, "", false)

	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil

}

// getImageDimension return image dimension , types is .jpg and .png
func getImageDimension(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, err
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}

func (insta *Instagram) SelfUserFollowers(maxid string) (response.UsersReponse, error) {
	return insta.UserFollowers(insta.Informations.UsernameId, maxid)
}

func (insta *Instagram) SelfUserFollowing(maxid string) (response.UsersReponse, error) {
	return insta.UserFollowing(insta.Informations.UsernameId, maxid)
}

func (insta *Instagram) SelfTotalUserFollowers() (response.UsersReponse, error) {
	resp := response.UsersReponse{}
	for {
		temp_resp, err := insta.SelfUserFollowers(resp.NextMaxID)
		if err != nil {
			return response.UsersReponse{}, err
		}
		for _, user := range temp_resp.Users {
			resp.Users = append(resp.Users, user)
		}
		resp.PageSize += temp_resp.PageSize
		if !temp_resp.BigList {
			return resp, nil
		}
		resp.NextMaxID = temp_resp.NextMaxID
		resp.Status = temp_resp.Status
	}
}

func (insta *Instagram) SelfTotalUserFollowing() (response.UsersReponse, error) {
	resp := response.UsersReponse{}
	for {
		temp_resp, err := insta.SelfUserFollowing(resp.NextMaxID)
		if err != nil {
			return response.UsersReponse{}, err
		}
		for _, user := range temp_resp.Users {
			resp.Users = append(resp.Users, user)
		}
		resp.PageSize += temp_resp.PageSize
		if !temp_resp.BigList {
			return resp, nil
		}
		resp.NextMaxID = temp_resp.NextMaxID
		resp.Status = temp_resp.Status
	}
}

func (insta *Instagram) GetRecentActivity() ([]byte, error) {
	err := insta.sendRequest("news/inbox/?", "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) GetFollowingRecentActivity() ([]byte, error) {
	err := insta.sendRequest("news/?", "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) SearchUsername(query string) (response.SearchUserResponse, error) {
	err := insta.sendRequest("users/search/?ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&is_typeahead=true&query="+url.QueryEscape(query)+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.SearchUserResponse{}, err
	}

	result := response.SearchUserResponse{}
	json.Unmarshal([]byte(lastJson), &result)

	return result, nil
}

func (insta *Instagram) SearchTags(query string) ([]byte, error) {
	err := insta.sendRequest("tags/search/?is_typeahead=true&q="+query+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) SearchFacebookUsers(query string) ([]byte, error) {
	err := insta.sendRequest("fbsearch/topsearch/?context=blended&query="+query+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return []byte{}, err
	}

	return []byte(lastJson), nil
}

func (insta *Instagram) DirectMessage(recipient string, message string) (response.DirectMessageResponse, error) {
	recipients, err := json.Marshal([][]string{[]string{recipient}})
	if err != nil {
		return response.DirectMessageResponse{}, err
	}
	threads, err := json.Marshal([]string{"0"})
	if err != nil {
		return response.DirectMessageResponse{}, err
	}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.SetBoundary(insta.Informations.UUID)
	w.WriteField("recipient_users", string(recipients))
	w.WriteField("client_context", insta.Informations.UUID)
	w.WriteField("thread_ids", string(threads))
	w.WriteField("text", message)
	w.Close()

	req, err := http.NewRequest("POST", GOINSTA_API_URL+"direct_v2/threads/broadcast/text/", &b)
	if err != nil {
		return response.DirectMessageResponse{}, err
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-en")
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", GOINSTA_USER_AGENT)

	tempjar := newJar()
	for key, value := range cookiejar.cookies { // make a copy of session
		tempjar.cookies[key] = value
	}

	client := &http.Client{
		Jar: tempjar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return response.DirectMessageResponse{}, err
	}
	defer resp.Body.Close()

	lastResponse = resp
	cookie = resp.Header.Get("Set-Cookie")

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return response.DirectMessageResponse{}, fmt.Errorf(string(body))
	}

	lastJson = string(body)

	result := response.DirectMessageResponse{}
	json.Unmarshal(body, &result)
	return result, nil
}

func (insta *Instagram) GetReelsTrayFeed() {
	insta.sendRequest("feed/reels_tray/", "", false)
}
