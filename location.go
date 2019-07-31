package goinsta

import (
	"encoding/json"
	"fmt"
)

type LocationInstance struct {
	inst *Instagram
}

func newLocation(inst *Instagram) *LocationInstance {
	return &LocationInstance{inst: inst}
}

type LayoutSection struct {
	LayoutType    string `json:"layout_type"`
	LayoutContent struct {
		Medias []struct {
			Media Item `json:"media"`
		} `json:"medias"`
	} `json:"layout_content"`
	FeedType        string `json:"feed_type"`
	ExploreItemInfo struct {
		NumColumns      int  `json:"num_columns"`
		TotalNumColumns int  `json:"total_num_columns"`
		AspectRatio     int  `json:"aspect_ratio"`
		Autoplay        bool `json:"autoplay"`
	} `json:"explore_item_info"`
}

type Section struct {
	Sections      []LayoutSection `json:"sections"`
	MoreAvailable bool            `json:"more_available"`
	NextPage      int             `json:"next_page"`
	NextMediaIds  []int64         `json:"next_media_ids"`
	NextMaxID     string          `json:"next_max_id"`
	Status        string          `json:"status"`
}

func (l *LocationInstance) Feeds(locationID int64) (*Section, error) {
	// TODO: use pagination for location feeds.
	insta := l.inst
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlFeedLocations, locationID),
			Query: map[string]string{
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
				"_csrftoken":     insta.token,
				"_uuid":          insta.uuid,
			},
			IsPost: true,
		},
	)
	if err != nil {
		return nil, err
	}

	section := &Section{}
	err = json.Unmarshal(body, section)
	return section, err
}
