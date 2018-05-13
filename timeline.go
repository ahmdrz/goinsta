package goinsta

import (
	"encoding/json"
)

type Timeline struct {
	inst *Instagram
}

func newTimeline(inst *Instagram) *Timeline {
	time := &Timeline{
		inst: inst,
	}
	return time
}

// Get returns latest media from timeline.
//
// For pagination use FeedMedia.Next()
func (time *Timeline) Get() (*FeedMedia, error) {
	insta := time.inst

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlTimeline,
			Query: map[string]string{
				"max_id":         "",
				"rank_token":     insta.rankToken,
				"ranked_content": "true",
			},
		},
	)
	if err == nil {
		media := &FeedMedia{}
		err = json.Unmarshal(body, media)
		media.inst = insta
		media.endpoint = urlTimeline
		return media, err
	}
	return nil, err
}

func (time *Timeline) Stories() ([]StoryMedia, error) {
	body, err := time.inst.sendSimpleRequest(urlStories)
	if err == nil {
		resp := &timeStoryResp{}
		err = json.Unmarshal(body, &resp)
		if err != nil {
			return nil, err
		}
		for i := range resp.Media {
			resp.Media[i].inst = time.inst
			resp.Media[i].endpoint = urlStories
		}
		return resp.Media, nil
	}
	return nil, err
}
