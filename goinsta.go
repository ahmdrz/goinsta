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
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ahmdrz/goinsta/response"
)

// GetSessions return current instagram session and cookies
// Maybe need for webpages that use this API
func (insta *Instagram) GetSessions(url *url.URL) []*http.Cookie {
	return insta.Cookiejar.Cookies(url)
}

// SetCookies can enable us to set cookie, it'll be help for webpage that use this API without Login-again.
func (insta *Instagram) SetCookies(url *url.URL, cookies []*http.Cookie) error {
	if insta.Cookiejar == nil {
		var err error
		insta.Cookiejar, err = cookiejar.New(nil) //newJar()
		if err != nil {
			return err
		}
	}
	insta.Cookiejar.SetCookies(url, cookies)
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
	GOINSTA_USER_AGENT      = "Instagram 10.26.0 Android (18/4.3; 320dpi; 720x1280; Xiaomi; HM 1SW; armani; qcom; en_US)"
	GOINSTA_IG_SIG_KEY      = "4f8732eb9ba7d1c8e8897a75d6474d4eb3f5279137431b2aafb71fafe2abe178"
	GOINSTA_EXPERIMENTS     = "ig_promote_reach_objective_fix_universe,ig_android_universe_video_production,ig_search_client_h1_2017_holdout,ig_android_live_follow_from_comments_universe,ig_android_carousel_non_square_creation,ig_android_live_analytics,ig_android_follow_all_dialog_confirmation_copy,ig_android_stories_server_coverframe,ig_android_video_captions_universe,ig_android_offline_location_feed,ig_android_direct_inbox_retry_seen_state,ig_android_ontact_invite_universe,ig_android_live_broadcast_blacklist,ig_android_insta_video_reconnect_viewers,ig_android_ad_async_ads_universe,ig_android_search_clear_layout_universe,ig_android_shopping_reporting,ig_android_stories_surface_universe,ig_android_verified_comments_universe,ig_android_preload_media_ahead_in_current_reel,android_instagram_prefetch_suggestions_universe,ig_android_reel_viewer_fetch_missing_reels_universe,ig_android_direct_search_share_sheet_universe,ig_android_business_promote_tooltip,ig_android_direct_blue_tab,ig_android_async_network_tweak_universe,ig_android_elevate_main_thread_priority_universe,ig_android_stories_gallery_nux,ig_android_instavideo_remove_nux_comments,ig_video_copyright_whitelist,ig_react_native_inline_insights_with_relay,ig_android_direct_thread_message_animation,ig_android_draw_rainbow_client_universe,ig_android_direct_link_style,ig_android_live_heart_enhancements_universe,ig_android_rtc_reshare,ig_android_preload_item_count_in_reel_viewer_buffer,ig_android_users_bootstrap_service,ig_android_auto_retry_post_mode,ig_android_shopping,ig_android_main_feed_seen_state_dont_send_info_on_tail_load,ig_fbns_preload_default,ig_android_gesture_dismiss_reel_viewer,ig_android_tool_tip,ig_android_ad_logger_funnel_logging_universe,ig_android_gallery_grid_column_count_universe,ig_android_business_new_ads_payment_universe,ig_android_direct_links,ig_android_audience_control,ig_android_live_encore_consumption_settings_universe,ig_perf_android_holdout,ig_android_cache_contact_import_list,ig_android_links_receivers,ig_android_ad_impression_backtest,ig_android_list_redesign,ig_android_stories_separate_overlay_creation,ig_android_stop_video_recording_fix_universe,ig_android_render_video_segmentation,ig_android_live_encore_reel_chaining_universe,ig_android_sync_on_background_enhanced_10_25,ig_android_immersive_viewer,ig_android_mqtt_skywalker,ig_fbns_push,ig_android_ad_watchmore_overlay_universe,ig_android_react_native_universe,ig_android_profile_tabs_redesign_universe,ig_android_live_consumption_abr,ig_android_story_viewer_social_context,ig_android_hide_post_in_feed,ig_android_video_loopcount_int,ig_android_enable_main_feed_reel_tray_preloading,ig_android_camera_upsell_dialog,ig_android_ad_watchbrowse_universe,ig_android_internal_research_settings,ig_android_search_people_tag_universe,ig_android_react_native_ota,ig_android_enable_concurrent_request,ig_android_react_native_stories_grid_view,ig_android_business_stories_inline_insights,ig_android_log_mediacodec_info,ig_android_direct_expiring_media_loading_errors,ig_video_use_sve_universe,ig_android_cold_start_feed_request,ig_android_enable_zero_rating,ig_android_reverse_audio,ig_android_branded_content_three_line_ui_universe,ig_android_live_encore_production_universe,ig_stories_music_sticker,ig_android_stories_teach_gallery_location,ig_android_http_stack_experiment_2017,ig_android_stories_device_tilt,ig_android_pending_request_search_bar,ig_android_fb_topsearch_sgp_fork_request,ig_android_seen_state_with_view_info,ig_android_animation_perf_reporter_timeout,ig_android_new_block_flow,ig_android_story_tray_title_play_all_v2,ig_android_direct_address_links,ig_android_stories_archive_universe,ig_android_save_collections_cover_photo,ig_android_live_webrtc_livewith_production,ig_android_sign_video_url,ig_android_stories_video_prefetch_kb,ig_android_stories_create_flow_favorites_tooltip,ig_android_live_stop_broadcast_on_404,ig_android_live_viewer_invite_universe,ig_android_promotion_feedback_channel,ig_android_render_iframe_interval,ig_android_accessibility_logging_universe,ig_android_camera_shortcut_universe,ig_android_use_one_cookie_store_per_user_override,ig_profile_holdout_2017_universe,ig_android_stories_server_brushes,ig_android_ad_media_url_logging_universe,ig_android_shopping_tag_nux_text_universe,ig_android_comments_single_reply_universe,ig_android_stories_video_loading_spinner_improvements,ig_android_collections_cache,ig_android_comment_api_spam_universe,ig_android_facebook_twitter_profile_photos,ig_android_shopping_tag_creation_universe,ig_story_camera_reverse_video_experiment,ig_android_direct_bump_selected_recipients,ig_android_ad_cta_haptic_feedback_universe,ig_android_vertical_share_sheet_experiment,ig_android_family_bridge_share,ig_android_search,ig_android_insta_video_consumption_titles,ig_android_stories_gallery_preview_button,ig_android_fb_auth_education,ig_android_camera_universe,ig_android_me_only_universe,ig_android_instavideo_audio_only_mode,ig_android_user_profile_chaining_icon,ig_android_live_video_reactions_consumption_universe,ig_android_stories_hashtag_text,ig_android_post_live_badge_universe,ig_android_swipe_fragment_container,ig_android_search_users_universe,ig_android_live_save_to_camera_roll_universe,ig_creation_growth_holdout,ig_android_sticker_region_tracking,ig_android_unified_inbox,ig_android_live_new_watch_time,ig_android_offline_main_feed_10_11,ig_import_biz_contact_to_page,ig_android_live_encore_consumption_universe,ig_android_experimental_filters,ig_android_search_client_matching_2,ig_android_react_native_inline_insights_v2,ig_android_business_conversion_value_prop_v2,ig_android_redirect_to_low_latency_universe,ig_android_ad_show_new_awr_universe,ig_family_bridges_holdout_universe,ig_android_background_explore_fetch,ig_android_following_follower_social_context,ig_android_video_keep_screen_on,ig_android_ad_leadgen_relay_modern,ig_android_profile_photo_as_media,ig_android_insta_video_consumption_infra,ig_android_ad_watchlead_universe,ig_android_direct_prefetch_direct_story_json,ig_android_shopping_react_native,ig_android_top_live_profile_pics_universe,ig_android_direct_phone_number_links,ig_android_stories_weblink_creation,ig_android_direct_search_new_thread_universe,ig_android_histogram_reporter,ig_android_direct_on_profile_universe,ig_android_network_cancellation,ig_android_background_reel_fetch,ig_android_react_native_insights,ig_android_insta_video_audio_encoder,ig_android_family_bridge_bookmarks,ig_android_data_usage_network_layer,ig_android_universal_instagram_deep_links,ig_android_dash_for_vod_universe,ig_android_modular_tab_discover_people_redesign,ig_android_mas_sticker_upsell_dialog_universe,ig_android_ad_add_per_event_counter_to_logging_event,ig_android_sticky_header_top_chrome_optimization,ig_android_rtl,ig_android_biz_conversion_page_pre_select,ig_android_promote_from_profile_button,ig_android_live_broadcaster_invite_universe,ig_android_share_spinner,ig_android_text_action,ig_android_own_reel_title_universe,ig_promotions_unit_in_insights_landing_page,ig_android_business_settings_header_univ,ig_android_save_longpress_tooltip,ig_android_constrain_image_size_universe,ig_android_business_new_graphql_endpoint_universe,ig_ranking_following,ig_android_stories_profile_camera_entry_point,ig_android_universe_reel_video_production,ig_android_power_metrics,ig_android_sfplt,ig_android_offline_hashtag_feed,ig_android_live_skin_smooth,ig_android_direct_inbox_search,ig_android_stories_posting_offline_ui,ig_android_sidecar_video_upload_universe,ig_android_promotion_manager_entry_point_universe,ig_android_direct_reply_audience_upgrade,ig_android_swipe_navigation_x_angle_universe,ig_android_offline_mode_holdout,ig_android_live_send_user_location,ig_android_direct_fetch_before_push_notif,ig_android_non_square_first,ig_android_insta_video_drawing,ig_android_swipeablefilters_universe,ig_android_live_notification_control_universe,ig_android_analytics_logger_running_background_universe,ig_android_save_all,ig_android_reel_viewer_data_buffer_size,ig_direct_quality_holdout_universe,ig_android_family_bridge_discover,ig_android_react_native_restart_after_error_universe,ig_android_startup_manager,ig_story_tray_peek_content_universe,ig_android_profile,ig_android_high_res_upload_2,ig_android_http_service_same_thread,ig_android_scroll_to_dismiss_keyboard,ig_android_remove_followers_universe,ig_android_skip_video_render,ig_android_story_timestamps,ig_android_live_viewer_comment_prompt_universe,ig_profile_holdout_universe,ig_android_react_native_insights_grid_view,ig_stories_selfie_sticker,ig_android_stories_reply_composer_redesign,ig_android_streamline_page_creation,ig_explore_netego,ig_android_ig4b_connect_fb_button_universe,ig_android_feed_util_rect_optimization,ig_android_rendering_controls,ig_android_os_version_blocking,ig_android_encoder_width_safe_multiple_16,ig_search_new_bootstrap_holdout_universe,ig_android_snippets_profile_nux,ig_android_e2e_optimization_universe,ig_android_comments_logging_universe,ig_shopping_insights,ig_android_save_collections,ig_android_live_see_fewer_videos_like_this_universe,ig_android_show_new_contact_import_dialog,ig_android_live_view_profile_from_comments_universe,ig_fbns_blocked,ig_formats_and_feedbacks_holdout_universe,ig_android_reduce_view_pager_buffer,ig_android_instavideo_periodic_notif,ig_search_user_auto_complete_cache_sync_ttl,ig_android_marauder_update_frequency,ig_android_suggest_password_reset_on_oneclick_login,ig_android_promotion_entry_from_ads_manager_universe,ig_android_live_special_codec_size_list,ig_android_enable_share_to_messenger,ig_android_background_main_feed_fetch,ig_android_live_video_reactions_creation_universe,ig_android_channels_home,ig_android_sidecar_gallery_universe,ig_android_upload_reliability_universe,ig_migrate_mediav2_universe,ig_android_insta_video_broadcaster_infra_perf,ig_android_business_conversion_social_context,android_ig_fbns_kill_switch,ig_android_live_webrtc_livewith_consumption,ig_android_destroy_swipe_fragment,ig_android_react_native_universe_kill_switch,ig_android_stories_book_universe,ig_android_all_videoplayback_persisting_sound,ig_android_draw_eraser_universe,ig_direct_search_new_bootstrap_holdout_universe,ig_android_cache_layer_bytes_threshold,ig_android_search_hash_tag_and_username_universe,ig_android_business_promotion,ig_android_direct_search_recipients_controller_universe,ig_android_ad_show_full_name_universe,ig_android_anrwatchdog,ig_android_qp_kill_switch,ig_android_2fac,ig_direct_bypass_group_size_limit_universe,ig_android_promote_simplified_flow,ig_android_share_to_whatsapp,ig_android_hide_bottom_nav_bar_on_discover_people,ig_fbns_dump_ids,ig_android_hands_free_before_reverse,ig_android_skywalker_live_event_start_end,ig_android_live_join_comment_ui_change,ig_android_direct_search_story_recipients_universe,ig_android_direct_full_size_gallery_upload,ig_android_ad_browser_gesture_control,ig_channel_server_experiments,ig_android_video_cover_frame_from_original_as_fallback,ig_android_ad_watchinstall_universe,ig_android_ad_viewability_logging_universe,ig_android_new_optic,ig_android_direct_visual_replies,ig_android_stories_search_reel_mentions_universe,ig_android_threaded_comments_universe,ig_android_mark_reel_seen_on_Swipe_forward,ig_internal_ui_for_lazy_loaded_modules_experiment,ig_fbns_shared,ig_android_capture_slowmo_mode,ig_android_live_viewers_list_search_bar,ig_android_video_single_surface,ig_android_offline_reel_feed,ig_android_video_download_logging,ig_android_last_edits,ig_android_exoplayer_4142,ig_android_post_live_viewer_count_privacy_universe,ig_android_activity_feed_click_state,ig_android_snippets_haptic_feedback,ig_android_gl_drawing_marks_after_undo_backing,ig_android_mark_seen_state_on_viewed_impression,ig_android_live_backgrounded_reminder_universe,ig_android_live_hide_viewer_nux_universe,ig_android_live_monotonic_pts,ig_android_search_top_search_surface_universe,ig_android_user_detail_endpoint,ig_android_location_media_count_exp_ig,ig_android_comment_tweaks_universe,ig_android_ad_watchmore_entry_point_universe,ig_android_top_live_notification_universe,ig_android_add_to_last_post,ig_save_insights,ig_android_live_enhanced_end_screen_universe,ig_android_ad_add_counter_to_logging_event,ig_android_blue_token_conversion_universe,ig_android_exoplayer_settings,ig_android_progressive_jpeg,ig_android_offline_story_stickers,ig_android_gqls_typing_indicator,ig_android_chaining_button_tooltip,ig_android_video_prefetch_for_connectivity_type,ig_android_use_exo_cache_for_progressive,ig_android_samsung_app_badging,ig_android_ad_holdout_watchandmore_universe,ig_android_offline_commenting,ig_direct_stories_recipient_picker_button,ig_insights_feedback_channel_universe,ig_android_insta_video_abr_resize,ig_android_insta_video_sound_always_on"
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
	insta.Proxy = proxy
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
	insta.GetV2Inbox()
	insta.GetRecentActivity()
	insta.GetReelsTrayFeed()

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

// UploadPhotoFromReader can upload your photo stored in io.Reader with any quality , better to use 87
func (insta *Instagram) UploadPhotoFromReader(photo io.Reader, photo_caption string, upload_id int64, quality int, filter_type int) (response.UploadPhotoResponse, error) {
	photo_name := fmt.Sprintf("pending_media_%d.jpg", upload_id)

	//multipart request body
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	w.WriteField("upload_id", strconv.FormatInt(upload_id, 10))
	w.WriteField("_uuid", insta.Informations.UUID)
	w.WriteField("_csrftoken", insta.Informations.Token)
	w.WriteField("image_compression", `{"lib_name":"jt","lib_version":"1.3.0","quality":"`+strconv.Itoa(quality)+`"}`)

	fw, err := w.CreateFormFile("photo", photo_name)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}

	var buf bytes.Buffer

	rdr := io.TeeReader(photo, &buf)

	if _, err = io.Copy(fw, rdr); err != nil {
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
		Jar: insta.Cookiejar,
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
		w, h, err := getImageDimensionFromReader(&buf)
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

// UploadPhoto can upload your photo file, stored in filesystem with any quality , better to use 87
func (insta *Instagram) UploadPhoto(photo_path string, photo_caption string, upload_id int64, quality int, filter_type int) (response.UploadPhotoResponse, error) {
	f, err := os.Open(photo_path)
	if err != nil {
		return response.UploadPhotoResponse{}, err
	}
	defer f.Close()

	return insta.UploadPhotoFromReader(f, photo_caption, upload_id, quality, filter_type)
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

func (insta *Instagram) DisableComments(mediaID string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/disable_comments/", mediaID),
		PostData: generateSignature(data),
	})
}

func (insta *Instagram) EnableComments(mediaID string) ([]byte, error) {
	data, err := insta.prepareData(map[string]interface{}{
		"media_id": mediaID,
	})
	if err != nil {
		return []byte{}, err
	}

	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("media/%s/enable_comments/", mediaID),
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

func (insta *Instagram) Timeline(maxID string) (r response.FeedsResponse, err error) {
	data, err := insta.sendRequest(&reqOptions{
		Endpoint: "feed/timeline/",
		Query: map[string]string{
			"max_id":         maxID,
			"rank_token":     insta.Informations.RankToken,
			"ranked_content": "true",
		},
	})
	if err == nil {
		err = json.Unmarshal(data, &r)
	}

	return
}

// getImageDimensionFromReader return image dimension , types is .jpg and .png
func getImageDimensionFromReader(rdr io.Reader) (int, int, error) {
	image, _, err := image.DecodeConfig(rdr)
	if err != nil {
		return 0, 0, err
	}
	return image.Width, image.Height, nil
}

// getImageDimension return image dimension , types is .jpg and .png
func getImageDimension(imagePath string) (int, int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	return getImageDimensionFromReader(file)
}

func (insta *Instagram) SelfUserFollowers(maxID string) (response.UsersResponse, error) {
	return insta.UserFollowers(insta.LoggedInUser.ID, maxID)
}

func (insta *Instagram) SelfUserFollowing(maxID string) (response.UsersResponse, error) {
	return insta.UserFollowing(insta.LoggedInUser.ID, maxID)
}

func (insta *Instagram) SelfTotalUserFollowing() (response.UsersResponse, error) {
	return insta.TotalUserFollowing(insta.LoggedInUser.ID)
}

func (insta *Instagram) SelfTotalUserFollowers() (response.UsersResponse, error) {
	return insta.TotalUserFollowers(insta.LoggedInUser.ID)
}

func (insta *Instagram) TotalUserFollowing(userID int64) (response.UsersResponse, error) {
	resp := response.UsersResponse{}
	for {
		temp_resp, err := insta.UserFollowing(userID, resp.NextMaxID)
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

func (insta *Instagram) TotalUserFollowers(userID int64) (response.UsersResponse, error) {
	resp := response.UsersResponse{}
	for {
		temp_resp, err := insta.UserFollowers(userID, resp.NextMaxID)
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

func (insta *Instagram) SearchTags(query string) (response.SearchTagsResponse, error) {
	result := response.SearchTagsResponse{}
	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "tags/search/",
		Query: map[string]string{
			"is_typeahead": "true",
			"rank_token":   insta.Informations.RankToken,
			"q":            query,
		},
	})
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)

	return result, err
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

	bytes, err := insta.sendSimpleRequest("feed/user/%d/story/", userID)
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
