package goinsta

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
)

// Item represents media items
//
// All Item has Images or Videos objects which contains the url(s).
// You can use Download function to get the best quality Image or Video from Item.
type Item struct {
	TakenAt          int     `json:"taken_at"`
	ID               int64   `json:"pk"`
	IDStr            string  `json:"id"`
	DeviceTimestamp  int64   `json:"device_timestamp"`
	MediaType        int     `json:"media_type"`
	Code             string  `json:"code"`
	ClientCacheKey   string  `json:"client_cache_key"`
	FilterType       int     `json:"filter_type"`
	CarouselParentID string  `json:"carousel_parent_id"`
	CarouselMedia    []Item  `json:"carousel_media,omitempty"`
	User             User    `json:"user"`
	CanViewerReshare bool    `json:"can_viewer_reshare"`
	Caption          Caption `json:"caption"`
	CaptionIsEdited  bool    `json:"caption_is_edited"`
	Likes            int     `json:"like_count"`
	HasLiked         bool    `json:"has_liked"`
	// _TopLikers can be `string` or `[]string`.
	// Use TopLikers function instead of getting it directly.
	_TopLikers                   interface{} `json:"top_likers"`
	CommentLikesEnabled          bool        `json:"comment_likes_enabled"`
	CommentThreadingEnabled      bool        `json:"comment_threading_enabled"`
	HasMoreComments              bool        `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int         `json:"max_num_visible_preview_comments"`
	// _PreviewComments can be `string` or `[]string`.
	// Use PreviewComments function instead of getting it directly.
	_PreviewComments     interface{} `json:"preview_comments,omitempty"`
	CommentCount         int         `json:"comment_count"`
	PhotoOfYou           bool        `json:"photo_of_you"`
	Usertags             Tag         `json:"usertags,omitempty"`
	FbUserTags           Tag         `json:"fb_user_tags"`
	CanViewerSave        bool        `json:"can_viewer_save"`
	OrganicTrackingToken string      `json:"organic_tracking_token"`
	// Images contains URL images in different versions.
	// Version = quality.
	Images          Images `json:"image_versions2,omitempty"`
	OriginalWidth   int    `json:"original_width,omitempty"`
	OriginalHeight  int    `json:"original_height,omitempty"`
	ImportedTakenAt int    `json:"imported_taken_at,omitempty"`
	Location        struct {
		Pk               int     `json:"pk"`
		Name             string  `json:"name"`
		Address          string  `json:"address"`
		City             string  `json:"city"`
		ShortName        string  `json:"short_name"`
		Lng              float64 `json:"lng"`
		Lat              float64 `json:"lat"`
		ExternalSource   string  `json:"external_source"`
		FacebookPlacesID int64   `json:"facebook_places_id"`
	} `json:"location,omitempty"`
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`

	// Only for stories
	StoryEvents              []interface{} `json:"story_events"`
	StoryHashtags            []interface{} `json:"story_hashtags"`
	StoryPolls               []interface{} `json:"story_polls"`
	StoryFeedMedia           []interface{} `json:"story_feed_media"`
	StorySoundOn             []interface{} `json:"story_sound_on"`
	CreativeConfig           interface{}   `json:"creative_config"`
	StoryLocations           []interface{} `json:"story_locations"`
	StorySliders             []interface{} `json:"story_sliders"`
	StoryQuestions           []interface{} `json:"story_questions"`
	StoryProductItems        []interface{} `json:"story_product_items"`
	SupportsReelReactions    bool          `json:"supports_reel_reactions"`
	ShowOneTapFbShareTooltip bool          `json:"show_one_tap_fb_share_tooltip"`
	HasSharedToFb            int           `json:"has_shared_to_fb"`
	Mentions                 []Mentions
	Videos                   []Video `json:"video_versions,omitempty"`
	HasAudio                 bool    `json:"has_audio,omitempty"`
	VideoDuration            float64 `json:"video_duration,omitempty"`
	IsDashEligible           int     `json:"is_dash_eligible,omitempty"`
	VideoDashManifest        string  `json:"video_dash_manifest,omitempty"`
	NumberOfQualities        int     `json:"number_of_qualities,omitempty"`
}

func getname(name string) string {
	nname := name
	i := 0
	for {
		_, err := os.Stat(name)
		if err != nil {
			break
		}
		name = fmt.Sprintf("%s.%d", nname, i)
		i++
	}
	return name
}

func download(inst *Instagram, url, dst string) error {
	file, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := inst.c.Get(url)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	return err
}

// Download downloads media item (video or image) with the best quality.
//
// Input parameters are folder and filename. If filename is "" will be saved with
// the default value name.
func (item *Item) Download(inst *Instagram, folder, name string) error {
	for _, c := range item.Images.Versions {
		if name == "" {
			name = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, path.Base(c.URL))
		} else {
			name = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, name)
		}
		name = getname(name)

		err := download(inst, c.URL, name)
		if err != nil {
			return err
		}
	}

	for _, video := range item.Videos {
		if name == "" {
			name = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, path.Base(video.URL))
		} else {
			name = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, name)
		}
		name = getname(name)

		err := download(inst, video.URL, name)
		if err != nil {
			return err
		}
	}
	return nil
}

// TopLikers returns string slice or single string (inside string slice)
// Depending on TopLikers parameter.
func (item *Item) TopLikers() []string {
	switch s := item._TopLikers.(type) {
	case string:
		return []string{s}
	case []string:
		return s
	}
	return nil
}

// PreviewComments returns string slice or single string (inside string slice)
// Depending on PreviewComments parameter.
func (item *Item) PreviewComments() []string {
	switch s := item._PreviewComments.(type) {
	case string:
		return []string{s}
	case []string:
		return s
	}
	return nil
}

type Media interface {
	Next() error
}

type StoryMedia struct {
	inst     *Instagram
	endpoint string
	uid      int64

	ID              int      `json:"id"`
	LatestReelMedia int      `json:"latest_reel_media"`
	ExpiringAt      int      `json:"expiring_at"`
	Seen            float64  `json:"seen"`
	CanReply        bool     `json:"can_reply"`
	CanReshare      bool     `json:"can_reshare"`
	ReelType        string   `json:"reel_type"`
	User            User     `json:"user"`
	Items           []Item   `json:"items"`
	ReelMentions    []string `json:"reel_mentions"`
	PrefetchCount   int      `json:"prefetch_count"`
	HasBestiesMedia bool     `json:"has_besties_media"`
	Status          string   `json:"status"`
}

// Next allows to paginate after calling:
// User.Stories
func (media *StoryMedia) Next() (err error) {
	var body []byte
	insta := media.inst
	endpoint := media.endpoint
	if media.uid != 0 {
		endpoint = fmt.Sprintf(endpoint, media.uid)
	}

	body, err = insta.sendSimpleRequest(endpoint)
	if err == nil {
		m := StoryMedia{}
		err = json.Unmarshal(body, &m)
		if err == nil {
			// TODO check NextID media
			err = ErrNoMore
			*media = m
			media.inst = insta
			media.endpoint = endpoint
		}
	}
	return err
}

// Media represent a set of media items
type FeedMedia struct {
	inst *Instagram

	uid       int64
	endpoint  string
	timestamp string

	Items               []Item `json:"items"`
	NumResults          int    `json:"num_results"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Status              string `json:"status"`
	// Can be int64 and string
	// this is why recomend Next() usage :')
	NextID interface{} `json:"next_max_id"`
}

// Next allows to paginate after calling:
// User.Feed
//
// returns ErrNoMore when list reach the end.
func (media *FeedMedia) Next() (err error) {
	var body []byte
	insta := media.inst
	endpoint := media.endpoint
	next := ""

	switch s := media.NextID.(type) {
	case string:
		next = s
	case int64:
		next = strconv.FormatInt(s, 10)
	}

	if media.uid != 0 {
		endpoint = fmt.Sprintf(endpoint, media.uid)
	}

	body, err = insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id":         next,
				"rank_token":     insta.rankToken,
				"min_timestamp":  media.timestamp,
				"ranked_content": "true",
			},
		},
	)
	if err == nil {
		m := FeedMedia{}
		err = json.Unmarshal(body, &m)
		if err == nil {
			*media = m
			media.inst = insta
			media.endpoint = endpoint
			if m.NextID == 0 || !m.MoreAvailable {
				err = ErrNoMore
			}
		}
	}
	return err
}
