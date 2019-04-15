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
			Media struct {
				TakenAt                      int    `json:"taken_at"`
				Pk                           int64  `json:"pk"`
				ID                           string `json:"id"`
				DeviceTimestamp              int64  `json:"device_timestamp"`
				MediaType                    int    `json:"media_type"`
				Code                         string `json:"code"`
				ClientCacheKey               string `json:"client_cache_key"`
				FilterType                   int    `json:"filter_type"`
				CommentLikesEnabled          bool   `json:"comment_likes_enabled"`
				CommentThreadingEnabled      bool   `json:"comment_threading_enabled"`
				HasMoreComments              bool   `json:"has_more_comments"`
				NextMaxID                    int64  `json:"next_max_id"`
				MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
				PreviewComments              []struct {
					Pk           int64  `json:"pk"`
					UserID       int    `json:"user_id"`
					Text         string `json:"text"`
					Type         int    `json:"type"`
					CreatedAt    int    `json:"created_at"`
					CreatedAtUtc int    `json:"created_at_utc"`
					ContentType  string `json:"content_type"`
					Status       string `json:"status"`
					BitFlags     int    `json:"bit_flags"`
					User         struct {
						Pk            int    `json:"pk"`
						Username      string `json:"username"`
						FullName      string `json:"full_name"`
						IsPrivate     bool   `json:"is_private"`
						ProfilePicURL string `json:"profile_pic_url"`
						ProfilePicID  string `json:"profile_pic_id"`
						IsVerified    bool   `json:"is_verified"`
					} `json:"user"`
					DidReportAsSpam bool  `json:"did_report_as_spam"`
					ShareEnabled    bool  `json:"share_enabled"`
					MediaID         int64 `json:"media_id"`
				} `json:"preview_comments"`
				CanViewMorePreviewComments bool `json:"can_view_more_preview_comments"`
				CommentCount               int  `json:"comment_count"`
				ImageVersions2             struct {
					Candidates []struct {
						Width  int    `json:"width"`
						Height int    `json:"height"`
						URL    string `json:"url"`
					} `json:"candidates"`
				} `json:"image_versions2"`
				OriginalWidth  int `json:"original_width"`
				OriginalHeight int `json:"original_height"`
				Location       struct {
					Pk               int64   `json:"pk"`
					Name             string  `json:"name"`
					Address          string  `json:"address"`
					City             string  `json:"city"`
					ShortName        string  `json:"short_name"`
					Lng              float64 `json:"lng"`
					Lat              float64 `json:"lat"`
					ExternalSource   string  `json:"external_source"`
					FacebookPlacesID int64   `json:"facebook_places_id"`
				} `json:"location"`
				Lat  float64 `json:"lat"`
				Lng  float64 `json:"lng"`
				User struct {
					Pk               int    `json:"pk"`
					Username         string `json:"username"`
					FullName         string `json:"full_name"`
					IsPrivate        bool   `json:"is_private"`
					ProfilePicURL    string `json:"profile_pic_url"`
					ProfilePicID     string `json:"profile_pic_id"`
					FriendshipStatus struct {
						Following       bool `json:"following"`
						OutgoingRequest bool `json:"outgoing_request"`
						IsBestie        bool `json:"is_bestie"`
					} `json:"friendship_status"`
					HasAnonymousProfilePicture bool `json:"has_anonymous_profile_picture"`
					IsUnpublished              bool `json:"is_unpublished"`
					IsFavorite                 bool `json:"is_favorite"`
				} `json:"user"`
				CanViewerReshare bool `json:"can_viewer_reshare"`
				Caption          struct {
					Pk           int64  `json:"pk"`
					UserID       int    `json:"user_id"`
					Text         string `json:"text"`
					Type         int    `json:"type"`
					CreatedAt    int    `json:"created_at"`
					CreatedAtUtc int    `json:"created_at_utc"`
					ContentType  string `json:"content_type"`
					Status       string `json:"status"`
					BitFlags     int    `json:"bit_flags"`
					User         struct {
						Pk               int    `json:"pk"`
						Username         string `json:"username"`
						FullName         string `json:"full_name"`
						IsPrivate        bool   `json:"is_private"`
						ProfilePicURL    string `json:"profile_pic_url"`
						ProfilePicID     string `json:"profile_pic_id"`
						FriendshipStatus struct {
							Following       bool `json:"following"`
							OutgoingRequest bool `json:"outgoing_request"`
							IsBestie        bool `json:"is_bestie"`
						} `json:"friendship_status"`
						HasAnonymousProfilePicture bool `json:"has_anonymous_profile_picture"`
						IsUnpublished              bool `json:"is_unpublished"`
						IsFavorite                 bool `json:"is_favorite"`
					} `json:"user"`
					DidReportAsSpam bool  `json:"did_report_as_spam"`
					ShareEnabled    bool  `json:"share_enabled"`
					MediaID         int64 `json:"media_id"`
				} `json:"caption"`
				CaptionIsEdited      bool          `json:"caption_is_edited"`
				LikeCount            int           `json:"like_count"`
				HasLiked             bool          `json:"has_liked"`
				TopLikers            []interface{} `json:"top_likers"`
				PhotoOfYou           bool          `json:"photo_of_you"`
				CanViewerSave        bool          `json:"can_viewer_save"`
				OrganicTrackingToken string        `json:"organic_tracking_token"`
			} `json:"media"`
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
