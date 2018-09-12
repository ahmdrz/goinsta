package goinsta

import (
	"encoding/json"
	"fmt"
)

// Feed is the object for all feed endpoints.
type Feed struct {
	inst *Instagram
}

// newFeed creates new Feed structure
func newFeed(inst *Instagram) *Feed {
	return &Feed{
		inst: inst,
	}
}

// Feed search by locationID
func (feed *Feed) LocationID(locationID int64) (*FeedLocation, error) {
	insta := feed.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlFeedLocationID, locationID),
			Query: map[string]string{
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
			},
		},
	)
	if err != nil {
		return nil, err
	}

	res := &FeedLocation{}
	err = json.Unmarshal(body, res)
	return res, err
}

// FeedLocation is the struct that fits the structure returned by instagram on LocationID search.
type FeedLocation struct {
	RankedItems         []Item   `json:"ranked_items"`
	Items               []Item   `json:"items"`
	NumResults          int      `json:"num_results"`
	NextID              string   `json:"next_max_id"`
	MoreAvailable       bool     `json:"more_available"`
	AutoLoadMoreEnabled bool     `json:"auto_load_more_enabled"`
	MediaCount          int      `json:"media_count"`
	Location            Location `json:"location"`
	Status              string   `json:"status"`
}

// Tags search by Tag in user Feed
//
// (sorry for returning FeedTag. See #FeedTag)
func (feed *Feed) Tags(tag string) (*FeedTag, error) {
	insta := feed.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlFeedTag, tag),
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

// FeedTag is the struct that fits the structure returned by instagram on TagSearch.
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
