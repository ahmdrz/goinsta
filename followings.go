package goinsta

import (
	"encoding/json"
	"fmt"
)

func (f *InstagramFollowings) User(userID int64, options ...string) (output FollowerAndFollowingResponse, err error) {
	if len(options) > 1 {
		return output, fmt.Errorf("Bad input as options , use only maxID if you need, inputs are %v", options)
	}
	maxID := ""
	if len(options) == 1 {
		maxID = options[0]
	}

	body, err := f.instagram.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf("friendships/%d/following/", userID),
		Query: map[string]string{
			"max_id":             maxID,
			"ig_sig_key_version": GOINSTA_SIG_KEY_VERSION,
			"rank_token":         f.instagram.rankToken,
		},
	})
	if err != nil {
		return output, err
	}

	err = json.Unmarshal(body, &output)

	return output, err
}

func (f *InstagramFollowings) Current(maxID ...string) (FollowerAndFollowingResponse, error) {
	return f.User(f.instagram.CurrentUser.ID, maxID...)
}
