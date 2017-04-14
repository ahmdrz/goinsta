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
	GOINSTA_USER_AGENT      = "Instagram 10.1.0 Android (18/4.3; 320dpi; 720x1280; Xiaomi; HM 1SW; armani; qcom; en_US)"
	GOINSTA_IG_SIG_KEY      = "f6bcd0639ea3d279c10d4ffe49bd59f1f5045cf73c62fc33c9e45766f6cf6c63"
	GOINSTA_EXPERIMENTS     = "ig_android_ad_holdout_16m5_universe,ig_android_progressive_jpeg,ig_creation_growth_holdout,ig_android_oppo_app_badging,ig_android_ad_remove_username_from_caption_universe,ig_android_enable_share_to_whatsapp,ig_android_direct_drawing_in_quick_cam_universe,ig_android_ad_always_send_ad_attribution_id_universe,ig_android_universe_video_production,ig_android_direct_plus_button,ig_android_ads_heatmap_overlay_universe,ig_android_http_stack_experiment_2016,ig_android_infinite_scrolling,ig_fbns_blocked,ig_android_post_auto_retry_v7_21,ig_fbns_push,ig_android_video_playback_bandwidth_threshold,ig_android_direct_link_preview,ig_android_direct_typing_indicator,ig_android_preview_capture,ig_android_feed_pill,ig_android_profile_link_iab,ig_android_story_caption,ig_android_network_cancellation,ig_android_histogram_reporter,ig_android_anrwatchdog,ig_android_search_client_matching,ig_android_follow_request_text_buttons,ig_android_feed_zoom,ig_android_drafts_universe,ig_android_disable_comment,ig_android_user_detail_endpoint,ig_android_os_version_blocking,ig_android_blocked_list,ig_android_event_creation,ig_android_high_res_upload_2,ig_android_2fac,ig_android_mark_reel_seen_on_Swipe_forward,ig_android_comment_redesign,ig_android_ad_sponsored_label_universe,ig_android_mentions_dismiss_rule,ig_android_disable_chroma_subsampling,ig_android_share_spinner,ig_android_video_reuse_surface,ig_explore_v3_android_universe,ig_android_media_favorites,ig_android_nux_holdout,ig_android_insta_video_universe,ig_android_search_null_state,ig_android_universe_reel_video_production,liger_instagram_android_univ,ig_android_direct_emoji_picker,ig_feed_holdout_universe,ig_android_direct_send_auto_retry_universe,ig_android_samsung_app_badging,ig_android_disk_usage,ig_android_business_promotion,ig_android_direct_swipe_to_inbox,ig_android_feed_reshare_button_nux,ig_android_react_native_boost_post,ig_android_boomerang_feed_attribution,ig_fbns_shared,ig_fbns_dump_ids,ig_android_react_native_universe,ig_show_promote_button_in_feed,ig_android_ad_metadata_behavior_universe,ig_android_video_loopcount_int,ig_android_inline_gallery_backoff_hours_universe,ig_android_rendering_controls,ig_android_profile_photo_as_media,ig_android_async_stack_image_cache,ig_video_max_duration_qe_preuniverse,ig_video_copyright_whitelist,ig_android_render_stories_with_content_override,ig_android_ad_intent_to_highlight_universe,ig_android_swipe_navigation_x_angle_universe,ig_android_disable_comment_public_test,ig_android_profile,ig_android_direct_blue_tab,ig_android_enable_share_to_messenger,ig_android_fetch_reel_tray_on_resume_universe,ig_android_promote_again,ig_feed_event_landing_page_channel,ig_ranking_following,ig_android_pending_request_search_bar,ig_android_feed_ufi_redesign,ig_android_pending_edits_dialog_universe,ig_android_business_conversion_flow_universe,ig_android_show_your_story_when_empty_universe,ig_android_ad_drop_cookie_early,ig_android_app_start_config,ig_android_fix_ise_two_phase,ig_android_ppage_toggle_universe,ig_android_pbia_normal_weight_universe,ig_android_profanity_filter,ig_ios_su_activity_feed,ig_android_search,ig_android_boomerang_entry,ig_android_mute_story,ig_android_inline_gallery_universe,ig_android_ad_remove_one_tap_indicator_universe,ig_android_view_count_decouple_likes_universe,ig_android_contact_button_redesign_v2,ig_android_periodic_analytics_upload_v2,ig_android_send_direct_typing_indicator,ig_android_ad_holdout_16h2m1_universe,ig_android_react_native_comment_moderation_settings,ig_video_use_sve_universe,ig_android_inline_gallery_no_backoff_on_launch_universe,ig_android_immersive_viewer,ig_android_discover_people_icon,ig_android_profile_follow_back_button,is_android_feed_seen_state,ig_android_dense_feed_unit_cards,ig_android_drafts_video_universe,ig_android_exoplayer,ig_android_add_to_last_post,ig_android_ad_remove_cta_chevron_universe,ig_android_ad_comment_cta_universe,ig_android_search_event_icon,ig_android_channels_home,ig_android_feed,ig_android_dv2_realtime_private_share,ig_android_non_square_first,ig_android_video_interleaved_v2,ig_android_video_cache_policy,ig_android_react_native_universe_kill_switch,ig_android_video_captions_universe,ig_android_follow_search_bar,ig_android_last_edits,ig_android_two_step_capture_flow,ig_android_video_download_logging,ig_android_share_link_to_whatsapp,ig_android_facebook_twitter_profile_photos,ig_android_swipeable_filters_blacklist,ig_android_ad_pbia_profile_tap_universe,ig_android_use_software_layer_for_kc_drawing_universe,ig_android_react_native_ota,ig_android_direct_mutually_exclusive_experiment_universe,ig_android_following_follower_social_context"
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
	proxyUrl = proxy
	return New(username, password)
}

// New try to fill Instagram struct
// New does not try to login , it will only fill
// Instagram struct
func New(username, password string) *Instagram {
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
	insta.cookiejar, _ = cookiejar.New(nil) //newJar()

	body, err := insta.sendRequest("si/fetch_headers/?challenge_type=signup&guid="+generateUUID(false), "", true)
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

	body, err = insta.sendRequest("accounts/login/", generateSignature(string(result)), true)
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
	insta.Timeline()
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
	_, err := insta.sendRequest("accounts/logout/", "", false)
	insta.cookiejar = nil
	return err
}

// UserFollowing return followings of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowing(userID int64, maxid string) (response.UsersResponse, error) {
	body, err := insta.sendRequest("friendships/"+strconv.FormatInt(userID, 10)+"/following/?max_id="+maxid+"&ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.UsersResponse{}, err
	}

	resp := response.UsersResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UserFollowers return followers of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowers(userID int64, maxid string) (response.UsersResponse, error) {
	body, err := insta.sendRequest("friendships/"+strconv.FormatInt(userID, 10)+"/followers/?max_id="+maxid+"&ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&rank_token="+insta.Informations.RankToken, "", false)
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

	body, err := insta.sendRequest("feed/user/"+strconv.FormatInt(userID, 10)+"/?rank_token="+insta.Informations.RankToken+"&maxid="+maxID+"&min_timestamp="+minTimestamp+"&ranked_content=true", "", false)
	if err != nil {
		return resp, err
	}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// MediaLikers return likers of a media , input is mediaid of a media
func (insta *Instagram) MediaLikers(mediaId string) (response.MediaLikersResponse, error) {
	body, err := insta.sendRequest("media/"+mediaId+"/likers/?", "", false)
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

	_, err = insta.sendRequest("qe/sync/", generateSignature(data), false)
	return err
}

// AutoCompleteUserList simulates Instagram app behavior
func (insta *Instagram) AutoCompleteUserList() error {
	_, err := insta.sendRequest("friendships/autocomplete_user_list/?version=2", "", false, false)
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
	_, err = insta.sendRequest("megaphone/log/", generateSignature(data), false)
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

	body, err := insta.sendRequest("qe/expose/", generateSignature(data), false)
	if err != nil {
		return err
	}

	fmt.Print(string(body))
	err = json.Unmarshal(body, &result)

	return err
}

// MediaInfo return media information
func (insta *Instagram) MediaInfo(mediaId string) (response.MediaInfoResponse, error) {
	result := response.MediaInfoResponse{}
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaId,
	})
	if err != nil {
		return result, err
	}

	body, err := insta.sendRequest("media/"+mediaId+"/info/", generateSignature(data), false)
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

	body, err := insta.sendRequest("accounts/set_public/", generateSignature(data), false)
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

	body, err := insta.sendRequest("accounts/set_private/", generateSignature(data), false)
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

	body, err := insta.sendRequest("accounts/current_user/?edit=true", generateSignature(data), false)
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

	body, err := insta.sendRequest("accounts/remove_profile_picture/", generateSignature(data), false)
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

	body, err := insta.sendRequest("users/"+strconv.FormatInt(userID, 10)+"/info/", generateSignature(data), false)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

// GetUsername return information of a user by username
func (insta *Instagram) GetUserByUsername(username string) (response.GetUsernameResponse, error) {
	body, err := insta.sendRequest("users/"+username+"/usernameinfo/", "", false)
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

	query := "?rank_token=" + insta.Informations.RankToken + "&latitude=" + lat + "&longitude=" + lng

	if search != "" {
		query += "&search_query=" + url.QueryEscape(search)
	} else {
		query += "&timestamp=" + string(time.Now().Unix())
	}
	query += "&ranked_content=true"

	body, err := insta.sendRequest("location_search/"+query, "", false)

	if err != nil {
		return response.SearchLocationResponse{}, err
	}

	resp := response.SearchLocationResponse{}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// GetLocationFeed return location feed data by locationID in Instagram
func (insta *Instagram) GetLocationFeed(locationID int64, maxID string) (response.LocationFeedResponse, error) {
	var query string
	var err error

	if maxID != "" {
		query += "?max_id=" + maxID
	}

	uri := fmt.Sprintf("feed/location/%d/", locationID) + query

	body, err := insta.sendRequest(uri, "", false)

	if err != nil {
		return response.LocationFeedResponse{}, err
	}

	resp := response.LocationFeedResponse{}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// GetTagRelated can get related tags by tags in instagram
func (insta *Instagram) GetTagRelated(tag string) (response.TagRelatedResponse, error) {
	visited := url.QueryEscape("[{\"id\":\"" + tag + "\",\"type\":\"hashtag\"}]")
	relatedTypes := url.QueryEscape("[\"hashtag\"]")
	body, err := insta.sendRequest("tags/"+tag+"/related?visited="+visited+"&related_types="+relatedTypes, "", false)

	if err != nil {
		return response.TagRelatedResponse{}, err
	}
	resp := response.TagRelatedResponse{}
	err = json.Unmarshal(body, &resp)
	return resp, err
}

// TagFeed search by tags in instagram
func (insta *Instagram) TagFeed(tag string) (response.TagFeedsResponse, error) {
	body, err := insta.sendRequest("feed/tag/"+tag+"/?rank_token="+insta.Informations.RankToken+"&ranked_content=true", "", false)
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

		body, err = insta.sendRequest("media/configure/?", generateSignature(data), false)
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

	body, err := insta.sendRequest("friendships/create/"+strconv.FormatInt(userID, 10)+"/", generateSignature(data), false)
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

	body, err := insta.sendRequest("friendships/destroy/"+strconv.FormatInt(userID, 10)+"/", generateSignature(data), false)
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

	return insta.sendRequest("friendships/block/"+strconv.FormatInt(userID, 10)+"/", generateSignature(data), false)
}

func (insta *Instagram) UnBlock(userID int64) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"user_id": userID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("friendships/unblock/"+strconv.FormatInt(userID, 10)+"/", generateSignature(data), false)
}

func (insta *Instagram) Like(mediaId string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaId,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/like/", generateSignature(data), false)
}

func (insta *Instagram) UnLike(mediaId string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaId,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/unlike/", generateSignature(data), false)
}

func (insta *Instagram) EditMedia(mediaId string, caption string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"caption_text": caption,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/edit_media/", generateSignature(data), false)
}

func (insta *Instagram) DeleteMedia(mediaId string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaId,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/delete/", generateSignature(data), false)
}

func (insta *Instagram) RemoveSelfTag(mediaId string) ([]byte, error) {
	data, err := insta.prepareData()
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/remove/", generateSignature(data), false)
}

func (insta *Instagram) Comment(mediaId, text string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"comment_text": text,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/comment/", generateSignature(data), false)
}

func (insta *Instagram) DeleteComment(mediaId, commentId string) ([]byte, error) {
	data, err := insta.prepareData()
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest("media/"+mediaId+"/comment/"+commentId+"/delete/", generateSignature(data), false)
}

func (insta *Instagram) GetRecentRecipients() ([]byte, error) {
	return insta.sendRequest("direct_share/recent_recipients/", "", false)
}

func (insta *Instagram) GetV2Inbox() (response.DirectListResponse, error) {
	result := response.DirectListResponse{}
	body, err := insta.sendRequest("direct_v2/inbox/?", "", false)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetDirectPendingRequests() (response.DirectPendingRequests, error) {
	result := response.DirectPendingRequests{}
	body, err := insta.sendRequest("direct_v2/pending_inbox/?", "", false)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetRankedRecipients() (response.DirectRankedRecipients, error) {
	result := response.DirectRankedRecipients{}
	body, err := insta.sendRequest("direct_v2/ranked_recipients/?", "", false)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetDirectThread(threadid string) (response.DirectThread, error) {
	result := response.DirectThread{}
	body, err := insta.sendRequest("direct_v2/threads/"+threadid+"/", "", false)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) Explore() (response.ExploreResponse, error) {
	result := response.ExploreResponse{}
	body, err := insta.sendRequest("discover/explore/", "", false)
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
	return insta.sendRequest("accounts/change_password/", generateSignature(data), false)
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

	return insta.sendRequest("feed/timeline/?rank_token="+insta.Informations.RankToken+"&ranked_content=true"+nextmaxid, "", false)
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

func (insta *Instagram) SelfUserFollowers(maxid string) (response.UsersResponse, error) {
	return insta.UserFollowers(insta.LoggedInUser.ID, maxid)
}

func (insta *Instagram) SelfUserFollowing(maxid string) (response.UsersResponse, error) {
	return insta.UserFollowing(insta.LoggedInUser.ID, maxid)
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
	return insta.sendRequest("news/inbox/?", "", false)
}

func (insta *Instagram) GetFollowingRecentActivity() (response.FollowingRecentActivityResponse, error) {
	result := response.FollowingRecentActivityResponse{}
	bytes, err := insta.sendRequest("news/?", "", false)
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
	body, err := insta.sendRequest("users/search/?ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&is_typeahead=true&query="+url.QueryEscape(query)+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
}

func (insta *Instagram) SearchTags(query string) ([]byte, error) {
	return insta.sendRequest("tags/search/?is_typeahead=true&q="+query+"&rank_token="+insta.Informations.RankToken, "", false)
}

func (insta *Instagram) SearchFacebookUsers(query string) ([]byte, error) {
	return insta.sendRequest("fbsearch/topsearch/?context=blended&query="+query+"&rank_token="+insta.Informations.RankToken, "", false)
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
	bytes, err := insta.sendRequest("feed/reels_tray/", "", false)
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

	bytes, err := insta.sendRequest("feed/user/"+strconv.FormatInt(userID, 10)+"/reel_media/", "", false)
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

	bytes, err := insta.sendRequest("friendships/show/"+strconv.FormatInt(userID, 10)+"/", generateSignature(data), false)
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
	bytes, err := insta.sendRequest("feed/popular/?people_teaser_supported=1&rank_token="+insta.Informations.RankToken+"&ranked_content=true&", "", false)
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
