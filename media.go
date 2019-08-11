package goinsta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	neturl "net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// StoryReelMention represent story reel mention
type StoryReelMention struct {
	X        float64 `json:"x"`
	Y        float64 `json:"y"`
	Z        int     `json:"z"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	IsPinned int     `json:"is_pinned"`
	IsHidden int     `json:"is_hidden"`
	User     User
}

// StoryCTA represent story cta
type StoryCTA struct {
	Links []struct {
		LinkType                                int         `json:"linkType"`
		WebURI                                  string      `json:"webUri"`
		AndroidClass                            string      `json:"androidClass"`
		Package                                 string      `json:"package"`
		DeeplinkURI                             string      `json:"deeplinkUri"`
		CallToActionTitle                       string      `json:"callToActionTitle"`
		RedirectURI                             interface{} `json:"redirectUri"`
		LeadGenFormID                           string      `json:"leadGenFormId"`
		IgUserID                                string      `json:"igUserId"`
		AppInstallObjectiveInvalidationBehavior interface{} `json:"appInstallObjectiveInvalidationBehavior"`
	} `json:"links"`
}

// Item represents media items
//
// All Item has Images or Videos objects which contains the url(s).
// You can use Download function to get the best quality Image or Video from Item.
type Item struct {
	media    Media
	Comments *Comments `json:"-"`

	TakenAt          int64   `json:"taken_at"`
	Pk               int64   `json:"pk"`
	ID               string  `json:"id"`
	CommentsDisabled bool    `json:"comments_disabled"`
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
	// Toplikers can be `string` or `[]string`.
	// Use TopLikers function instead of getting it directly.
	Toplikers                    interface{} `json:"top_likers"`
	Likers                       []User      `json:"likers"`
	CommentLikesEnabled          bool        `json:"comment_likes_enabled"`
	CommentThreadingEnabled      bool        `json:"comment_threading_enabled"`
	HasMoreComments              bool        `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int         `json:"max_num_visible_preview_comments"`
	// Previewcomments can be `string` or `[]string` or `[]Comment`.
	// Use PreviewComments function instead of getting it directly.
	Previewcomments interface{} `json:"preview_comments,omitempty"`
	CommentCount    int         `json:"comment_count"`
	PhotoOfYou      bool        `json:"photo_of_you"`
	// Tags are tagged people in photo
	Tags struct {
		In []Tag `json:"in"`
	} `json:"usertags,omitempty"`
	FbUserTags           Tag    `json:"fb_user_tags"`
	CanViewerSave        bool   `json:"can_viewer_save"`
	OrganicTrackingToken string `json:"organic_tracking_token"`
	// Images contains URL images in different versions.
	// Version = quality.
	Images          Images   `json:"image_versions2,omitempty"`
	OriginalWidth   int      `json:"original_width,omitempty"`
	OriginalHeight  int      `json:"original_height,omitempty"`
	ImportedTakenAt int64    `json:"imported_taken_at,omitempty"`
	Location        Location `json:"location,omitempty"`
	Lat             float64  `json:"lat,omitempty"`
	Lng             float64  `json:"lng,omitempty"`

	// Videos
	Videos            []Video `json:"video_versions,omitempty"`
	HasAudio          bool    `json:"has_audio,omitempty"`
	VideoDuration     float64 `json:"video_duration,omitempty"`
	ViewCount         float64 `json:"view_count,omitempty"`
	IsDashEligible    int     `json:"is_dash_eligible,omitempty"`
	VideoDashManifest string  `json:"video_dash_manifest,omitempty"`
	NumberOfQualities int     `json:"number_of_qualities,omitempty"`

	// Only for stories
	StoryEvents              []interface{}      `json:"story_events"`
	StoryHashtags            []interface{}      `json:"story_hashtags"`
	StoryPolls               []interface{}      `json:"story_polls"`
	StoryFeedMedia           []interface{}      `json:"story_feed_media"`
	StorySoundOn             []interface{}      `json:"story_sound_on"`
	CreativeConfig           interface{}        `json:"creative_config"`
	StoryLocations           []interface{}      `json:"story_locations"`
	StorySliders             []interface{}      `json:"story_sliders"`
	StoryQuestions           []interface{}      `json:"story_questions"`
	StoryProductItems        []interface{}      `json:"story_product_items"`
	StoryCTA                 []StoryCTA         `json:"story_cta"`
	ReelMentions             []StoryReelMention `json:"reel_mentions"`
	SupportsReelReactions    bool               `json:"supports_reel_reactions"`
	ShowOneTapFbShareTooltip bool               `json:"show_one_tap_fb_share_tooltip"`
	HasSharedToFb            int64              `json:"has_shared_to_fb"`
	Mentions                 []Mentions
	Audience                 string `json:"audience,omitempty"`
}

// MediaToString returns Item.MediaType as string.
func (item *Item) MediaToString() string {
	switch item.MediaType {
	case 1:
		return "photo"
	case 2:
		return "video"
	}
	return ""
}

func setToItem(item *Item, media Media) {
	item.media = media
	item.User.inst = media.instagram()
	item.Comments = newComments(item)
	for i := range item.CarouselMedia {
		item.CarouselMedia[i].User = item.User
		setToItem(&item.CarouselMedia[i], media)
	}
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

func download(inst *Instagram, url, dst string) (string, error) {
	file, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer file.Close()

	resp, err := inst.c.Get(url)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(file, resp.Body)
	return dst, err
}

type bestMedia struct {
	w, h int
	url  string
}

// GetBest returns best quality image or video.
//
// Arguments can be []Video or []Candidate
func GetBest(obj interface{}) string {
	m := bestMedia{}

	switch t := obj.(type) {
	// getting best video
	case []Video:
		for _, video := range t {
			if m.w < video.Width && video.Height > m.h && video.URL != "" {
				m.w = video.Width
				m.h = video.Height
				m.url = video.URL
			}
		}
		// getting best image
	case []Candidate:
		for _, image := range t {
			if m.w < image.Width && image.Height > m.h && image.URL != "" {
				m.w = image.Width
				m.h = image.Height
				m.url = image.URL
			}
		}
	}
	return m.url
}

var rxpTags = regexp.MustCompile(`#\w+`)

// Hashtags returns caption hashtags.
//
// Item media parent must be FeedMedia.
//
// See example: examples/media/hashtags.go
func (item *Item) Hashtags() []Hashtag {
	tags := rxpTags.FindAllString(item.Caption.Text, -1)

	hsh := make([]Hashtag, len(tags))

	i := 0
	for _, tag := range tags {
		hsh[i].Name = tag[1:]
		i++
	}

	for _, comment := range item.PreviewComments() {
		tags := rxpTags.FindAllString(comment.Text, -1)

		for _, tag := range tags {
			hsh = append(hsh, Hashtag{Name: tag[1:]})
		}
	}

	return hsh
}

// Delete deletes your media item. StoryMedia or FeedMedia
//
// See example: examples/media/mediaDelete.go
func (item *Item) Delete() error {
	insta := item.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"media_id": item.ID,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlMediaDelete, item.ID),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// SyncLikers fetch new likers of a media
//
// This function updates Item.Likers value
func (item *Item) SyncLikers() error {
	resp := respLikers{}
	insta := item.media.instagram()
	body, err := insta.sendSimpleRequest(urlMediaLikers, item.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err == nil {
		item.Likers = resp.Users
	}
	return err
}

// Unlike mark media item as unliked.
//
// See example: examples/media/unlike.go
func (item *Item) Unlike() error {
	insta := item.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"media_id": item.ID,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlMediaUnlike, item.ID),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// Like mark media item as liked.
//
// See example: examples/media/like.go
func (item *Item) Like() error {
	insta := item.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"media_id": item.ID,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlMediaLike, item.ID),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// Save saves media item.
//
// You can get saved media using Account.Saved()
func (item *Item) Save() error {
	insta := item.media.instagram()
	data, err := insta.prepareData(
		map[string]interface{}{
			"media_id": item.ID,
		},
	)
	if err != nil {
		return err
	}

	_, err = insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(urlMediaSave, item.ID),
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	return err
}

// Download downloads media item (video or image) with the best quality.
//
// Input parameters are folder and filename. If filename is "" will be saved with
// the default value name.
//
// If file exists it will be saved
// This function makes folder automatically
//
// This function returns an slice of location of downloaded items
// The returned values are the output path of images and videos.
//
// This function does not download CarouselMedia.
//
// See example: examples/media/itemDownload.go
func (item *Item) Download(folder, name string) (imgs, vds string, err error) {
	var u *neturl.URL
	var nname string
	imgFolder := path.Join(folder, "images")
	vidFolder := path.Join(folder, "videos")
	inst := item.media.instagram()

	os.MkdirAll(folder, 0777)
	os.MkdirAll(imgFolder, 0777)
	os.MkdirAll(vidFolder, 0777)

	vds = GetBest(item.Videos)
	if vds != "" {
		if name == "" {
			u, err = neturl.Parse(vds)
			if err != nil {
				return
			}

			nname = path.Join(vidFolder, path.Base(u.Path))
		} else {
			nname = path.Join(vidFolder, name)
		}
		nname = getname(nname)

		vds, err = download(inst, vds, nname)
		return "", vds, err
	}

	imgs = GetBest(item.Images.Versions)
	if imgs != "" {
		if name == "" {
			u, err = neturl.Parse(imgs)
			if err != nil {
				return
			}

			nname = path.Join(imgFolder, path.Base(u.Path))
		} else {
			nname = path.Join(imgFolder, name)
		}
		nname = getname(nname)

		imgs, err = download(inst, imgs, nname)
		return imgs, "", err
	}

	return imgs, vds, fmt.Errorf("cannot find any image or video")
}

// TopLikers returns string slice or single string (inside string slice)
// Depending on TopLikers parameter.
func (item *Item) TopLikers() []string {
	switch s := item.Toplikers.(type) {
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
	switch s := item.Previewcomments.(type) {
	case []interface{}:
		if len(s) == 0 {
			return nil
		}

		switch s[0].(type) {
		case interface{}:
			comments := make([]Comment, 0)
			for i := range s {
				if buf, err := json.Marshal(s[i]); err != nil {
					return nil
				} else {
					comment := &Comment{}

					if err = json.Unmarshal(buf, comment); err != nil {
						return nil
					} else {
						comments = append(comments, *comment)
					}
				}
			}
			return comments
		case string:
			comments := make([]Comment, 0)
			for i := range s {
				comments = append(comments, Comment{
					Text: s[i].(string),
				})
			}
			return comments
		}
	case string:
		comments := []Comment{
			{
				Text: s,
			},
		}
		return comments
	}
	return nil
}

// StoryIsCloseFriends returns a bool
// If the returned value is true the story was published only for close friends
func (item *Item) StoryIsCloseFriends() bool {
	return item.Audience == "besties"
}

//Media interface defines methods for both StoryMedia and FeedMedia.
type Media interface {
	// Next allows pagination
	Next(...interface{}) bool
	// Error returns error (in case it have been occurred)
	Error() error
	// ID returns media id
	ID() string
	// Delete removes media
	Delete() error

	instagram() *Instagram
}

//StoryMedia is the struct that handles the information from the methods to get info about Stories.
type StoryMedia struct {
	inst     *Instagram
	endpoint string
	uid      int64

	err error

	Pk              interface{} `json:"id"`
	LatestReelMedia int64       `json:"latest_reel_media"`
	ExpiringAt      float64     `json:"expiring_at"`
	HaveBeenSeen    float64     `json:"seen"`
	CanReply        bool        `json:"can_reply"`
	Title           string      `json:"title"`
	CanReshare      bool        `json:"can_reshare"`
	ReelType        string      `json:"reel_type"`
	User            User        `json:"user"`
	Items           []Item      `json:"items"`
	ReelMentions    []string    `json:"reel_mentions"`
	PrefetchCount   int         `json:"prefetch_count"`
	// this field can be int or bool
	HasBestiesMedia      interface{} `json:"has_besties_media"`
	StoryRankingToken    string      `json:"story_ranking_token"`
	Broadcasts           []Broadcast `json:"broadcasts"`
	FaceFilterNuxVersion int         `json:"face_filter_nux_version"`
	HasNewNuxStory       bool        `json:"has_new_nux_story"`
	Status               string      `json:"status"`
}

// Delete removes instragram story.
//
// See example: examples/media/deleteStories.go
func (media *StoryMedia) Delete() error {
	insta := media.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"media_id": media.ID(),
		},
	)
	if err == nil {
		_, err = insta.sendRequest(
			&reqOptions{
				Endpoint: fmt.Sprintf(urlMediaDelete, media.ID()),
				Query:    generateSignature(data),
				IsPost:   true,
			},
		)
	}
	return err
}

// ID returns Story id
func (media *StoryMedia) ID() string {
	switch id := media.Pk.(type) {
	case int64:
		return strconv.FormatInt(id, 10)
	case string:
		return id
	}
	return ""
}

func (media *StoryMedia) instagram() *Instagram {
	return media.inst
}

func (media *StoryMedia) setValues() {
	for i := range media.Items {
		setToItem(&media.Items[i], media)
	}
}

// Error returns error happened any error
func (media StoryMedia) Error() error {
	return media.err
}

// Seen marks story as seen.
/*
func (media *StoryMedia) Seen() error {
	insta := media.inst
	data, err := insta.prepareData(
		map[string]interface{}{
			"container_module":   "feed_timeline",
			"live_vods_skipped":  "",
			"nuxes_skipped":      "",
			"nuxes":              "",
			"reels":              "", // TODO xd
			"live_vods":          "",
			"reel_media_skipped": "",
		},
	)
	if err == nil {
		_, err = insta.sendRequest(
			&reqOptions{
				Endpoint: urlMediaSeen, // reel=1&live_vod=0
				Query:    generateSignature(data),
				IsPost:   true,
				UseV2:    true,
			},
		)
	}
	return err
}
*/

type trayRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Sync function is used when Highlight must be sync.
// Highlight must be sync when User.Highlights does not return any object inside StoryMedia slice.
//
// This function does NOT update Stories items.
//
// This function updates StoryMedia.Items
func (media *StoryMedia) Sync() error {
	insta := media.inst
	query := []trayRequest{
		{"SUPPORTED_SDK_VERSIONS", "9.0,10.0,11.0,12.0,13.0,14.0,15.0,16.0,17.0,18.0,19.0,20.0,21.0,22.0,23.0,24.0"},
		{"FACE_TRACKER_VERSION", "10"},
		{"segmentation", "segmentation_enabled"},
		{"COMPRESSION", "ETC2_COMPRESSION"},
	}
	qjson, err := json.Marshal(query)
	if err != nil {
		return err
	}

	id := media.Pk.(string)
	data, err := insta.prepareData(
		map[string]interface{}{
			"user_ids":                   []string{id},
			"supported_capabilities_new": b2s(qjson),
		},
	)
	if err != nil {
		return err
	}

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: urlReelMedia,
			Query:    generateSignature(data),
			IsPost:   true,
		},
	)
	if err == nil {
		resp := trayResp{}
		err = json.Unmarshal(body, &resp)
		if err == nil {
			m, ok := resp.Reels[id]
			if ok {
				media.Items = m.Items
				media.setValues()
				return nil
			}
			err = fmt.Errorf("cannot find %s structure in response", id)
		}
	}
	return err
}

// Next allows pagination after calling:
// User.Stories
//
//
// returns false when list reach the end
// if StoryMedia.Error() is ErrNoMore no problem have been occurred.
func (media *StoryMedia) Next(params ...interface{}) bool {
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
			media.err = ErrNoMore // TODO: See if stories has pagination
			media.setValues()
			return true
		}
	}
	media.err = err
	return false
}

// FeedMedia represent a set of media items
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
	// this is why we recommend Next() usage :')
	NextID interface{} `json:"next_max_id"`
}

// Delete deletes all items in media. Take care...
//
// See example: examples/media/mediaDelete.go
func (media *FeedMedia) Delete() error {
	for i := range media.Items {
		media.Items[i].Delete()
	}
	return nil
}

func (media *FeedMedia) instagram() *Instagram {
	return media.inst
}

// SetInstagram set instagram
func (media *FeedMedia) SetInstagram(inst *Instagram) {
	media.inst = inst
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
			IsPost:   false,
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
		setToItem(&media.Items[i], media)
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
	case json.Number:
		return string(s)
	}
	return ""
}

// Next allows pagination after calling:
// User.Feed
// Params: ranked_content is set to "true" by default, you can set it to false by either passing "false" or false as parameter.
// returns false when list reach the end.
// if FeedMedia.Error() is ErrNoMore no problem have been occurred.
func (media *FeedMedia) Next(params ...interface{}) bool {
	if media.err != nil {
		return false
	}

	insta := media.inst
	endpoint := media.endpoint
	next := media.ID()
	ranked := "true"

	if media.uid != 0 {
		endpoint = fmt.Sprintf(endpoint, media.uid)
	}

	for _, param := range params {
		switch s := param.(type) {
		case string:
			if _, err := strconv.ParseBool(s); err == nil {
				ranked = s
			}
		case bool:
			if !s {
				ranked = "false"
			}
		}
	}
	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: endpoint,
			Query: map[string]string{
				"max_id":         next,
				"rank_token":     insta.rankToken,
				"min_timestamp":  media.timestamp,
				"ranked_content": ranked,
			},
		},
	)
	if err == nil {
		m := FeedMedia{}
		d := json.NewDecoder(bytes.NewReader(body))
		d.UseNumber()
		err = d.Decode(&m)
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

// UploadPhoto post image from io.Reader to instagram.
func (insta *Instagram) UploadPhoto(photo io.Reader, photoCaption string, quality int, filterType int) (Item, error) {
	out := Item{}

	config, err := insta.postPhoto(photo, photoCaption, quality, filterType, false)
	if err != nil {
		return out, err
	}
	data, err := insta.prepareData(config)
	if err != nil {
		return out, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "media/configure/?",
		Query:    generateSignature(data),
		IsPost:   true,
	})
	if err != nil {
		return out, err
	}
	var uploadResult struct {
		Media    Item   `json:"media"`
		UploadID string `json:"upload_id"`
		Status   string `json:"status"`
	}
	err = json.Unmarshal(body, &uploadResult)
	if err != nil {
		return out, err
	}

	if uploadResult.Status != "ok" {
		return out, fmt.Errorf("invalid status, result: %s", uploadResult.Status)
	}

	return uploadResult.Media, nil
}

func (insta *Instagram) postPhoto(photo io.Reader, photoCaption string, quality int, filterType int, isSidecar bool) (map[string]interface{}, error) {
	uploadID := time.Now().Unix()
	photoName := fmt.Sprintf("pending_media_%d.jpg", uploadID)
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	w.WriteField("upload_id", strconv.FormatInt(uploadID, 10))
	w.WriteField("_uuid", insta.uuid)
	w.WriteField("_csrftoken", insta.token)
	var compression = map[string]interface{}{
		"lib_name":    "jt",
		"lib_version": "1.3.0",
		"quality":     quality,
	}
	cBytes, _ := json.Marshal(compression)
	w.WriteField("image_compression", toString(cBytes))
	if isSidecar {
		w.WriteField("is_sidecar", toString(1))
	}
	fw, err := w.CreateFormFile("photo", photoName)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	rdr := io.TeeReader(photo, &buf)
	if _, err = io.Copy(fw, rdr); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", goInstaAPIUrl+"upload/photo/", &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-IG-Capabilities", "3Q4=")
	req.Header.Set("X-IG-Connection-Type", "WIFI")
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-type", w.FormDataContentType())
	req.Header.Set("Connection", "close")
	req.Header.Set("User-Agent", goInstaUserAgent)

	resp, err := insta.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("invalid status code, result: %s", resp.Status)
	}
	var result struct {
		UploadID       string      `json:"upload_id"`
		XsharingNonces interface{} `json:"xsharing_nonces"`
		Status         string      `json:"status"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.Status != "ok" {
		return nil, fmt.Errorf("unknown error, status: %s", result.Status)
	}
	width, height, err := getImageDimensionFromReader(&buf)
	if err != nil {
		return nil, err
	}
	config := map[string]interface{}{
		"media_folder": "Instagram",
		"source_type":  4,
		"caption":      photoCaption,
		"upload_id":    strconv.FormatInt(uploadID, 10),
		"device":       goInstaDeviceSettings,
		"edits": map[string]interface{}{
			"crop_original_size": []int{width * 1.0, height * 1.0},
			"crop_center":        []float32{0.0, 0.0},
			"crop_zoom":          1.0,
			"filter_type":        filterType,
		},
		"extra": map[string]interface{}{
			"source_width":  width,
			"source_height": height,
		},
	}
	return config, nil
}

// UploadAlbum post image from io.Reader to instagram.
func (insta *Instagram) UploadAlbum(photos []io.Reader, photoCaption string, quality int, filterType int) (Item, error) {
	out := Item{}

	var childrenMetadata []map[string]interface{}
	for _, photo := range photos {
		config, err := insta.postPhoto(photo, photoCaption, quality, filterType, true)
		if err != nil {
			return out, err
		}

		childrenMetadata = append(childrenMetadata, config)
	}
	albumUploadID := time.Now().Unix()

	config := map[string]interface{}{
		"caption":           photoCaption,
		"client_sidecar_id": albumUploadID,
		"children_metadata": childrenMetadata,
	}
	data, err := insta.prepareData(config)
	if err != nil {
		return out, err
	}

	body, err := insta.sendRequest(&reqOptions{
		Endpoint: "media/configure_sidecar/?",
		Query:    generateSignature(data),
		IsPost:   true,
	})
	if err != nil {
		return out, err
	}

	var uploadResult struct {
		Media           Item   `json:"media"`
		ClientSideCarID int64  `json:"client_sidecar_id"`
		Status          string `json:"status"`
	}
	err = json.Unmarshal(body, &uploadResult)
	if err != nil {
		return out, err
	}

	if uploadResult.Status != "ok" {
		return out, fmt.Errorf("invalid status, result: %s", uploadResult.Status)
	}

	return uploadResult.Media, nil
}
