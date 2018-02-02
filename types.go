package goinsta

import (
	"net/http"
	"net/http/cookiejar"
)

type Instagram struct {
	username  string
	password  string
	deviceID  string
	uuid      string
	rankToken string
	token     string
	phoneID   string

	FriendShip FriendShip
	Users      Users

	CurrentUser CurrentUser

	isLoggedIn bool
	cookiejar  *cookiejar.Jar
	transport  http.Transport
	proxy      string
}

type FriendShip struct {
	instagram *Instagram
}

type Users struct {
	instagram *Instagram
}

type UserResponse struct {
	ID                         int64  `json:"pk"`
	Username                   string `json:"username"`
	FullName                   string `json:"full_name"`
	IsPrivate                  bool   `json:"is_private"`
	ProfilePicURL              string `json:"profile_pic_url"`
	ProfilePicID               string `json:"profile_pic_id"`
	IsVerified                 bool   `json:"is_verified"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	IsBusiness                 bool   `json:"is_business"`
	CanSeeOrganicInsights      bool   `json:"can_see_organic_insights"`
	ShowInsightsTerms          bool   `json:"show_insights_terms"`
	AllowContactsSync          bool   `json:"allow_contacts_sync"`
	PhoneNumber                string `json:"phone_number"`
	CountryCode                int    `json:"country_code"`
	NationalNumber             int64  `json:"national_number"`
	ReelAutoArchive            string `json:"reel_auto_archive"`
	MediaCount                 int    `json:"media_count"`
	GeoMediaCount              int    `json:"geo_media_count"`
	FollowerCount              int    `json:"follower_count"`
	FollowingCount             int    `json:"following_count"`
	Biography                  string `json:"biography"`
	ExternalURL                string `json:"external_url"`
	UsertagsCount              int    `json:"usertags_count"`
	IsFavorite                 bool   `json:"is_favorite"`
	HasChaining                bool   `json:"has_chaining"`
	HdProfilePicVersions       []struct {
		Width  int    `json:"width"`
		Height int    `json:"height"`
		URL    string `json:"url"`
	} `json:"hd_profile_pic_versions"`
	HdProfilePicURLInfo struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"hd_profile_pic_url_info"`
	ProfileContext                 string `json:"profile_context"`
	ProfileContextLinksWithUserIds []struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"profile_context_links_with_user_ids"`
	ProfileContextMutualFollowIds []interface{} `json:"profile_context_mutual_follow_ids"`
}

type FollowerAndFollowingResponse struct {
	Users     []UserResponse `json:"users"`
	BigList   bool           `json:"big_list"`
	NextMaxID string         `json:"next_max_id"`
	PageSize  int            `json:"page_size"`
	Status    string         `json:"status"`
}

type FriendShipsShowResponse struct {
	Following       bool   `json:"following"`
	FollowedBy      bool   `json:"followed_by"`
	Blocking        bool   `json:"blocking"`
	IsPrivate       bool   `json:"is_private"`
	IncomingRequest bool   `json:"incoming_request"`
	OutgoingRequest bool   `json:"outgoing_request"`
	IsBlockingReel  bool   `json:"is_blocking_reel"`
	IsMutingReel    bool   `json:"is_muting_reel"`
	IsBestie        bool   `json:"is_bestie"`
	Status          string `json:"status"`
}

type FriendShipsPendingResponse struct {
	Users    []UserResponse `json:"users"`
	BigList  bool           `json:"big_list"`
	PageSize int            `json:"page_size"`
	Status   string         `json:"status"`
}

type CurrentUser struct {
	instagram *Instagram
	UserResponse
}
