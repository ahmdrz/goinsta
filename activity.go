package goinsta

import (
	"encoding/json"
	"strconv"
)

// Activity is the activity of your instagram account
type Activity struct {
	inst *Instagram
}

// Activities ...
type Activities struct {
	inst *Instagram

	err      error
	endpoint string

	AutoLoadMoreEnabled bool  `json:"auto_load_more_enabled"`
	NextID              int64 `json:"next_max_id"`
	Stories             []struct {
		Type      int `json:"type"`
		StoryType int `json:"story_type"`
		Args      struct {
			Text  string `json:"text"`
			Links []struct {
				Start int    `json:"start"`
				End   int    `json:"end"`
				Type  string `json:"type"`
				ID    string `json:"id"`
			} `json:"links"`
			ProfileID    int    `json:"profile_id"`
			ProfileImage string `json:"profile_image"`
			Media        []struct {
				ID    string `json:"id"`
				Image string `json:"image"`
			} `json:"media"`
			Timestamp int    `json:"timestamp"`
			Tuuid     string `json:"tuuid"`
		} `json:"args"`
		Counts struct {
		} `json:"counts"`
		Pk string `json:"pk"`
	} `json:"stories"`
	Status string `json:"status"`
}

// Next can be used to paginate over activities.
func (nact *Activities) Next() bool {
	// TODO: Bug
	if nact.err != nil {
		return false
	}
	insta := nact.inst

	endpoint := nact.endpoint

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id": strconv.FormatInt(nact.NextID, 10),
			},
		},
	)
	if err != nil {
		nact.err = err
		return false
	}

	nnact := Activities{}
	err = json.Unmarshal(body, &nnact)
	if err == nil {
		*nact = nnact
		nact.inst = insta
		nact.endpoint = endpoint
		if nact.NextID == 0 {
			nact.err = ErrNoMore
		}
		return true
	}
	nact.err = err
	return false
}

func newActivity(inst *Instagram) *Activity {
	act := &Activity{
		inst: inst,
	}
	return act
}

// Following allows to receive recent following activity.
//
// Use Next to paginate over activities
//
// See example: examples/activity/following.go
func (act *Activity) Following() *Activities {
	insta := act.inst

	nact := &Activities{inst: insta}
	nact.endpoint = urlActivityFollowing

	return nact
}

// Recent allows to receive recent activity (notifications).
//
// Use Activities.Next to paginate over activities.
//
// See example: examples/activity/recent.go
func (act *Activity) Recent() *Activities {
	insta := act.inst

	nact := &Activities{inst: insta}
	nact.endpoint = urlActivityRecent

	return nact
}
