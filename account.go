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
type Account struct {
	// Activity is recent activity
	//Activity *Activity
	// Tray is your disponible friend's stories
	//Tray *Tray

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
func (account *Account) Followers() (*Users, error) {
	endpoint := fmt.Sprintf(urlFollowers, account.ID)
	body, err := account.inst.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id":             "",
				"ig_sig_key_version": goInstaSigKeyVersion,
				"rank_token":         account.inst.rankToken,
			},
		},
	)
	if err == nil {
		users := &Users{}
		err = json.Unmarshal(body, users)
		users.inst = account.inst
		users.endpoint = endpoint
		return users, err
	}
	return nil, err
}

// Following returns a list of user following.
//
// Users.Next can be used to paginate
func (account *Account) Following() (*Users, error) {
	endpoint := fmt.Sprintf(urlFollowing, account.ID)
	body, err := account.inst.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id":             "",
				"ig_sig_key_version": goInstaSigKeyVersion,
				"rank_token":         account.inst.rankToken,
			},
		},
	)
	if err == nil {
		users := &Users{}
		err = json.Unmarshal(body, users)
		users.inst = account.inst
		users.endpoint = endpoint
		return users, err
	}
	return nil, err
}
