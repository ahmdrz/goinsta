package goinsta

import (
	"encoding/json"
	"fmt"
)

type accountResp struct {
	Status  string  `json:"status"`
	Account Account `json:"logged_in_user"`
}

// Account is personal account object
//
// See examples: examples/account/*
type Account struct {
	inst *Instagram

	CanSeeOrganicInsights      bool    `json:"can_see_organic_insights"`
	ShowInsightsTerms          bool    `json:"show_insights_terms"`
	IsBusiness                 bool    `json:"is_business"`
	Nametag                    Nametag `json:"nametag"`
	ID                         int64   `json:"pk"`
	Username                   string  `json:"username"`
	FullName                   string  `json:"full_name"`
	HasAnonymousProfilePicture bool    `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool    `json:"is_private"`
	IsVerified                 bool    `json:"is_verified"`
	ProfilePicURL              string  `json:"profile_pic_url"`
	ProfilePicID               string  `json:"profile_pic_id"`
	AllowedCommenterType       string  `json:"allowed_commenter_type"`
	ReelAutoArchive            string  `json:"reel_auto_archive"`
	AllowContactsSync          bool    `json:"allow_contacts_sync"`
	PhoneNumber                string  `json:"phone_number"`
	CanBoostPost               bool    `json:"can_boost_post"`
}

// Sync updates account information
func (account *Account) Sync() error {
	insta := account.inst
	data, err := insta.prepareData()
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSyncProfile,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := profResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*account = resp.Account
			account.inst = insta
		}
	}
	return err
}

// ChangePassword changes current password.
//
// GoInsta does not store current instagram password (for security reasons)
// If you want to change your password you must parse old and new password.
//
// See example: examples/account/changePass.go
func (account *Account) ChangePassword(old, new string) error {
	insta := account.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"old_password":  old,
			"new_password1": new,
			"new_password2": new,
		},
	)
	if err == nil {
		_, err = insta.sendRequest(
			&reqOptions{
				Endpoint: urlChangePass,
				Query:    generateSignature(data),
				IsPost:   true,
			},
		)
	}
	return err
}

type profResp struct {
	Status  string  `json:"status"`
	Account Account `json:"user"`
}

// RemoveProfilePic removes current profile picture
//
// This function updates current Account information.
//
// See example: examples/account/removeProfilePic.go
func (account *Account) RemoveProfilePic() error {
	insta := account.inst
	data, err := insta.prepareData()
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlRemoveProfPic,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := profResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*account = resp.Account
			account.inst = insta
		}
	}
	return err
}

// SetPrivate sets account to private mode.
//
// This function updates current Account information.
//
// See example: examples/account/setPrivate.go
func (account *Account) SetPrivate() error {
	insta := account.inst
	data, err := insta.prepareData()
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSetPrivate,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := profResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*account = resp.Account
			account.inst = insta
		}
	}
	return err
}

// SetPublic sets account to public mode.
//
// This function updates current Account information.
//
// See example: examples/account/setPublic.go
func (account *Account) SetPublic() error {
	insta := account.inst
	data, err := insta.prepareData()
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSetPublic,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := profResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			*account = resp.Account
			account.inst = insta
		}
	}
	return err
}

// Followers returns a list of user followers.
//
// Users.Next can be used to paginate
//
// See example: examples/account/followers.go
func (account *Account) Followers() *Users {
	endpoint := fmt.Sprintf(urlFollowers, account.ID)
	users := &Users{}
	users.inst = account.inst
	users.endpoint = endpoint
	return users
}

// Following returns a list of user following.
//
// Users.Next can be used to paginate
//
// See example: examples/account/following.go
func (account *Account) Following() *Users {
	endpoint := fmt.Sprintf(urlFollowing, account.ID)
	users := &Users{}
	users.inst = account.inst
	users.endpoint = endpoint
	return users
}

// Feed returns current account feed
//
// minTime is the minimum timestamp of media.
//
// For pagination use FeedMedia.Next()
func (account *Account) Feed(minTime []byte) (*FeedMedia, error) {
	insta := account.inst
	timestamp := b2s(minTime)

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlUserFeed, account.ID),
			Query: map[string]string{
				"max_id":         "",
				"rank_token":     insta.rankToken,
				"min_timestamp":  timestamp,
				"ranked_content": "true",
			},
		},
	)
	if err == nil {
		media := &FeedMedia{}
		err = json.Unmarshal(body, media)
		media.inst = insta
		media.endpoint = urlUserFeed
		media.uid = account.ID
		return media, err
	}
	return nil, err
}

// Stories returns account stories.
//
// Use StoryMedia.Next for pagination.
//
// See example: examples/account/stories.go
func (account *Account) Stories() *StoryMedia {
	media := &StoryMedia{}
	media.uid = account.ID
	media.inst = account.inst
	media.endpoint = urlUserStories
	return media
}

// Tags returns media where account is tagged in
//
// For pagination use FeedMedia.Next()
func (account *Account) Tags(minTimestamp []byte) (*FeedMedia, error) {
	timestamp := b2s(minTimestamp)
	body, err := account.inst.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlUserTags, account.ID),
			Query: map[string]string{
				"max_id":         "",
				"rank_token":     account.inst.rankToken,
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
	media.inst = account.inst
	media.endpoint = urlUserTags
	media.uid = account.ID
	return media, err
}
