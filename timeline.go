package goinsta

import (
	"encoding/json"
)

// Timeline is the object to represent the main feed on instagram, the first page that shows the latest feeds of my following contacts.
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
func (time *Timeline) Get() *FeedMedia {
	insta := time.inst
	media := &FeedMedia{}
	media.inst = insta
	media.endpoint = urlTimeline
	return media
}

// Stories returns slice of StoryMedia
func (time *Timeline) Stories() (*Tray, error) {
	body, err := time.inst.sendSimpleRequest(urlStories)
	if err == nil {
		tray := &Tray{}
		err = json.Unmarshal(body, tray)
		if err != nil {
			return nil, err
		}
		tray.set(time.inst, urlStories)
		return tray, nil
	}
	return nil, err
}
