package goinsta

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Users struct {
	inst *Instagram

	// It's a bit confusing have the same structure
	// in the Instagram strucure and in the multiple users
	// calls

	err      error
	endpoint string

	Status   string `json:"status"`
	BigList  bool   `json:"big_list"`
	Users    []User `json:"users"`
	PageSize int    `json:"page_size"`
	NextID   string `json:"next_max_id"`
}

func newUsers(inst *Instagram) *Users {
	users := &Users{inst: inst}

	return users
}

// SetInstagram sets new instagram to user structure
func (users *Users) SetInstagram(inst *Instagram) {
	users.inst = inst
}

var ErrNoMore = errors.New("User list end reached")

// Next allows to paginate after calling:
// Account.Follow* and User.Follow*
//
// New user list is stored inside Users
//
// returns false when list reach the end.
func (users *Users) Next() bool {
	if users.err != nil {
		return false
	}

	insta := users.inst
	endpoint := users.endpoint

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id":             users.NextID,
				"ig_sig_key_version": goInstaSigKeyVersion,
				"rank_token":         insta.rankToken,
			},
		},
	)
	if err == nil {
		usrs := Users{}
		err = json.Unmarshal(body, &usrs)
		if err == nil {
			*users = usrs
			if !usrs.BigList || usrs.NextID == "" {
				users.err = ErrNoMore
			}
			users.inst = insta
			users.endpoint = endpoint
			return true
		}
	}
	users.err = err
	return false
}

type userResp struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}

// User is the representation of instagram's user profile
type User struct {
	inst *Instagram

	ID                         int64        `json:"pk"`
	Username                   string       `json:"username"`
	FullName                   string       `json:"full_name"`
	Biography                  string       `json:"biography"`
	ProfilePicURL              string       `json:"profile_pic_url"`
	Email                      string       `json:"email"`
	PhoneNumber                string       `json:"phone_number"`
	IsBusiness                 bool         `json:"is_business"`
	Gender                     int          `json:"gender"`
	ProfilePicID               string       `json:"profile_pic_id"`
	HasAnonymousProfilePicture bool         `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool         `json:"is_private"`
	IsUnpublished              bool         `json:"is_unpublished"`
	AllowedCommenterType       string       `json:"allowed_commenter_type"`
	IsVerified                 bool         `json:"is_verified"`
	MediaCount                 int          `json:"media_count"`
	FollowerCount              int          `json:"follower_count"`
	FollowingCount             int          `json:"following_count"`
	GeoMediaCount              int          `json:"geo_media_count"`
	ExternalURL                string       `json:"external_url"`
	HasBiographyTranslation    bool         `json:"has_biography_translation"`
	ExternalLynxURL            string       `json:"external_lynx_url"`
	HdProfilePicURLInfo        PicURLInfo   `json:"hd_profile_pic_url_info"`
	HdProfilePicVersions       []PicURLInfo `json:"hd_profile_pic_versions"`
	UsertagsCount              int          `json:"usertags_count"`
	HasChaining                bool         `json:"has_chaining"`
	IsFavorite                 bool         `json:"is_favorite"`
	ReelAutoArchive            string       `json:"reel_auto_archive"`
	School                     School       `json:"school"`
	PublicEmail                string       `json:"public_email"`
	PublicPhoneNumber          string       `json:"public_phone_number"`
	PublicPhoneCountryCode     string       `json:"public_phone_country_code"`
	ContactPhoneNumber         string       `json:"contact_phone_number"`
	CityID                     int64        `json:"city_id"`
	CityName                   string       `json:"city_name"`
	AddressStreet              string       `json:"address_street"`
	DirectMessaging            string       `json:"direct_messaging"`
	Latitude                   float64      `json:"latitude"`
	Longitude                  float64      `json:"longitude"`
	Category                   string       `json:"category"`
	BusinessContactMethod      string       `json:"business_contact_method"`
	IsCallToActionEnabled      bool         `json:"is_call_to_action_enabled"`
	FbPageCallToActionID       string       `json:"fb_page_call_to_action_id"`
	Zip                        string       `json:"zip"`
	Friendship                 Friendship   `json:"friendship_status"`
}

// Following returns a list of user following.
//
// Users.Next can be used to paginate
func (user *User) Following() *Users {
	users := &Users{}
	users.inst = user.inst
	users.endpoint = fmt.Sprintf(urlFollowing, user.ID)
	return users
}

// Followers returns a list of user followers.
//
// Users.Next can be used to paginate
func (user *User) Followers() *Users {
	users := &Users{}
	users.inst = user.inst
	users.endpoint = fmt.Sprintf(urlFollowers, user.ID)
	return users
}

// Block blocks user
//
// This function updates current User.Friendship structure.
func (user *User) Block() error {
	insta := user.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"user_id": user.ID,
		},
	)
	if err == nil {
		body, err := insta.sendRequest(
			&reqOptions{
				Endpoint: fmt.Sprintf(urlUserBlock, user.ID),
				Query:    generateSignature(data),
				IsPost:   true,
			},
		)
		if err == nil {
			resp := friendResp{}
			err = json.Unmarshal(body, &resp)
			user.Friendship = resp.Friendship
		}
	}
	return err
}

// Unblock unblocks user
//
// This function updates current User.Friendship structure.
func (user *User) Unblock() error {
	insta := user.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"user_id": user.ID,
		},
	)
	if err == nil {
		body, err := insta.sendRequest(
			&reqOptions{
				Endpoint: fmt.Sprintf(urlUserUnblock, user.ID),
				Query:    generateSignature(data),
				IsPost:   true,
			},
		)
		if err == nil {
			resp := friendResp{}
			err = json.Unmarshal(body, &resp)
			user.Friendship = resp.Friendship
		}
	}
	return err
}

// Follow started following some user
//
// This function performs a follow call. If user is private
// you have to wait until he/she accepts you.
//
// If the account is public User.Friendship will be updated
func (user *User) Follow() error {
	insta := user.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"user_id": user.ID,
		},
	)
	if err == nil {
		body, err := insta.sendRequest(
			&reqOptions{
				Endpoint: fmt.Sprintf(urlUserFollow, user.ID),
				Query:    generateSignature(data),
				IsPost:   true,
			},
		)
		if err == nil {
			resp := friendResp{}
			err = json.Unmarshal(body, &resp)
			user.Friendship = resp.Friendship
		}
	}
	return err
}

// Unfollow unfollows user
//
// User.Friendship will be updated
func (user *User) Unfollow() error {
	insta := user.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"user_id": user.ID,
		},
	)
	if err == nil {
		body, err := insta.sendRequest(
			&reqOptions{
				Endpoint: fmt.Sprintf(urlUserUnfollow, user.ID),
				Query:    generateSignature(data),
				IsPost:   true,
			},
		)
		if err == nil {
			resp := friendResp{}
			err = json.Unmarshal(body, &resp)
			user.Friendship = resp.Friendship
		}
	}
	return err
}

func (user *User) friendShip() error {
	insta := user.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"user_id": user.ID,
		},
	)
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlFriendship, user.ID),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		err = json.Unmarshal(body, &user.Friendship)
	}
	return err
}

// Feed returns user feeds (media)
//
// minTime is the minimum timestamp of media.
//
// For pagination use FeedMedia.Next()
func (user *User) Feed(minTime []byte) *FeedMedia {
	insta := user.inst
	timestamp := b2s(minTime)

	media := &FeedMedia{}
	media.timestamp = timestamp
	media.inst = insta
	media.endpoint = urlUserFeed
	media.uid = user.ID
	return media
}

// Stories returns user stories
//
// Use StoryMedia.Next for pagination.
//
// See example: examples/user/stories.go
func (user *User) Stories() *StoryMedia {
	media := &StoryMedia{}
	media.uid = user.ID
	media.inst = user.inst
	media.endpoint = urlUserStories
	return media
}

// Tags returns media where user is tagged in
//
// For pagination use FeedMedia.Next()
func (user *User) Tags(minTimestamp []byte) (*FeedMedia, error) {
	timestamp := b2s(minTimestamp)
	body, err := user.inst.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlUserTags, user.ID),
			Query: map[string]string{
				"max_id":         "",
				"rank_token":     user.inst.rankToken,
				"min_timestamp":  timestamp,
				"ranked_content": "true",
			},
		},
	)
	if err != nil {
		return nil, err
	}

	media := &FeedMedia{}
	err = json.Unmarshal(body, media)
	media.inst = user.inst
	media.endpoint = urlUserTags
	media.uid = user.ID
	return media, err
}
