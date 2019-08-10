package goinsta

import (
	"encoding/json"
	"fmt"
)

// Hashtag is used for getting the media that matches a hashtag on instagram.
type Hashtag struct {
	inst *Instagram
	err  error

	Name string `json:"name"`

	Sections []struct {
		LayoutType    string `json:"layout_type"`
		LayoutContent struct {
			// F*ck you instagram.
			// Why you do this f*cking horribly structure?!?
			// Media []Media IS EASY. CHECK IT!
			Medias []struct {
				Item Item `json:"media"`
			} `json:"medias"`
		} `json:"layout_content"`
		FeedType        string `json:"feed_type"`
		ExploreItemInfo struct {
			NumColumns      int     `json:"num_columns"`
			TotalNumColumns int     `json:"total_num_columns"`
			AspectRatio     float32 `json:"aspect_ratio"`
			Autoplay        bool    `json:"autoplay"`
		} `json:"explore_item_info"`
	} `json:"sections"`
	MediaCount          int     `json:"media_count"`
	ID                  int64   `json:"id"`
	MoreAvailable       bool    `json:"more_available"`
	NextID              string  `json:"next_max_id"`
	NextPage            int     `json:"next_page"`
	NextMediaIds        []int64 `json:"next_media_ids"`
	AutoLoadMoreEnabled bool    `json:"auto_load_more_enabled"`
	Status              string  `json:"status"`
}

func (h *Hashtag) setValues() {
	for i := range h.Sections {
		for j := range h.Sections[i].LayoutContent.Medias {
			m := &FeedMedia{
				inst: h.inst,
			}
			setToItem(&h.Sections[i].LayoutContent.Medias[j].Item, m)
		}
	}
}

// NewHashtag returns initialised hashtag structure
// Name parameter is hashtag name
func (inst *Instagram) NewHashtag(name string) *Hashtag {
	return &Hashtag{
		inst: inst,
		Name: name,
	}
}

// Sync updates Hashtag information preparing it to Next call.
func (h *Hashtag) Sync() error {
	insta := h.inst

	body, err := insta.sendSimpleRequest(urlTagSync, h.Name)
	if err == nil {
		var resp struct {
			Name       string `json:"name"`
			ID         int64  `json:"id"`
			MediaCount int    `json:"media_count"`
		}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			h.Name = resp.Name
			h.ID = resp.ID
			h.MediaCount = resp.MediaCount
			h.setValues()
		}
	}
	return err
}

// Next paginates over hashtag pages (xd).
func (h *Hashtag) Next() bool {
	if h.err != nil {
		return false
	}
	insta := h.inst
	name := h.Name
	body, err := insta.sendRequest(
		&reqOptions{
			Query: map[string]string{
				"max_id":     h.NextID,
				"rank_token": insta.rankToken,
				"page":       fmt.Sprintf("%d", h.NextPage),
			},
			Endpoint: fmt.Sprintf(urlTagContent, name),
			IsPost:   false,
		},
	)
	if err == nil {
		ht := &Hashtag{}
		err = json.Unmarshal(body, ht)
		if err == nil {
			*h = *ht
			h.inst = insta
			h.Name = name
			if !h.MoreAvailable {
				h.err = ErrNoMore
			}
			h.setValues()
			return true
		}
	}
	h.err = err
	return false
}

// Error returns hashtag error
func (h *Hashtag) Error() error {
	return h.err
}

// Stories returns hashtag stories.
func (h *Hashtag) Stories() (*StoryMedia, error) {
	body, err := h.inst.sendSimpleRequest(
		urlTagStories, h.Name,
	)
	if err == nil {
		var resp struct {
			Story  StoryMedia `json:"story"`
			Status string     `json:"status"`
		}
		err = json.Unmarshal(body, &resp)
		return &resp.Story, err
	}
	return nil, err
}
