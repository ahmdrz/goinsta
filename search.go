package goinsta

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Search struct {
	inst *Instagram
}

type SearchResult struct {
	HasMore    bool   `json:"has_more"`
	RankToken  string `json:"rank_token"`
	Status     string `json:"status"`
	NumResults int    `json:"num_results"`

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

// User search by username
func (search *Search) User(user string) (*SearchResult, error) {
	insta := search.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlSearchUser,
			Query: map[string]string{
				"ig_sig_key_version": goInstaSigKeyVersion,
				"is_typeahead":       "true",
				"query":              user,
				"rank_token":         insta.rankToken,
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

// FeedTags search by Tag in user Feed
//
// (sorry for returning FeedTag. See #FeedTag)
func (search *Search) FeedTags(tag string) (*FeedTag, error) {
	insta := search.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlSearchFeedTag, tag),
			Query: map[string]string{
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
			},
		},
	)
	if err != nil {
		return nil, err
	}
	res := &FeedTag{}
	err = json.Unmarshal(body, res)
	return res, err
}

// Instagram's database is f*cking shit.
// We all hate nodejs (seems that they uses nodejs and mongoldb)
// I don't know why FeedTags returns this aberration structure.
type FeedTag struct {
	RankedItems         []Item     `json:"ranked_items"`
	Images              []Item     `json:"items"`
	NumResults          int        `json:"num_results"`
	NextID              string     `json:"next_max_id"`
	MoreAvailable       bool       `json:"more_available"`
	AutoLoadMoreEnabled bool       `json:"auto_load_more_enabled"`
	Story               StoryMedia `json:"story"`
	Status              string     `json:"status"`
}
