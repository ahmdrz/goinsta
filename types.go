package goinsta

import (
	"fmt"
	"net/http"
)

// Instagram represent the main API handler
//
// ...
type Instagram struct {
	user string
	pass string
	// device id
	dID string
	// uuid
	uuid string
	// rankToken
	rankToken string
	// token
	token string
	// phone id
	pid string

	// Instagram objects

	// Users is the user interaction
	Profiles *Profiles
	// Account stores all personal data of the user and his/her options.
	Account *Account
	// Search performs searching of multiple things (users, locations...)
	Search *Search
	//
	Timeline *Timeline
	//
	Activity *Activity

	logged bool

	c *http.Client
}

// School is void structure (yet).
type School struct {
}

// PicURLInfo repre
type PicURLInfo struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type instaError struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

func instaToErr(ierr instaError) error {
	return fmt.Errorf("%s: %s (%s)", ierr.Status, ierr.Message, ierr.ErrorType)
}

type Nametag struct {
	Mode          int    `json:"mode"`
	Gradient      int    `json:"gradient"`
	Emoji         string `json:"emoji"`
	SelfieSticker int    `json:"selfie_sticker"`
}

type friendResp struct {
	Status     string     `json:"status"`
	Friendship Friendship `json:"friendship_status"`
}

type Friendship struct {
	IncomingRequest bool `json:"incoming_request"`
	FollowedBy      bool `json:"followed_by"`
	OutgoingRequest bool `json:"outgoing_request"`
	Following       bool `json:"following"`
	Blocking        bool `json:"blocking"`
	IsPrivate       bool `json:"is_private"`
}

// Images are different quality images
type Images struct {
	Versions []Candidate `json:"candidates"`
}

type Candidate struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

type Tag struct {
	In []struct {
		User                  User        `json:"user"`
		Position              []float64   `json:"position"`
		StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
		DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
	} `json:"in"`
}

// Caption is media caption
type Caption struct {
	ID              int64  `json:"pk"`
	UserID          int    `json:"user_id"`
	Text            string `json:"text"`
	Type            int    `json:"type"`
	CreatedAt       int    `json:"created_at"`
	CreatedAtUtc    int    `json:"created_at_utc"`
	ContentType     string `json:"content_type"`
	Status          string `json:"status"`
	BitFlags        int    `json:"bit_flags"`
	User            User   `json:"user"`
	DidReportAsSpam bool   `json:"did_report_as_spam"`
	MediaID         int64  `json:"media_id"`
	HasTranslation  bool   `json:"has_translation"`
}

type Mentions struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int     `json:"z"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	IsPinned int     `json:"is_pinned"`
	User     User    `json:"user"`
}

// Videos are different quality videos
type Videos struct {
	Type   int    `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
	ID     string `json:"id"`
}

type timeStoryResp struct {
	Status string       `json:"status"`
	Media  []StoryMedia `json:"tray"`
}
