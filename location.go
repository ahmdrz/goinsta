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
	inst *Instagram
	err error
	endpoint string

	Sections      []LayoutSection `json:"sections"`
	MoreAvailable bool            `json:"more_available"`
	NextPage      int             `json:"next_page"`
	NextMediaIds  []int64         `json:"next_media_ids"`
	NextMaxID     string          `json:"next_max_id"`
	Status        string          `json:"status"`
}

// Next allows to paginate after calling:
// Account.Follow* and User.Follow*
//
// New user list is stored inside Users
//
// returns false when list reach the end.
func (section *Section) Next() bool {
	if section.err != nil {
		return false
	}

	insta := section.inst
	endpoint := section.endpoint

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id": section.NextMaxID,
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
				"_csrftoken":     insta.token,
				"_uuid":          insta.uuid,
			},
			IsPost: true,
		},
	)
	if err == nil {
		newSection := Section{}
		err = json.Unmarshal(body, &newSection)

		if err != nil {
			section.err = err
			return false
		}

		newSection.inst = section.inst
		newSection.endpoint = section.endpoint
		*section = newSection

		return true
	}

	section.err = err
	return false
}

func (l *LocationInstance) Feeds(locationID int64) *Section {
	endpoint := fmt.Sprintf(urlFeedLocations, locationID)
	return &Section{inst:l.inst, endpoint:endpoint}
}
