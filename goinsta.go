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
	"net/http/cookiejar"
)

// GetSessions return current instagram session and cookies
// Maybe need for webpages that use this API
func (insta *Instagram) GetSessions(url *url.URL) []*http.Cookie {
	return insta.cookiejar.Cookies(url)
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

// All requests will use proxy server (example http://<ip>:<port>)
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
	insta.cookiejar, _ = cookiejar.New(nil)//newJar()

	body, err := insta.sendRequest("si/fetch_headers/?challenge_type=signup&guid="+generateUUID(false), "", true)
	if err != nil {
		return fmt.Errorf("Login failed for %s error %s :", insta.Informations.Username, err.Error())
	}

	data := insta.cookie[strings.Index(insta.cookie, "csrftoken=")+10:]
	data = data[:strings.Index(data, ";")]

	result, _ := json.Marshal(map[string]interface{}{
		"guid":                insta.Informations.UUID,
		"login_attempt_count": 0,
		"_csrftoken":          data,
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

	insta.Informations.Token = data
	insta.Informations.UsernameId = strconv.FormatInt(Result.LoggedInUser.PK, 10)
	insta.Informations.RankToken = insta.Informations.UsernameId + "_" + insta.Informations.UUID
	insta.IsLoggedIn = true
	insta.LoggedInUser = Result.LoggedInUser

	return nil
}

// Logout of Instagram
func (insta *Instagram) Logout() error {
	_, err := insta.sendRequest("accounts/logout/", "", false)
	insta.cookiejar = nil
	return err
}

// UserFollowings return followings of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowing(userid, maxid string) (response.UsersReponse, error) {
	body, err := insta.sendRequest("friendships/"+userid+"/following/?max_id="+maxid+"&ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.UsersReponse{}, err
	}

	resp := response.UsersReponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UserFollowers return followers of specific user
// skip maxid with empty string for get first page
func (insta *Instagram) UserFollowers(userid, maxid string) (response.UsersReponse, error) {
	body, err := insta.sendRequest("friendships/"+userid+"/followers/?max_id="+maxid+"&ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.UsersReponse{}, err
	}

	resp := response.UsersReponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// FirstUserFeed latest users feed
func (insta *Instagram) FirstUserFeed(userid string) (response.UserFeedResponse, error) {
	body, err := insta.sendRequest("feed/user/"+userid+"/?rank_token="+insta.Informations.RankToken+"&maxid=&min_timestamp=&ranked_content=true", "", false)
	if err != nil {
		return response.UserFeedResponse{}, err
	}
	resp := response.UserFeedResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// UserFeed has tree mode ,
// If input was one string that we call maxid , mode is pagination
// If input was two string can pagination by timestamp and maxid
// If input was empty default value will select.
func (insta *Instagram) UserFeed(strings ...string) (response.FeedsResponse, error) {
	var body []byte
	var err error

	if len(strings) == 2 { // maxid and timestamp
		body, err = insta.sendRequest("feed/user/"+insta.Informations.UsernameId+"/?rank_token="+insta.Informations.RankToken+"&maxid="+strings[0]+"&min_timestamp="+strings[1]+"&ranked_content=true", "", false)
		if err != nil {
			return response.FeedsResponse{}, err
		}
	} else if len(strings) == 1 { // only maxid
		body, err = insta.sendRequest("feed/user/"+insta.Informations.UsernameId+"/?rank_token="+insta.Informations.RankToken+"&maxid="+strings[0]+"&ranked_content=true", "", false)
		if err != nil {
			return response.FeedsResponse{}, err
		}
	} else if len(strings) == 0 { // nothing (current user)
		body, err = insta.sendRequest("feed/user/"+insta.Informations.UsernameId+"/?rank_token="+insta.Informations.RankToken+"&ranked_content=true", "", false)
		if err != nil {
			return response.FeedsResponse{}, err
		}
	} else {
		return response.FeedsResponse{}, fmt.Errorf("invalid input")

	}

	resp := response.FeedsResponse{}
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

// Expose , expose instagram
// return error if status was not 'ok' or runtime error
func (insta *Instagram) Expose() error {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"id":         insta.Informations.UsernameId,
		"experiment": "ig_android_profile_contextual_feed",
	})

	body, err := insta.sendRequest("qe/expose/", generateSignature(string(bytes)), false)
	if err != nil {
		return err
	}

	resp := response.StatusResponse{}

	err = json.Unmarshal(body, &resp)

	return err
}

// MediaInfo return media information
func (insta *Instagram) MediaInfo(mediaId string) (response.MediaInfoResponse, error) {

	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"media_id":   mediaId,
	})

	body, err := insta.sendRequest("media/"+mediaId+"/info/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.MediaInfoResponse{}, err
	}

	resp := response.MediaInfoResponse{}

	err = json.Unmarshal(body, &resp)

	return resp, err
}

// SetPublicAccount Sets account to public
func (insta *Instagram) SetPublicAccount() (response.ProfileDataResponse, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})

	body, err := insta.sendRequest("accounts/set_public/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// SetPrivateAccount Sets account to private
func (insta *Instagram) SetPrivateAccount() (response.ProfileDataResponse, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})

	body, err := insta.sendRequest("accounts/set_private/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// GetProfileData return current user information
func (insta *Instagram) GetProfileData() (response.ProfileDataResponse, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})

	body, err := insta.sendRequest("accounts/current_user/?edit=true", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// RemoveProfilePicture will remove current logged in user profile picture
func (insta *Instagram) RemoveProfilePicture() (response.ProfileDataResponse, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})

	body, err := insta.sendRequest("accounts/remove_profile_picture/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.ProfileDataResponse{}, err
	}

	resp := response.ProfileDataResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// GetUserID return information of a user by user ID
func (insta *Instagram) GetUserID(userid string) (response.GetUsernameResponse, error) {
	bytes, err := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
	})
	if err != nil {
		return response.GetUsernameResponse{}, err
	}

	body, err := insta.sendRequest("users/"+userid+"/info/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.GetUsernameResponse{}, err
	}

	resp := response.GetUsernameResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

// GetUsername return information of a user by username
func (insta *Instagram) GetUsername(username string) (response.GetUsernameResponse, error) {
	body, err := insta.sendRequest("users/"+username+"/usernameinfo/", "", false)
	if err != nil {
		return response.GetUsernameResponse{}, err
	}

	resp := response.GetUsernameResponse{}
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

	/*tempjar := newJar()
	for key, value := range insta.cookiejar.cookies { // make a copy of session
		tempjar.cookies[key] = value
	}*/

	client := &http.Client{
		Jar: insta.cookiejar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	defer resp.Body.Close()

	insta.cookie = resp.Header.Get("Set-Cookie")

	body, _ := ioutil.ReadAll(resp.Body)

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

		var config map[string]interface{} = map[string]interface{} {
			"_csrftoken": insta.Informations.Token,
			"media_folder": "Instagram",
			"source_type": 4,
			"_uid": insta.Informations.UsernameId,
			"_uuid": insta.Informations.UUID,
			"caption": photo_caption,
			"upload_id": strconv.FormatInt(upload_id, 10),
			"device": GOINSTA_DEVICE_SETTINGS,
			"edits": map[string]interface{} {
				"crop_original_size": []int{w * 1.0, h * 1.0},
				"crop_center":        []float32{0.0, 0.0},
				"crop_zoom":          1.0,
				"filter_type":        filter_type,
			},
			"extra": map[string]interface{} {
				"source_width":  w,
				"source_height": h,
			},
		}

		bytes, _ := json.Marshal(config)

		body, err = insta.sendRequest("media/configure/?", generateSignature(string(bytes)), false)
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

// Follow one of instagram users with userid , you can find userid in GetUsername
func (insta *Instagram) Follow(userid string) (response.FollowResponse, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})

	body, err := insta.sendRequest("friendships/create/"+userid+"/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.FollowResponse{}, err
	}

	resp := response.FollowResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
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

	body, err := insta.sendRequest("friendships/destroy/"+userid+"/", generateSignature(string(bytes)), false)
	if err != nil {
		return response.UnFollowResponse{}, err
	}

	resp := response.UnFollowResponse{}
	err = json.Unmarshal(body, &resp)

	return resp, err
}

func (insta *Instagram) Block(userid string) ([]byte, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})

	return insta.sendRequest("friendships/block/"+userid+"/", generateSignature(string(bytes)), false)
}

func (insta *Instagram) UnBlock(userid string) ([]byte, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"user_id":    userid,
	})

	return insta.sendRequest("friendships/unblock/"+userid+"/", generateSignature(string(bytes)), false)
}

func (insta *Instagram) Like(mediaId string) ([]byte, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"media_id":   mediaId,
	})

	return insta.sendRequest("media/"+mediaId+"/like/", generateSignature(string(bytes)), false)
}

func (insta *Instagram) UnLike(mediaId string) ([]byte, error) {
	bytes, _ := json.Marshal(map[string]interface{}{
		"_uuid":      insta.Informations.UUID,
		"_uid":       insta.Informations.UsernameId,
		"_csrftoken": insta.Informations.Token,
		"media_id":   mediaId,
	})

	return insta.sendRequest("media/"+mediaId+"/unlike", generateSignature(string(bytes)), false)
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

	bytes, _ := json.Marshal(Data)

	return insta.sendRequest("media/"+mediaId+"/edit_media/", generateSignature(string(bytes)), false)
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

	bytes, _ := json.Marshal(Data)

	return insta.sendRequest("media/"+mediaId+"/delete/", generateSignature(string(bytes)), false)
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

	bytes, _ := json.Marshal(Data)

	return insta.sendRequest("media/"+mediaId+"/remove/", generateSignature(string(bytes)), false)
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

	bytes, _ := json.Marshal(Data)

	return insta.sendRequest("media/"+mediaId+"/comment/", generateSignature(string(bytes)), false)
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

	bytes, _ := json.Marshal(Data)

	return insta.sendRequest("media/"+mediaId+"/comment/"+commentId+"/delete/", generateSignature(string(bytes)), false)
}

func (insta *Instagram) GetRecentRecipients() ([]byte, error) {
	return insta.sendRequest("direct_share/recent_recipients/", "", false)
}

func (insta *Instagram) GetV2Inbox() ([]byte, error) {
	return insta.sendRequest("direct_v2/inbox/?", "", false)
}

func (insta *Instagram) GetDirectPendingRequests() (response.DirectPendingRequests, error) {
	body, err := insta.sendRequest("direct_v2/pending_inbox/?", "", false)
	if err != nil {
		return response.DirectPendingRequests{}, err
	}

	result := response.DirectPendingRequests{}
	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetRankedRecipients() (response.DirectRankedRecipients, error) {
	body, err := insta.sendRequest("direct_v2/ranked_recipients/?", "", false)
	if err != nil {
		return response.DirectRankedRecipients{}, err
	}

	result := response.DirectRankedRecipients{}
	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) GetDirectThread(threadid string) (response.DirectThread, error) {
	body, err := insta.sendRequest("direct_v2/threads/"+threadid+"/", "", false)
	if err != nil {
		return response.DirectThread{}, err
	}

	result := response.DirectThread{}
	err = json.Unmarshal(body, &result)
	return result, err
}

func (insta *Instagram) Explore() (response.ExploreResponse, error) {
	body, err := insta.sendRequest("discover/explore/", "", false)
	if err != nil {
		return response.ExploreResponse{}, err
	}

	result := response.ExploreResponse{}
	err = json.Unmarshal(body, &result)

	return result, err
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

	bytes, _ := json.Marshal(Data)

	return insta.sendRequest("accounts/change_password/", generateSignature(string(bytes)), false)
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
		resp.Users = append(resp.Users, temp_resp.Users...)
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

func (insta *Instagram) GetFollowingRecentActivity() ([]byte, error) {
	return insta.sendRequest("news/?", "", false)
}

func (insta *Instagram) SearchUsername(query string) (response.SearchUserResponse, error) {
	body, err := insta.sendRequest("users/search/?ig_sig_key_version="+GOINSTA_SIG_KEY_VERSION+"&is_typeahead=true&query="+url.QueryEscape(query)+"&rank_token="+insta.Informations.RankToken, "", false)
	if err != nil {
		return response.SearchUserResponse{}, err
	}

	result := response.SearchUserResponse{}
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
	recipients, err := json.Marshal([][]string{{recipient}})
	if err != nil {
		return response.DirectMessageResponse{}, err
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
		return response.DirectMessageResponse{}, err
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
		return response.DirectMessageResponse{}, err
	}
	defer resp.Body.Close()

	insta.cookie = resp.Header.Get("Set-Cookie")

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return response.DirectMessageResponse{}, fmt.Errorf(string(body))
	}

	result := response.DirectMessageResponse{}
	json.Unmarshal(body, &result)
	return result, nil
}

func (insta *Instagram) GetReelsTrayFeed() {
	insta.sendRequest("feed/reels_tray/", "", false)
}

func (insta *Instagram) GetTrayFeeds(id string) {
	insta.sendRequest("feed/reels_tray/?tray_session_id="+id, "", false)
}

func (insta *Instagram) GetUserStories(id string) {
	insta.sendRequest("feed/user/"+id+"/reel_media", "", false)
}
