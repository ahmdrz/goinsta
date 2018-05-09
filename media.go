package goinsta

import (
	"encoding/json"
	"fmt"
)

// Item represents media items
type Item struct {
	TakenAt                      int          `json:"taken_at"`
	ID                           int64        `json:"pk"`
	IDStr                        string       `json:"id"`
	DeviceTimestamp              int64        `json:"device_timestamp"`
	MediaType                    int          `json:"media_type"`
	Code                         string       `json:"code"`
	ClientCacheKey               string       `json:"client_cache_key"`
	FilterType                   int          `json:"filter_type"`
	CarouselParentID             string       `json:"carousel_parent_id"`
	CarouselMedia                []Item       `json:"carousel_media,omitempty"`
	User                         User         `json:"user"`
	CanViewerReshare             bool         `json:"can_viewer_reshare"`
	Caption                      Caption      `json:"caption"`
	CaptionIsEdited              bool         `json:"caption_is_edited"`
	LikeCount                    int          `json:"like_count"`
	HasLiked                     bool         `json:"has_liked"`
	TopLikers                    []User       `json:"top_likers"`
	CommentLikesEnabled          bool         `json:"comment_likes_enabled"`
	CommentThreadingEnabled      bool         `json:"comment_threading_enabled"`
	HasMoreComments              bool         `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int          `json:"max_num_visible_preview_comments"`
	PreviewComments              []string     `json:"preview_comments"`
	CommentCount                 int          `json:"comment_count"`
	PhotoOfYou                   bool         `json:"photo_of_you"`
	Usertags                     Tag          `json:"usertags,omitempty"`
	FbUserTags                   Tag          `json:"fb_user_tags"`
	CanViewerSave                bool         `json:"can_viewer_save"`
	OrganicTrackingToken         string       `json:"organic_tracking_token"`
	ImageVersions                ImageVersion `json:"image_versions2,omitempty"`
	OriginalWidth                int          `json:"original_width,omitempty"`
	OriginalHeight               int          `json:"original_height,omitempty"`
}

// Media represent a set of media items
type Media struct {
	inst *Instagram

	uid      int64
	endpoint string

	Items               []Item `json:"items"`
	NumResults          int    `json:"num_results"`
	MoreAvailable       bool   `json:"more_available"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Status              string `json:"status"`
	NextID              string `json:"next_max_id"`
}

// Next allows to paginate after calling:
// User.Feed
//
// returns ErrNoMore when list reach the end.
func (media *Media) Next() error {
	insta := media.inst
	endpoint := media.endpoint

	body, err := insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(endpoint, media.uid),
			Query: map[string]string{
				"max_id":         media.NextID,
				"rank_token":     insta.rankToken,
				"min_timestamp":  "",
				"ranked_content": "true",
			},
		},
	)
	if err == nil {
		m := Media{}
		err = json.Unmarshal(body, &m)
		if err == nil {
			*media = m
			media.inst = insta
			media.endpoint = endpoint
			if m.NextID == "" || m.MoreAvailable {
				err = ErrNoMore
			}
		}
	}
	return err
}
