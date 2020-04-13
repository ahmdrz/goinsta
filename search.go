package goinsta

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Search is the object for all searches like Facebook, Location or Tag search.
type Search struct {
	inst *Instagram
}

// SearchResult handles the data for the results given by each type of Search.
type SearchResult struct {
	HasMore    bool   `json:"has_more"`
	RankToken  string `json:"rank_token"`
	Status     string `json:"status"`
	NumResults int64  `json:"num_results"`

	// User search results
	Users []User `json:"users"`

	// Tag search results
	Tags []struct {
		ID               int64       `json:"id"`
		Name             string      `json:"name"`
		MediaCount       int         `json:"media_count"`
		FollowStatus     interface{} `json:"follow_status"`
		Following        interface{} `json:"following"`
		AllowFollowing   interface{} `json:"allow_following"`
		AllowMutingStory interface{} `json:"allow_muting_story"`
		ProfilePicURL    interface{} `json:"profile_pic_url"`
		NonViolating     interface{} `json:"non_violating"`
		RelatedTags      interface{} `json:"related_tags"`
		DebugInfo        interface{} `json:"debug_info"`
	} `json:"results"`

	// Location search result
	RequestID string `json:"request_id"`
	Venues    []struct {
		ExternalIDSource string  `json:"external_id_source"`
		ExternalID       string  `json:"external_id"`
		Lat              float64 `json:"lat"`
		Lng              float64 `json:"lng"`
		Address          string  `json:"address"`
		Name             string  `json:"name"`
	} `json:"venues"`

	// Facebook
	// Facebook also uses `Users`
	Places   []interface{} `json:"places"`
	Hashtags []struct {
		Position int `json:"position"`
		Hashtag  struct {
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			MediaCount int    `json:"media_count"`
		} `json:"hashtag"`
	} `json:"hashtags"`
	ClearClientCache bool `json:"clear_client_cache"`
}

// newSearch creates new Search structure
func newSearch(inst *Instagram) *Search {
	search := &Search{
		inst: inst,
	}
	return search
}

// User search by username, you can use count optional parameter to get more than 50 items.
func (search *Search) User(user string, countParam ...int) (*SearchResult, error) {
	count := 50
	if len(countParam) > 0 {
		count = countParam[0]
	}
	insta := search.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchUser,
			Query: map[string]string{
				"ig_sig_key_version": goInstaSigKeyVersion,
				"is_typeahead":       "true",
				"q":                  user,
				"count":              fmt.Sprintf("%d", count),
				"rank_token":         insta.rankToken,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	for id := range res.Users {
		res.Users[id].inst = insta
	}
	return res, err
}

// Tags search by tag
func (search *Search) Tags(tag string) (*SearchResult, error) {
	insta := search.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchTag,
			Query: map[string]string{
				"is_typeahead": "true",
				"rank_token":   insta.rankToken,
				"q":            tag,
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	return res, err
}

// Location search by location.
// DEPRECATED - Instagram does not allow Location search method.
// Lat and Lng (Latitude & Longitude) cannot be ""
func (search *Search) Location(lat, lng, location string) (*SearchResult, error) {
	insta := search.inst
	q := map[string]string{
		"rank_token":     insta.rankToken,
		"latitude":       lat,
		"longitude":      lng,
		"ranked_content": "true",
	}

	if location != "" {
		q["search_query"] = location
	} else {
		q["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchLocation,
			Query:    q,
		},
	)
	if err != nil {
		return nil, err
	}

	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	return res, err
}

// Facebook search by facebook user.
func (search *Search) Facebook(user string) (*SearchResult, error) {
	insta := search.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchFacebook,
			Query: map[string]string{
				"query":      user,
				"rank_token": insta.rankToken,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	res := &SearchResult{}
	err = json.Unmarshal(body, res)
	return res, err
}
