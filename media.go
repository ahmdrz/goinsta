package goinsta

import (
	"encoding/json"
	"fmt"
	"io"
	neturl "net/url"
	"os"
	"path"
	"strconv"
	"strings"
)

type mediaType byte

const (
	feedItem = iota
	storyItem
)

// Item represents media items
//
// All Item has Images or Videos objects which contains the url(s).
// You can use Download function to get the best quality Image or Video from Item.
type Item struct {
	inst     *Instagram `json:"-"`
	father   mediaType  `json:"-"`
	Comments *Comments  `json:"-"`

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
	Likers                       []User      `json:"likers"`
	CommentLikesEnabled          bool        `json:"comment_likes_enabled"`
	CommentThreadingEnabled      bool        `json:"comment_threading_enabled"`
	HasMoreComments              bool        `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int         `json:"max_num_visible_preview_comments"`
	// _PreviewComments can be `string` or `[]string` or `[]Comment`.
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

// Comments allows user to interact with media (item) comments.
// You can Add or Delete by index or by user name.
type Comments struct {
	inst *Instagram

	media    Media
	endpoint string
	err      error

	Items                          []Comment `json:"comments"`
	CommentCount                   int       `json:"comment_count"`
	Caption                        Caption   `json:"caption"`
	CaptionIsEdited                bool      `json:"caption_is_edited"`
	HasMoreComments                bool      `json:"has_more_comments"`
	HasMoreHeadloadComments        bool      `json:"has_more_headload_comments"`
	MediaHeaderDisplay             string    `json:"media_header_display"`
	DisplayRealtimeTypingIndicator bool      `json:"display_realtime_typing_indicator"`
	NextID                         string    `json:"next_max_id"`
	LastID                         string    `json:"next_min_id"`
	Status                         string    `json:"status"`

	//PreviewComments                []Comment   `json:"preview_comments"`
}

func newComments(media Media) *Comments {
	c := &Comments{
		media: media,
	}
	return c
}

func (comments Comments) Error() error {
	return comments.err
}

// Next allows comment pagination.
//
// This function support concurrency methods to get comments using Last and Next ID
//
// New comments are stored in Comments.Items
func (comments *Comments) Next() bool {
	if comments.err != nil {
		return false
	}

	insta := comments.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"can_support_threading": true,
			"max_id":                comments.NextID,
			"min_id":                comments.LastID,
		},
	)
	if err != nil {
		comments.err = err
		return false
	}

	media := comments.media
	endpoint := comments.endpoint

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		c := Comments{}
		err = json.Unmarshal(body, &c)
		if err == nil {
			*comments = c
			comments.endpoint = endpoint
			comments.media = media
			if !comments.HasMoreComments || comments.NextID == "" {
				comments.err = ErrNoMore
			}
			return true
		}
	}
	comments.err = err
	return false
}

// Sync prepare Comments to receive comments.
// Use Next to receive comments.
func (comments *Comments) Sync() {
	media := comments.media
	endpoint := fmt.Sprintf(urlCommentSync, comments.media.ID())
	comments.media = media
	comments.endpoint = endpoint
	return
}

// Add push a comment in media.
//
// If parent media is a Story this function will send a private message
// replying the Instagram story. TODO
func (comments *Comments) Add(msg string) error {
	insta := comments.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"comment_text": msg,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlCommentAdd, comments.media.ID()),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// Del deletes comment.
func (comments *Comments) Del(comment *Comment) error {
	insta := comments.media.instagram()

	data, err := insta.prepareData()
	if err != nil {
		return err
	}
	id := comment.getid()

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlCommentDelete, comments.media.ID(), id),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// DelByID removes comment using id.
//
// See example: examples/media/commentsDelByID.go
func (comments *Comments) DelByID(id string) error {
	return comments.Del(&Comment{idstr: id})
}

// DelMine removes all of your comments limited by parsed parameter.
//
// If limit is <= 0 DelMine will delete all your comments.
//
// See example: examples/media/commentsDelMine.go
func (comments *Comments) DelMine(limit int) error {
	i := 0
	if limit <= 0 {
		i = limit - 1
	}
	comments.Sync()

	insta := comments.media.instagram()
floop:
	for comments.Next() {
		for _, c := range comments.Items {
			if c.UserID == insta.Account.ID || c.User.ID == insta.Account.ID {
				if i >= limit {
					break floop
				}
				comments.Del(&c)
				i++
			}
		}
	}
	if err := comments.Error(); err != nil && err != ErrNoMore {
		return err
	}
	return nil
}

func setToItem(item *Item, t mediaType, media Media) {
	item.father = t
	item.inst = media.instagram()
	item.Comments = newComments(media)
}

func getname(name string) string {
	nname := name
	i := 1
	for {
		ext := path.Ext(name)

		_, err := os.Stat(name)
		if err != nil {
			break
		}
		if ext != "" {
			nname = strings.Replace(nname, ext, "", -1)
		}
		name = fmt.Sprintf("%s.%d%s", nname, i, ext)
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

type bestMedia struct {
	w, h int
	url  string
}

func getBest(obj interface{}) []string {
	m := make(map[string]bestMedia)

	switch t := obj.(type) {
	// getting best video
	case []Video:
		for _, video := range t {
			v, ok := m[video.ID]
			if !ok {
				m[video.ID] = bestMedia{
					w:   video.Width,
					h:   video.Height,
					url: video.URL,
				}
			} else {
				if v.w < video.Width && video.Height > v.h {
					m[video.ID] = bestMedia{
						w:   video.Width,
						h:   video.Height,
						url: video.URL,
					}
				}
			}
		}
		// getting best image
	case []Candidate:
		for _, image := range t {
			url, err := neturl.Parse(image.URL)
			if err != nil {
				continue
			}

			base := path.Base(url.Path)
			i, ok := m[base]
			if !ok {
				m[base] = bestMedia{
					w:   image.Width,
					h:   image.Height,
					url: image.URL,
				}
			} else {
				if i.w < image.Width && image.Height > i.h {
					m[base] = bestMedia{
						w:   image.Width,
						h:   image.Height,
						url: image.URL,
					}
				}
			}
		}
	}
	s := []string{}
	// getting best to return in string slice
	for _, v := range m {
		s = append(s, v.url)
	}
	m = nil
	return s
}

// Download downloads media item (video or image) with the best quality.
//
// Input parameters are folder and filename. If filename is "" will be saved with
// the default value name.
//
// If file exists it will be saved
func (item *Item) Download(inst *Instagram, folder, name string) error {
	os.MkdirAll(folder, 0777)
	for _, url := range getBest(item.Images.Versions) {
		nname := name
		if nname == "" {
			u, err := neturl.Parse(url)
			if err != nil {
				return err
			}

			nname = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, path.Base(u.Path))
		} else {
			nname = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, nname)
		}
		nname = getname(nname)

		err := download(inst, url, nname)
		if err != nil {
			return err
		}
	}

	for _, url := range getBest(item.Videos) {
		nname := name
		if nname == "" {
			u, err := neturl.Parse(url)
			if err != nil {
				return err
			}

			nname = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, path.Base(u.Path))
		} else {
			nname = fmt.Sprintf("%s%c%s", folder, os.PathSeparator, nname)
		}
		nname = getname(nname)

		err := download(inst, url, nname)
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

// PreviewComments returns string slice or single string (inside Comment slice)
// Depending on PreviewComments parameter.
// If PreviewComments are string or []string only the Text field will be filled.
func (item *Item) PreviewComments() []Comment {
	switch s := item._PreviewComments.(type) {
	case []Comment:
		return s
	case []string:
		comments := make([]Comment, 0)
		for i := range s {
			comments = append(comments, Comment{
				Text: s[i],
			})
		}
		return comments
	case string:
		comments := []Comment{
			Comment{
				Text: s,
			},
		}
		return comments
	}
	return nil
}

type Media interface {
	// Next allows pagination
	Next() bool
	// Error returns error (in case it have been occurred)
	Error() error
	// ID returns media id
	ID() string

	instagram() *Instagram
}

type StoryMedia struct {
	inst     *Instagram
	endpoint string
	uid      int64

	err error

	Pk              int64 `json:"id"`
	LatestReelMedia int   `json:"latest_reel_media"`
	ExpiringAt      int   `json:"expiring_at"`
	//Seen            float64  `json:"seen"`
	CanReply        bool     `json:"can_reply"`
	CanReshare      bool     `json:"can_reshare"`
	ReelType        string   `json:"reel_type"`
	User            User     `json:"user"`
	Items           []Item   `json:"items"`
	ReelMentions    []string `json:"reel_mentions"`
	PrefetchCount   int      `json:"prefetch_count"`
	HasBestiesMedia int      `json:"has_besties_media"`
	Status          string   `json:"status"`
}

// ID returns Story id
func (media *StoryMedia) ID() string {
	return strconv.FormatInt(media.Pk, 10)
}

func (media *StoryMedia) instagram() *Instagram {
	return media.inst
}

func (media *StoryMedia) setValues() {
	for i := range media.Items {
		setToItem(&media.Items[i], storyItem, media)
	}
}

// Error returns error happend any error
func (media StoryMedia) Error() error {
	return media.err
}

// Seen marks story as seen.
func (media *StoryMedia) Seen() error {
	// TODO
	return nil
}

// Next allows to paginate after calling:
// User.Stories
//
// returns false when list reach the end
// if StoryMedia.Error() is ErrNoMore no problem have been occurred.
func (media *StoryMedia) Next() bool {
	if media.err != nil {
		return false
	}

	insta := media.inst
	endpoint := media.endpoint
	if media.uid != 0 {
		endpoint = fmt.Sprintf(endpoint, media.uid)
	}

	body, err := insta.sendSimpleRequest(endpoint)
	if err == nil {
		m := StoryMedia{}
		err = json.Unmarshal(body, &m)
		if err == nil {
			// TODO check NextID media
			*media = m
			media.inst = insta
			media.endpoint = endpoint
			media.err = ErrNoMore
			media.setValues()
			return true
		}
	}
	media.err = err
	return false
}

// Media represent a set of media items
type FeedMedia struct {
	inst *Instagram

	err error

	uid       int64
	endpoint  string
	timestamp string

	Items               []Item `json:"items"`
	NumResults          int    `json:"num_results"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Status              string `json:"status"`
	// Can be int64 and string
	// this is why we recomend Next() usage :')
	NextID interface{} `json:"next_max_id"`
}

func (media *FeedMedia) instagram() *Instagram {
	return media.inst
}

// AcquireFeed returns initilised FeedMedia
//
// Use FeedMedia.Sync() to update FeedMedia information.
func AcquireFeed(inst *Instagram) *FeedMedia {
	return &FeedMedia{inst: inst}
}

// SetID sets media ID
// this value can be int64 or string
func (media *FeedMedia) SetID(id interface{}) {
	media.NextID = id
}

// Sync updates media values.
func (media *FeedMedia) Sync() error {
	id := media.ID()
	insta := media.inst

	data, err := insta.prepareData(
		map[string]interface{}{
			"media_id": id,
		},
	)
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlMediaInfo, id),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err != nil {
		return err
	}

	m := FeedMedia{}
	err = json.Unmarshal(body, &m)
	*media = m
	media.endpoint = urlMediaInfo
	media.inst = insta
	media.NextID = id
	media.setValues()
	return err
}

func (media *FeedMedia) setValues() {
	for i := range media.Items {
		setToItem(&media.Items[i], feedItem, media)
	}
}

func (media FeedMedia) Error() error {
	return media.err
}

// ID returns media id.
func (media *FeedMedia) ID() string {
	switch s := media.NextID.(type) {
	case string:
		return s
	case int64:
		return strconv.FormatInt(s, 10)
	}
	return ""
}

// Next allows to paginate after calling:
// User.Feed
//
// returns false when list reach the end.
// if FeedMedia.Error() is ErrNoMore no problem have been occurred.
func (media *FeedMedia) Next() bool {
	if media.err != nil {
		return false
	}

	insta := media.inst
	endpoint := media.endpoint
	next := media.ID()

	if media.uid != 0 {
		endpoint = fmt.Sprintf(endpoint, media.uid)
	}

	body, err := insta.sendRequest(
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
				media.err = ErrNoMore
			}
			media.setValues()
			return true
		}
	}
	return false
}
