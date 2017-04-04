package response

import (
	"strconv"
)

// StatusResponse Status struct point to if response is ok or not
type StatusResponse struct {
	Status string `json:"status"`
}

// Int64Pagination Pagination every pagination have next_max_id
type Int64Pagination struct {
	NextMaxID int64 `json:"next_max_id"`
}

// StringPagination Pagination every pagination have next_max_id
type StringPagination struct {
	NextMaxID string `json:"next_max_id"`
}

// UsersResponse
type UsersResponse struct {
	StatusResponse
	BigList  bool   `json:"big_list"`
	Users    []User `json:"users"`
	PageSize int    `json:"page_size"`
	StringPagination
}

// User , Instagram user informations
type User struct {
	Username                   string `json:"username"`
	HasAnonymousProfilePicture bool   `json:"has_anonymouse_profile_picture"`
	ProfilePictureID           string `json:"profile_pic_id"`
	ProfilePictureURL          string `json:"profile_pic_url"`
	FullName                   string `json:"full_name"`
	PK                         int64  `json:"pk"`
	IsVerified                 bool   `json:"is_verified"`
	IsPrivate                  bool   `json:"is_private"`
	IsFavorite                 bool   `json:"is_favorite"`
}

// StringID return PK as string
func (user User) StringID() string {
	return strconv.FormatInt(user.PK, 10)
}

// FeedsResponse struct contains array of media and can pagination
type FeedsResponse struct {
	StatusResponse
	Items         []MediaItemResponse `json:"items"`
	NumResults    int                 `json:"num_results"`
	AutoLoadMore  bool                `json:"auto_load_more_enabled"`
	MoreAvailable bool                `json:"more_available"`
	StringPagination
}

// TagFeedsResponse struct contains array of MediaItemResponse
// and can pagination
// and array of MediaItemResponse for ranked_items
type TagFeedsResponse struct {
	FeedsResponse
	RankedItems []MediaItemResponse `json:"ranked_items"`
}

// TagRelatedResponse struct contains array of related tags,
// and status
type TagRelatedResponse struct {
	Status  string `json:"status"`
	Related []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"related"`
}

// SearchLocationResponse struct contains array of location venues and status
type SearchLocationResponse struct {
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Venues    []struct {
		ExternalIDSource string  `json:"external_id_source"`
		ExternalID       string  `json:"external_id"`
		Lat              float64 `json:"lat"`
		Lng              float64 `json:"lng"`
		Address          string  `json:"address"`
		Name             string  `json:"name"`
	} `json:"venues"`
}

// MediaItemResponse struct for each media item
type MediaItemResponse struct {
	TakenAt                      int64             `json:"taken_at"`
	PK                           int64             `json:"pk"`
	ID                           string            `json:"id"`
	DeviceTimeStamp              int64             `json:"device_timestamp"`
	MediaType                    int               `json:"media_type"`
	Code                         string            `json:"code"`
	ClientCacheKey               string            `json:"client_cache_key"`
	FilterType                   int               `json:"filter_type"`
	ImageVersions                ImageVersions     `json:"image_versions2"`
	OriginalWidth                int               `json:"original_width"`
	OriginalHeight               int               `json:"original_height"`
	Location                     Location          `json:"location"`
	Lat                          float32           `json:"lat"`
	Lng                          float32           `json:"lng"`
	User                         User              `json:"user"`
	OrganicTrackingToken         string            `json:"organic_tracking_token"`
	LikeCount                    int               `json:"like_count"`
	TopLikers                    []string          `json:"top_likers,omitempty"`
	HasLiked                     bool              `json:"has_liked"`
	HasMoreComments              bool              `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int               `json:"max_num_visible_preview_comments"`
	PreviewComments              []CommentResponse `json:"preview_comments,omitempty"`
	Comments                     []CommentResponse `json:"comments,omitempty"`
	CommentCount                 int               `json:"comment_count"`
	Caption                      Caption           `json:"caption,omitempty"`
	CaptionIsEdited              bool              `json:"caption_is_edited"`
	PhotoOfYou                   bool              `json:"photo_of_you"`
	Int64Pagination
}

// LocationFeedResponse ...
type LocationFeedResponse struct {
	Status              string              `json:"status"`
	AutoLoadMoreEnabled bool                `json:"auto_load_more_enabled"`
	MediaCount          int64               `json:"media_count"`
	NumResults          int64               `json:"num_results"`
	MoreAvailable       bool                `json:"more_available"`
	NextMaxID           string              `json:"next_max_id"`
	Items               []MediaItemResponse `json:"items"`
	RankedItems         []MediaItemResponse `json:"ranked_items"`
}

// ImageVersions struct for image information , urls and etc
type ImageVersions struct {
	Candidates []ImageCandidate `json:"candidates"`
}

// ImageCandidate have urls and image width , height
type ImageCandidate struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Caption struct point to caption of a media
type Caption struct {
	CommentResponse
	HasTranslation bool `json:"has_translation"`
}

// Location struct mean where photo or video taken
type Location struct {
	ExternalSource   string  `json:"external_source"`
	City             string  `json:"city,omitempty"`
	Name             string  `json:"name"`
	FacebookPlacesID int64   `json:"facebook_places_id"`
	Address          string  `json:"address"`
	Lat              float32 `json:"lat"`
	Lng              float32 `json:"lng"`
	PK               int64   `json:"pk"`
}

// CommentResponse struct is a object for comment under media
type CommentResponse struct {
	StatusResponse
	UserID       int64  `json:"user_id"`
	CreatedAtUTC int64  `json:"created_at_utc"`
	CreatedAt    int64  `json:"created_at"`
	BitFlags     int    `json:"bit_flags"`
	User         User   `json:"user"`
	ContentType  string `json:"content_type"`
	Text         string `json:"text"`
	MediaID      int64  `json:"media_id"`
	PK           int64  `json:"pk"`
	Type         int    `json:"type"`
}

// MediaLikersResponse struct for get array of users that like a media
type MediaLikersResponse struct {
	StatusResponse
	UserCount int    `json:"user_count"`
	Users     []User `json:"users"`
}

// ProfileUserResponse struct is current logged in user profile data
// It's very similar to User struct but have more features
// Gender -> 1 male , 2 female , 3 unknown
type ProfileUserResponse struct {
	User
	//Birthday -> what the hell is ?
	PhoneNumber             string           `json:"phone_number"`
	HDProfilePicVersions    []ImageCandidate `json:"hd_profile_pic_versions"`
	Gender                  int              `json:"gender"`
	ShowConversionEditEntry bool             `json:"show_conversion_edit_entry"`
	ExternalLynxURL         string           `json:"external_lynx_url"`
	Biography               string           `json:"biography"`
	HDProfilePicURLInfo     ImageCandidate   `json:"hd_profile_pic_url_info"`
	Email                   string           `json:"email"`
	ExternalURL             string           `json:"external_url"`
}

// ProfileDataResponse have StatusResponse and ProfileUserResponse
type ProfileDataResponse struct {
	StatusResponse
	User ProfileUserResponse `json:"user"`
}

// GetUserID return userinformation
type GetUserID struct {
	StatusResponse
	User UsernameResponse `json:"user"`
}

// GetUsernameResponse return special userinformation
type GetUsernameResponse struct {
	User struct {
		IsPrivate            bool   `json:"is_private"`
		ExternalLynxURL      string `json:"external_lynx_url"`
		IsVerified           bool   `json:"is_verified"`
		MediaCount           int    `json:"media_count"`
		AutoExpandChaining   bool   `json:"auto_expand_chaining"`
		IsFavorite           bool   `json:"is_favorite"`
		FullName             string `json:"full_name"`
		Pk                   int    `json:"pk"`
		FollowingCount       int    `json:"following_count"`
		ExternalURL          string `json:"external_url"`
		ProfilePicURL        string `json:"profile_pic_url"`
		FollowerCount        int    `json:"follower_count"`
		HdProfilePicVersions []struct {
			Height int    `json:"height"`
			Width  int    `json:"width"`
			URL    string `json:"url"`
		} `json:"hd_profile_pic_versions"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		ProfilePicID               string `json:"profile_pic_id"`
		UserTagsCount              int    `json:"usertags_count"`
		Username                   string `json:"username"`
		HdProfilePicURLInfo        struct {
			Height int    `json:"height"`
			Width  int    `json:"width"`
			URL    string `json:"url"`
		} `json:"hd_profile_pic_url_info"`
		GeoMediaCount int    `json:"geo_media_count"`
		IsBusiness    bool   `json:"is_business"`
		Biography     string `json:"biography"`
		HasChaining   bool   `json:"has_chaining"`
	} `json:"user"`
	Status string `json:"status"`
}

// UsernameResponse information of each instagram users
type UsernameResponse struct {
	User
	ExternalURL         string         `json:"external_url"`
	Biography           string         `json:"biography"`
	HDProfilePicURLInfo ImageCandidate `json:"hd_profile_pic_url_info"`
	UserTagsCount       int            `json:"usertags_count"`
	MediaCount          int            `json:"media_count"`
	FollowingCount      int            `json:"following_count"`
	IsBusiness          bool           `json:"is_business"`
	AutoExpandChaining  bool           `json:"auto_expand_chaining"`
	HasChaining         bool           `json:"has_chaining"`
	FollowerCount       int            `json:"follower_count"`
	GeoMediaCount       int            `json:"geo_media_count"`
}

// UploadResponse struct information of upload method
type UploadResponse struct {
	StatusResponse
	UploadID string `json:"upload_id,omitempty"`
	Message  string `json:"message"`
}

// UploadPhotoResponse struct is for uploaded photo response.
type UploadPhotoResponse struct {
	StatusResponse
	Media    MediaItemResponse `json:"media"`
	UploadID string            `json:"upload_id"`
}

// FriendShipResponse struct is for user friendship_status
type FriendShipResponse struct {
	IncomingRequest bool `json:"incoming_request"`
	FollowedBy      bool `json:"followed_by"`
	OutgoingRequest bool `json:"outgoing_request"`
	Following       bool `json:"following"`
	Blocking        bool `json:"blocking"`
	IsPrivate       bool `json:"is_private"`
}

// FollowResponse contains follow response
type FollowResponse struct {
	StatusResponse
	FriendShipStatus FriendShipResponse `json:"friendship_status"`
}

// UnFollowResponse contains UnFollowResponse
type UnFollowResponse struct {
	StatusResponse
	FriendShipStatus FriendShipResponse `json:"friendship_status"`
}

// DirectPendingRequests contains direct pending response
type DirectPendingRequests struct {
	Status               string `json:"status"`
	SeqID                int    `json:"seq_id"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	Inbox                struct {
		UnseenCount   int   `json:"unseen_count"`
		UnseenCountTs int64 `json:"unseen_count_ts"`
		Threads       []struct {
			Named bool `json:"named"`
			Users []struct {
				Username                   string `json:"username"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				FriendshipStatus           struct {
					Following       bool `json:"following"`
					IncomingRequest bool `json:"incoming_request"`
					OutgoingRequest bool `json:"outgoing_request"`
					Blocking        bool `json:"blocking"`
					IsPrivate       bool `json:"is_private"`
				} `json:"friendship_status"`
				ProfilePicURL string `json:"profile_pic_url"`
				ProfilePicID  string `json:"profile_pic_id"`
				FullName      string `json:"full_name"`
				Pk            int    `json:"pk"`
				IsVerified    bool   `json:"is_verified"`
				IsPrivate     bool   `json:"is_private"`
			} `json:"users"`
			ViewerID         int64         `json:"viewer_id"`
			MoreAvailableMin bool          `json:"more_available_min"`
			ThreadID         string        `json:"thread_id"`
			ImageVersions2   ImageVersions `json:"image_versions2"`
			LastActivityAt   int64         `json:"last_activity_at"`
			NextMaxID        string        `json:"next_max_id"`
			IsSpam           bool          `json:"is_spam"`
			LeftUsers        []interface{} `json:"left_users"`
			NextMinID        string        `json:"next_min_id"`
			Muted            bool          `json:"muted"`
			Items            []struct {
				ItemID     string `json:"item_id"`
				ItemType   string `json:"item_type"`
				MediaShare struct {
					TakenAt         int           `json:"taken_at"`
					Pk              int64         `json:"pk"`
					ID              string        `json:"id"`
					DeviceTimestamp int64         `json:"device_timestamp"`
					MediaType       int           `json:"media_type"`
					Code            string        `json:"code"`
					ClientCacheKey  string        `json:"client_cache_key"`
					FilterType      int           `json:"filter_type"`
					ImageVersions2  ImageVersions `json:"image_versions2"`
					OriginalWidth   int           `json:"original_width"`
					OriginalHeight  int           `json:"original_height"`
					ViewCount       float64       `json:"view_count"`
					User            struct {
						Username                   string `json:"username"`
						HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
						IsUnpublished              bool   `json:"is_unpublished"`
						IsFavorite                 bool   `json:"is_favorite"`
						FriendshipStatus           struct {
							Following       bool `json:"following"`
							OutgoingRequest bool `json:"outgoing_request"`
						} `json:"friendship_status"`
						ProfilePicURL string `json:"profile_pic_url"`
						ProfilePicID  string `json:"profile_pic_id"`
						FullName      string `json:"full_name"`
						Pk            int    `json:"pk"`
						IsPrivate     bool   `json:"is_private"`
					} `json:"user"`
					OrganicTrackingToken         string `json:"organic_tracking_token"`
					LikeCount                    int    `json:"like_count"`
					HasLiked                     bool   `json:"has_liked"`
					HasMoreComments              bool   `json:"has_more_comments"`
					NextMaxID                    int64  `json:"next_max_id"`
					MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
					PreviewComments              []struct {
						Status       string `json:"status"`
						UserID       int    `json:"user_id"`
						CreatedAtUtc int    `json:"created_at_utc"`
						CreatedAt    int    `json:"created_at"`
						BitFlags     int    `json:"bit_flags"`
						User         struct {
							Username      string `json:"username"`
							ProfilePicURL string `json:"profile_pic_url"`
							ProfilePicID  string `json:"profile_pic_id"`
							FullName      string `json:"full_name"`
							Pk            int    `json:"pk"`
							IsVerified    bool   `json:"is_verified"`
							IsPrivate     bool   `json:"is_private"`
						} `json:"user"`
						ContentType    string `json:"content_type"`
						Text           string `json:"text"`
						MediaID        int64  `json:"media_id"`
						Pk             int64  `json:"pk"`
						Type           int    `json:"type"`
						HasTranslation bool   `json:"has_translation,omitempty"`
					} `json:"preview_comments"`
					Comments []struct {
						Status       string `json:"status"`
						UserID       int    `json:"user_id"`
						CreatedAtUtc int    `json:"created_at_utc"`
						CreatedAt    int    `json:"created_at"`
						BitFlags     int    `json:"bit_flags"`
						User         struct {
							Username      string `json:"username"`
							ProfilePicURL string `json:"profile_pic_url"`
							ProfilePicID  string `json:"profile_pic_id"`
							FullName      string `json:"full_name"`
							Pk            int    `json:"pk"`
							IsVerified    bool   `json:"is_verified"`
							IsPrivate     bool   `json:"is_private"`
						} `json:"user"`
						ContentType    string `json:"content_type"`
						Text           string `json:"text"`
						MediaID        int64  `json:"media_id"`
						Pk             int64  `json:"pk"`
						Type           int    `json:"type"`
						HasTranslation bool   `json:"has_translation,omitempty"`
					} `json:"comments"`
					CommentCount int `json:"comment_count"`
					Caption      struct {
						Status       string `json:"status"`
						UserID       int    `json:"user_id"`
						CreatedAtUtc int    `json:"created_at_utc"`
						CreatedAt    int    `json:"created_at"`
						BitFlags     int    `json:"bit_flags"`
						User         struct {
							Username                   string `json:"username"`
							HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
							IsUnpublished              bool   `json:"is_unpublished"`
							IsFavorite                 bool   `json:"is_favorite"`
							FriendshipStatus           struct {
								Following       bool `json:"following"`
								OutgoingRequest bool `json:"outgoing_request"`
							} `json:"friendship_status"`
							ProfilePicURL string `json:"profile_pic_url"`
							ProfilePicID  string `json:"profile_pic_id"`
							FullName      string `json:"full_name"`
							Pk            int    `json:"pk"`
							IsPrivate     bool   `json:"is_private"`
						} `json:"user"`
						ContentType    string `json:"content_type"`
						Text           string `json:"text"`
						MediaID        int64  `json:"media_id"`
						Pk             int64  `json:"pk"`
						HasTranslation bool   `json:"has_translation"`
						Type           int    `json:"type"`
					} `json:"caption"`
					CaptionIsEdited bool `json:"caption_is_edited"`
					PhotoOfYou      bool `json:"photo_of_you"`
					VideoVersions   []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Type   int    `json:"type"`
						Height int    `json:"height"`
					} `json:"video_versions"`
					HasAudio      bool `json:"has_audio"`
					VideoDuration int  `json:"video_duration"`
				} `json:"media_share"`
				UserID    int   `json:"user_id"`
				Timestamp int64 `json:"timestamp"`
			} `json:"items"`
			ThreadType       string `json:"thread_type"`
			MoreAvailableMax bool   `json:"more_available_max"`
			ThreadTitle      string `json:"thread_title"`
			Canonical        bool   `json:"canonical"`
			Inviter          struct {
				Username                   string `json:"username"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				ProfilePicURL              string `json:"profile_pic_url"`
				ProfilePicID               string `json:"profile_pic_id"`
				FullName                   string `json:"full_name"`
				Pk                         int    `json:"pk"`
				IsVerified                 bool   `json:"is_verified"`
				IsPrivate                  bool   `json:"is_private"`
			} `json:"inviter"`
			Pending bool `json:"pending"`
		} `json:"threads"`
		MoreAvailable bool `json:"more_available"`
	} `json:"inbox"`
}

// DirectRankedRecipients contains direct ranked_items recipients
type DirectRankedRecipients struct {
	Status           string `json:"status"`
	Filtered         bool   `json:"filtered"`
	Expires          int    `json:"expires"`
	RankedRecipients []struct {
		Thread struct {
			Named bool `json:"named"`
			Users []struct {
				Username                   string `json:"username"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				ProfilePicURL              string `json:"profile_pic_url"`
				ProfilePicID               string `json:"profile_pic_id"`
				FullName                   string `json:"full_name"`
				Pk                         int    `json:"pk"`
				IsVerified                 bool   `json:"is_verified"`
				IsPrivate                  bool   `json:"is_private"`
			} `json:"users"`
			ThreadType  string `json:"thread_type"`
			ThreadID    string `json:"thread_id"`
			ThreadTitle string `json:"thread_title"`
			Pending     bool   `json:"pending"`
		} `json:"thread"`
	} `json:"ranked_recipients"`
}

// DirectThread is a thread of directs
type DirectThread struct {
	Status string `json:"status"`
	Thread struct {
		Named bool `json:"named"`
		Users []struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			FriendshipStatus           struct {
				Following       bool `json:"following"`
				IncomingRequest bool `json:"incoming_request"`
				OutgoingRequest bool `json:"outgoing_request"`
				Blocking        bool `json:"blocking"`
				IsPrivate       bool `json:"is_private"`
			} `json:"friendship_status"`
			ProfilePicURL string `json:"profile_pic_url"`
			ProfilePicID  string `json:"profile_pic_id"`
			FullName      string `json:"full_name"`
			Pk            int    `json:"pk"`
			IsVerified    bool   `json:"is_verified"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"users"`
		ViewerID         int64         `json:"viewer_id"`
		MoreAvailableMin bool          `json:"more_available_min"`
		ThreadID         string        `json:"thread_id"`
		ImageVersions2   ImageVersions `json:"image_versions2"`
		LastActivityAt   int64         `json:"last_activity_at"`
		NextMaxID        string        `json:"next_max_id"`
		Canonical        bool          `json:"canonical"`
		LeftUsers        []interface{} `json:"left_users"`
		NextMinID        string        `json:"next_min_id"`
		Muted            bool          `json:"muted"`
		Items            []struct {
			ItemID     string `json:"item_id"`
			ItemType   string `json:"item_type"`
			MediaShare struct {
				TakenAt         int           `json:"taken_at"`
				Pk              int64         `json:"pk"`
				ID              string        `json:"id"`
				DeviceTimestamp int           `json:"device_timestamp"`
				MediaType       int           `json:"media_type"`
				Code            string        `json:"code"`
				ClientCacheKey  string        `json:"client_cache_key"`
				FilterType      int           `json:"filter_type"`
				ImageVersions2  ImageVersions `json:"image_versions2"`
				OriginalWidth   int           `json:"original_width"`
				OriginalHeight  int           `json:"original_height"`
				ViewCount       float64       `json:"view_count"`
				User            struct {
					Username                   string `json:"username"`
					HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
					IsUnpublished              bool   `json:"is_unpublished"`
					IsFavorite                 bool   `json:"is_favorite"`
					FriendshipStatus           struct {
						Following       bool `json:"following"`
						OutgoingRequest bool `json:"outgoing_request"`
					} `json:"friendship_status"`
					ProfilePicURL string `json:"profile_pic_url"`
					ProfilePicID  string `json:"profile_pic_id"`
					FullName      string `json:"full_name"`
					Pk            int    `json:"pk"`
					IsPrivate     bool   `json:"is_private"`
				} `json:"user"`
				OrganicTrackingToken         string `json:"organic_tracking_token"`
				LikeCount                    int    `json:"like_count"`
				HasLiked                     bool   `json:"has_liked"`
				HasMoreComments              bool   `json:"has_more_comments"`
				NextMaxID                    int64  `json:"next_max_id"`
				MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
				PreviewComments              []struct {
					Status       string `json:"status"`
					UserID       int    `json:"user_id"`
					CreatedAtUtc int    `json:"created_at_utc"`
					CreatedAt    int    `json:"created_at"`
					BitFlags     int    `json:"bit_flags"`
					User         struct {
						Username      string `json:"username"`
						ProfilePicURL string `json:"profile_pic_url"`
						FullName      string `json:"full_name"`
						Pk            int    `json:"pk"`
						IsVerified    bool   `json:"is_verified"`
						IsPrivate     bool   `json:"is_private"`
					} `json:"user"`
					ContentType string `json:"content_type"`
					Text        string `json:"text"`
					MediaID     int64  `json:"media_id"`
					Pk          int64  `json:"pk"`
					Type        int    `json:"type"`
				} `json:"preview_comments"`
				Comments []struct {
					Status       string `json:"status"`
					UserID       int    `json:"user_id"`
					CreatedAtUtc int    `json:"created_at_utc"`
					CreatedAt    int    `json:"created_at"`
					BitFlags     int    `json:"bit_flags"`
					User         struct {
						Username      string `json:"username"`
						ProfilePicURL string `json:"profile_pic_url"`
						FullName      string `json:"full_name"`
						Pk            int    `json:"pk"`
						IsVerified    bool   `json:"is_verified"`
						IsPrivate     bool   `json:"is_private"`
					} `json:"user"`
					ContentType string `json:"content_type"`
					Text        string `json:"text"`
					MediaID     int64  `json:"media_id"`
					Pk          int64  `json:"pk"`
					Type        int    `json:"type"`
				} `json:"comments"`
				CommentCount int `json:"comment_count"`
				Caption      struct {
					Status       string `json:"status"`
					UserID       int    `json:"user_id"`
					CreatedAtUtc int    `json:"created_at_utc"`
					CreatedAt    int    `json:"created_at"`
					BitFlags     int    `json:"bit_flags"`
					User         struct {
						Username                   string `json:"username"`
						HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
						IsUnpublished              bool   `json:"is_unpublished"`
						IsFavorite                 bool   `json:"is_favorite"`
						FriendshipStatus           struct {
							Following       bool `json:"following"`
							OutgoingRequest bool `json:"outgoing_request"`
						} `json:"friendship_status"`
						ProfilePicURL string `json:"profile_pic_url"`
						ProfilePicID  string `json:"profile_pic_id"`
						FullName      string `json:"full_name"`
						Pk            int    `json:"pk"`
						IsPrivate     bool   `json:"is_private"`
					} `json:"user"`
					ContentType string `json:"content_type"`
					Text        string `json:"text"`
					MediaID     int64  `json:"media_id"`
					Pk          int64  `json:"pk"`
					Type        int    `json:"type"`
				} `json:"caption"`
				CaptionIsEdited bool `json:"caption_is_edited"`
				PhotoOfYou      bool `json:"photo_of_you"`
				VideoVersions   []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Type   int    `json:"type"`
					Height int    `json:"height"`
				} `json:"video_versions"`
				HasAudio      bool    `json:"has_audio"`
				VideoDuration float64 `json:"video_duration"`
			} `json:"media_share"`
			UserID    int   `json:"user_id"`
			Timestamp int64 `json:"timestamp"`
		} `json:"items"`
		ThreadType       string `json:"thread_type"`
		MoreAvailableMax bool   `json:"more_available_max"`
		ThreadTitle      string `json:"thread_title"`
		LastSeenAt       struct {
			Num1572292791 struct {
				ItemID    string `json:"item_id"`
				Timestamp string `json:"timestamp"`
			} `json:"1572292791"`
			Num4043092277 struct {
				ItemID    string `json:"item_id"`
				Timestamp string `json:"timestamp"`
			} `json:"4043092277"`
		} `json:"last_seen_at"`
		Inviter struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			ProfilePicURL              string `json:"profile_pic_url"`
			ProfilePicID               string `json:"profile_pic_id"`
			FullName                   string `json:"full_name"`
			Pk                         int    `json:"pk"`
			IsVerified                 bool   `json:"is_verified"`
			IsPrivate                  bool   `json:"is_private"`
		} `json:"inviter"`
		Pending bool `json:"pending"`
	} `json:"thread"`
}

// UserFeedResponse contains user feeds
type UserFeedResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []struct {
		TakenAt         int64         `json:"taken_at"`
		Pk              int64         `json:"pk"`
		ID              string        `json:"id"`
		DeviceTimestamp int64         `json:"device_timestamp"`
		MediaType       int           `json:"media_type"`
		Code            string        `json:"code"`
		ClientCacheKey  string        `json:"client_cache_key"`
		FilterType      int           `json:"filter_type"`
		ImageVersions2  ImageVersions `json:"image_versions2"`
		OriginalWidth   int           `json:"original_width"`
		OriginalHeight  int           `json:"original_height"`
		User            struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			IsUnpublished              bool   `json:"is_unpublished"`
			IsFavorite                 bool   `json:"is_favorite"`
			ProfilePicURL              string `json:"profile_pic_url"`
			ProfilePicID               string `json:"profile_pic_id"`
			FullName                   string `json:"full_name"`
			Pk                         int    `json:"pk"`
			IsVerified                 bool   `json:"is_verified"`
			IsPrivate                  bool   `json:"is_private"`
		} `json:"user"`
		OrganicTrackingToken         string        `json:"organic_tracking_token"`
		LikeCount                    int           `json:"like_count"`
		TopLikers                    []interface{} `json:"top_likers"`
		HasLiked                     bool          `json:"has_liked"`
		HasMoreComments              bool          `json:"has_more_comments"`
		MaxNumVisiblePreviewComments int           `json:"max_num_visible_preview_comments"`
		PreviewComments              []interface{} `json:"preview_comments"`
		Comments                     []interface{} `json:"comments"`
		CommentCount                 int           `json:"comment_count"`
		Caption                      struct {
			Status       string `json:"status"`
			UserID       int    `json:"user_id"`
			CreatedAtUtc int    `json:"created_at_utc"`
			CreatedAt    int    `json:"created_at"`
			BitFlags     int    `json:"bit_flags"`
			User         struct {
				Username                   string `json:"username"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsFavorite                 bool   `json:"is_favorite"`
				ProfilePicURL              string `json:"profile_pic_url"`
				ProfilePicID               string `json:"profile_pic_id"`
				FullName                   string `json:"full_name"`
				Pk                         int    `json:"pk"`
				IsVerified                 bool   `json:"is_verified"`
				IsPrivate                  bool   `json:"is_private"`
			} `json:"user"`
			ContentType    string `json:"content_type"`
			Text           string `json:"text"`
			MediaID        int64  `json:"media_id"`
			Pk             int64  `json:"pk"`
			HasTranslation bool   `json:"has_translation"`
			Type           int    `json:"type"`
		} `json:"caption"`
		CaptionIsEdited bool `json:"caption_is_edited"`
		PhotoOfYou      bool `json:"photo_of_you"`
		UserTags        struct {
			In []struct {
				Position    []float64   `json:"position"`
				TimeInVideo interface{} `json:"time_in_video"`
				User        struct {
					Username      string `json:"username"`
					ProfilePicURL string `json:"profile_pic_url"`
					FullName      string `json:"full_name"`
					Pk            int64  `json:"pk"`
					IsVerified    bool   `json:"is_verified"`
					IsPrivate     bool   `json:"is_private"`
				} `json:"user"`
			} `json:"in"`
		} `json:"usertags,omitempty"`
		ViewCount     float64 `json:"view_count,omitempty"`
		VideoVersions []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Type   int    `json:"type"`
			Height int    `json:"height"`
		} `json:"video_versions,omitempty"`
		HasAudio      bool    `json:"has_audio,omitempty"`
		VideoDuration float64 `json:"video_duration,omitempty"`
		NextMaxID     int64   `json:"next_max_id,omitempty"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	NextMaxID     string `json:"next_max_id"`
}

// DirectMessageResponse contains direct messages
type DirectMessageResponse struct {
	Status  string `json:"status"`
	Threads []struct {
		Named bool `json:"named"`
		Users []struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			FriendshipStatus           struct {
				Following       bool `json:"following"`
				IncomingRequest bool `json:"incoming_request"`
				OutgoingRequest bool `json:"outgoing_request"`
				Blocking        bool `json:"blocking"`
				IsPrivate       bool `json:"is_private"`
			} `json:"friendship_status"`
			ProfilePicURL string `json:"profile_pic_url"`
			ProfilePicID  string `json:"profile_pic_id"`
			FullName      string `json:"full_name"`
			Pk            int    `json:"pk"`
			IsVerified    bool   `json:"is_verified"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"users"`
		ViewerID         int64         `json:"viewer_id"`
		MoreAvailableMin bool          `json:"more_available_min"`
		ThreadID         string        `json:"thread_id"`
		LastActivityAt   int64         `json:"last_activity_at"`
		NextMaxID        string        `json:"next_max_id"`
		Canonical        bool          `json:"canonical"`
		LeftUsers        []interface{} `json:"left_users"`
		NextMinID        string        `json:"next_min_id"`
		Muted            bool          `json:"muted"`
		Items            []struct {
			UserID        int64  `json:"user_id"`
			Text          string `json:"text"`
			ItemType      string `json:"item_type"`
			Timestamp     int64  `json:"timestamp"`
			ItemID        string `json:"item_id"`
			ClientContext string `json:"client_context"`
		} `json:"items"`
		ThreadType       string `json:"thread_type"`
		MoreAvailableMax bool   `json:"more_available_max"`
		ThreadTitle      string `json:"thread_title"`
		LastSeenAt       struct {
			Num1572292791 struct {
				ItemID    string `json:"item_id"`
				Timestamp string `json:"timestamp"`
			} `json:"1572292791"`
			Num4178028611 struct {
				ItemID    string `json:"item_id"`
				Timestamp string `json:"timestamp"`
			} `json:"4178028611"`
		} `json:"last_seen_at"`
		Inviter struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			ProfilePicURL              string `json:"profile_pic_url"`
			ProfilePicID               string `json:"profile_pic_id"`
			FullName                   string `json:"full_name"`
			Pk                         int64  `json:"pk"`
			IsVerified                 bool   `json:"is_verified"`
			IsPrivate                  bool   `json:"is_private"`
		} `json:"inviter"`
		Pending bool `json:"pending"`
	} `json:"threads"`
}

// SearchUserResponse is for user search response
type SearchUserResponse struct {
	HasMore    bool   `json:"has_more"`
	Status     string `json:"status"`
	NumResults int    `json:"num_results"`
	Users      []struct {
		Username                   string `json:"username"`
		HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
		Byline                     string `json:"byline"`
		FriendshipStatus           struct {
			Following       bool `json:"following"`
			IncomingRequest bool `json:"incoming_request"`
			OutgoingRequest bool `json:"outgoing_request"`
			IsPrivate       bool `json:"is_private"`
		} `json:"friendship_status"`
		UnseenCount          int     `json:"unseen_count"`
		MutualFollowersCount float64 `json:"mutual_followers_count"`
		ProfilePicURL        string  `json:"profile_pic_url"`
		FullName             string  `json:"full_name"`
		FollowerCount        int     `json:"follower_count"`
		Pk                   int     `json:"pk"`
		IsVerified           bool    `json:"is_verified"`
		IsPrivate            bool    `json:"is_private"`
		ProfilePicID         string  `json:"profile_pic_id,omitempty"`
	} `json:"users"`
}

// ExploreResponse is data from explore in Instagram
type ExploreResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []struct {
		Stories struct {
			Tray []struct {
				CanReply   bool `json:"can_reply"`
				ExpiringAt int  `json:"expiring_at"`
				User       struct {
					Username         string `json:"username"`
					FriendshipStatus struct {
						IncomingRequest bool `json:"incoming_request"`
						FollowedBy      bool `json:"followed_by"`
						OutgoingRequest bool `json:"outgoing_request"`
						Following       bool `json:"following"`
						Blocking        bool `json:"blocking"`
						IsPrivate       bool `json:"is_private"`
					} `json:"friendship_status"`
					ProfilePicURL string `json:"profile_pic_url"`
					ProfilePicID  string `json:"profile_pic_id"`
					FullName      string `json:"full_name"`
					Pk            int    `json:"pk"`
					IsVerified    bool   `json:"is_verified"`
					IsPrivate     bool   `json:"is_private"`
				} `json:"user"`
				SourceToken        string `json:"source_token"`
				Seen               int    `json:"seen"`
				LatestReelMedia    int    `json:"latest_reel_media"`
				ID                 int    `json:"id"`
				RankedPosition     int    `json:"ranked_position"`
				SeenRankedPosition int    `json:"seen_ranked_position"`
			} `json:"tray"`
			ID         int64 `json:"id"`
			IsPortrait bool  `json:"is_portrait"`
		} `json:"stories,omitempty"`
		Media struct {
			TakenAt         int           `json:"taken_at"`
			Pk              int64         `json:"pk"`
			ID              string        `json:"id"`
			DeviceTimestamp int64         `json:"device_timestamp"`
			MediaType       int           `json:"media_type"`
			Code            string        `json:"code"`
			ClientCacheKey  string        `json:"client_cache_key"`
			FilterType      int           `json:"filter_type"`
			ImageVersions2  ImageVersions `json:"image_versions2"`
			OriginalWidth   int           `json:"original_width"`
			OriginalHeight  int           `json:"original_height"`
			User            struct {
				Username                   string `json:"username"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsFavorite                 bool   `json:"is_favorite"`
				FriendshipStatus           struct {
					Following       bool `json:"following"`
					OutgoingRequest bool `json:"outgoing_request"`
				} `json:"friendship_status"`
				ProfilePicURL string `json:"profile_pic_url"`
				ProfilePicID  string `json:"profile_pic_id"`
				FullName      string `json:"full_name"`
				Pk            int    `json:"pk"`
				IsVerified    bool   `json:"is_verified"`
				IsPrivate     bool   `json:"is_private"`
			} `json:"user"`
			OrganicTrackingToken         string `json:"organic_tracking_token"`
			LikeCount                    int    `json:"like_count"`
			HasLiked                     bool   `json:"has_liked"`
			HasMoreComments              bool   `json:"has_more_comments"`
			NextMaxID                    int64  `json:"next_max_id"`
			MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
			PreviewComments              []struct {
				Status       string `json:"status"`
				UserID       int    `json:"user_id"`
				CreatedAtUtc int    `json:"created_at_utc"`
				CreatedAt    int    `json:"created_at"`
				BitFlags     int    `json:"bit_flags"`
				User         struct {
					Username      string `json:"username"`
					ProfilePicURL string `json:"profile_pic_url"`
					ProfilePicID  string `json:"profile_pic_id"`
					FullName      string `json:"full_name"`
					Pk            int    `json:"pk"`
					IsVerified    bool   `json:"is_verified"`
					IsPrivate     bool   `json:"is_private"`
				} `json:"user"`
				ContentType string `json:"content_type"`
				Text        string `json:"text"`
				MediaID     int64  `json:"media_id"`
				Pk          int64  `json:"pk"`
				Type        int    `json:"type"`
			} `json:"preview_comments"`
			CommentCount int `json:"comment_count"`
			Caption      struct {
				Status       string `json:"status"`
				UserID       int    `json:"user_id"`
				CreatedAtUtc int    `json:"created_at_utc"`
				CreatedAt    int    `json:"created_at"`
				BitFlags     int    `json:"bit_flags"`
				User         struct {
					Username                   string `json:"username"`
					HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
					IsUnpublished              bool   `json:"is_unpublished"`
					IsFavorite                 bool   `json:"is_favorite"`
					FriendshipStatus           struct {
						Following       bool `json:"following"`
						OutgoingRequest bool `json:"outgoing_request"`
					} `json:"friendship_status"`
					ProfilePicURL string `json:"profile_pic_url"`
					ProfilePicID  string `json:"profile_pic_id"`
					FullName      string `json:"full_name"`
					Pk            int    `json:"pk"`
					IsVerified    bool   `json:"is_verified"`
					IsPrivate     bool   `json:"is_private"`
				} `json:"user"`
				ContentType    string `json:"content_type"`
				Text           string `json:"text"`
				MediaID        int64  `json:"media_id"`
				Pk             int64  `json:"pk"`
				HasTranslation bool   `json:"has_translation"`
				Type           int    `json:"type"`
			} `json:"caption"`
			CaptionIsEdited    bool   `json:"caption_is_edited"`
			PhotoOfYou         bool   `json:"photo_of_you"`
			Algorithm          string `json:"algorithm"`
			ExploreContext     string `json:"explore_context"`
			ExploreSourceToken string `json:"explore_source_token"`
			Explore            struct {
				Explanation string `json:"explanation"`
				ActorID     int    `json:"actor_id"`
				SourceToken string `json:"source_token"`
			} `json:"explore"`
			ImpressionToken string `json:"impression_token"`
		} `json:"media,omitempty"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	NextMaxID     string `json:"next_max_id"`
	MaxID         string `json:"max_id"`
}

// MediaInfoResponse contains media information
type MediaInfoResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []struct {
		TakenAt         int           `json:"taken_at"`
		Pk              int64         `json:"pk"`
		ID              string        `json:"id"`
		DeviceTimestamp int           `json:"device_timestamp"`
		MediaType       int           `json:"media_type"`
		Code            string        `json:"code"`
		ClientCacheKey  string        `json:"client_cache_key"`
		FilterType      int           `json:"filter_type"`
		ImageVersions2  ImageVersions `json:"image_versions2"`
		OriginalWidth   int           `json:"original_width"`
		OriginalHeight  int           `json:"original_height"`
		Location        struct {
			ExternalSource   string  `json:"external_source"`
			City             string  `json:"city"`
			Name             string  `json:"name"`
			FacebookPlacesID int64   `json:"facebook_places_id"`
			Address          string  `json:"address"`
			Lat              float64 `json:"lat"`
			Pk               int     `json:"pk"`
			Lng              float64 `json:"lng"`
		} `json:"location"`
		ViewCount float64 `json:"view_count"`
		Lat       float64 `json:"lat"`
		Lng       float64 `json:"lng"`
		User      struct {
			Username                   string `json:"username"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			IsUnpublished              bool   `json:"is_unpublished"`
			IsFavorite                 bool   `json:"is_favorite"`
			FriendshipStatus           struct {
				Following       bool `json:"following"`
				OutgoingRequest bool `json:"outgoing_request"`
			} `json:"friendship_status"`
			ProfilePicURL string `json:"profile_pic_url"`
			ProfilePicID  string `json:"profile_pic_id"`
			FullName      string `json:"full_name"`
			Pk            int    `json:"pk"`
			IsVerified    bool   `json:"is_verified"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"user"`
		OrganicTrackingToken         string        `json:"organic_tracking_token"`
		LikeCount                    int           `json:"like_count"`
		TopLikers                    []interface{} `json:"top_likers"`
		HasLiked                     bool          `json:"has_liked"`
		HasMoreComments              bool          `json:"has_more_comments"`
		NextMaxID                    int64         `json:"next_max_id"`
		MaxNumVisiblePreviewComments int           `json:"max_num_visible_preview_comments"`
		PreviewComments              []struct {
			Status       string `json:"status"`
			UserID       int    `json:"user_id"`
			CreatedAtUtc int    `json:"created_at_utc"`
			CreatedAt    int    `json:"created_at"`
			BitFlags     int    `json:"bit_flags"`
			User         struct {
				Username      string `json:"username"`
				ProfilePicURL string `json:"profile_pic_url"`
				ProfilePicID  string `json:"profile_pic_id"`
				FullName      string `json:"full_name"`
				Pk            int    `json:"pk"`
				IsVerified    bool   `json:"is_verified"`
				IsPrivate     bool   `json:"is_private"`
			} `json:"user"`
			ContentType string `json:"content_type"`
			Text        string `json:"text"`
			MediaID     int64  `json:"media_id"`
			Pk          int64  `json:"pk"`
			Type        int    `json:"type"`
		} `json:"preview_comments"`
		CommentCount int `json:"comment_count"`
		Caption      struct {
			Status       string `json:"status"`
			UserID       int    `json:"user_id"`
			CreatedAtUtc int    `json:"created_at_utc"`
			CreatedAt    int    `json:"created_at"`
			BitFlags     int    `json:"bit_flags"`
			User         struct {
				Username                   string `json:"username"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsFavorite                 bool   `json:"is_favorite"`
				FriendshipStatus           struct {
					Following       bool `json:"following"`
					OutgoingRequest bool `json:"outgoing_request"`
				} `json:"friendship_status"`
				ProfilePicURL string `json:"profile_pic_url"`
				ProfilePicID  string `json:"profile_pic_id"`
				FullName      string `json:"full_name"`
				Pk            int    `json:"pk"`
				IsVerified    bool   `json:"is_verified"`
				IsPrivate     bool   `json:"is_private"`
			} `json:"user"`
			ContentType string `json:"content_type"`
			Text        string `json:"text"`
			MediaID     int64  `json:"media_id"`
			Pk          int64  `json:"pk"`
			Type        int    `json:"type"`
		} `json:"caption"`
		CaptionIsEdited bool `json:"caption_is_edited"`
		PhotoOfYou      bool `json:"photo_of_you"`
		VideoVersions   []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Type   int    `json:"type"`
			Height int    `json:"height"`
		} `json:"video_versions"`
		HasAudio      bool    `json:"has_audio"`
		VideoDuration float64 `json:"video_duration"`
	} `json:"items"`
	MoreAvailable       bool `json:"more_available"`
	CommentLikesEnabled bool `json:"comment_likes_enabled"`
}

// UserFriendShipResponse is about user_friend_ship response
type UserFriendShipResponse struct {
	Following       bool   `json:"following"`
	FollowedBy      bool   `json:"followed_by"`
	Status          string `json:"status"`
	IsPrivate       bool   `json:"is_private"`
	IsMutingReel    bool   `json:"is_muting_reel"`
	OutgoingRequest bool   `json:"outgoing_request"`
	IsBlockingReel  bool   `json:"is_blocking_reel"`
	Blocking        bool   `json:"blocking"`
	IncomingRequest bool   `json:"incoming_request"`
}

// GetPopularFeedResponse contains popular feeds
type GetPopularFeedResponse struct {
	MaxID               string `json:"max_id"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	NextMaxID           string `json:"next_max_id"`
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	Items               []struct {
		TakenAt         int           `json:"taken_at"`
		Pk              int64         `json:"pk"`
		ID              string        `json:"id"`
		DeviceTimestamp int64         `json:"device_timestamp"`
		MediaType       int           `json:"media_type"`
		Code            string        `json:"code"`
		ClientCacheKey  string        `json:"client_cache_key"`
		FilterType      int           `json:"filter_type"`
		ImageVersions2  ImageVersions `json:"image_versions2"`
		OriginalWidth   int           `json:"original_width"`
		OriginalHeight  int           `json:"original_height"`
		User            struct {
			Username                   string `json:"username"`
			IsPrivate                  bool   `json:"is_private"`
			FullName                   string `json:"full_name"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			Pk                         int    `json:"pk"`
			ProfilePicID               string `json:"profile_pic_id"`
			IsVerified                 bool   `json:"is_verified"`
			IsUnpublished              bool   `json:"is_unpublished"`
			IsFavorite                 bool   `json:"is_favorite"`
			ProfilePicURL              string `json:"profile_pic_url"`
			FriendshipStatus           struct {
				OutgoingRequest bool `json:"outgoing_request"`
				Following       bool `json:"following"`
			} `json:"friendship_status"`
		} `json:"user"`
		OrganicTrackingToken         string `json:"organic_tracking_token"`
		LikeCount                    int    `json:"like_count"`
		HasLiked                     bool   `json:"has_liked"`
		CommentLikesEnabled          bool   `json:"comment_likes_enabled"`
		HasMoreComments              bool   `json:"has_more_comments"`
		NextMaxID                    int64  `json:"next_max_id,omitempty"`
		MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
		PreviewComments              []struct {
			MediaID      int64  `json:"media_id"`
			BitFlags     int    `json:"bit_flags"`
			Text         string `json:"text"`
			Type         int    `json:"type"`
			Status       string `json:"status"`
			Pk           int64  `json:"pk"`
			CreatedAtUtc int    `json:"created_at_utc"`
			CreatedAt    int    `json:"created_at"`
			User         struct {
				Username      string `json:"username"`
				IsPrivate     bool   `json:"is_private"`
				FullName      string `json:"full_name"`
				Pk            int64  `json:"pk"`
				ProfilePicID  string `json:"profile_pic_id"`
				IsVerified    bool   `json:"is_verified"`
				ProfilePicURL string `json:"profile_pic_url"`
			} `json:"user"`
			ContentType string `json:"content_type"`
			UserID      int64  `json:"user_id"`
		} `json:"preview_comments"`
		CommentCount int `json:"comment_count"`
		Caption      struct {
			CreatedAt      int    `json:"created_at"`
			CreatedAtUtc   int    `json:"created_at_utc"`
			HasTranslation bool   `json:"has_translation"`
			UserID         int    `json:"user_id"`
			MediaID        int64  `json:"media_id"`
			Text           string `json:"text"`
			Type           int    `json:"type"`
			Pk             int64  `json:"pk"`
			Status         string `json:"status"`
			BitFlags       int    `json:"bit_flags"`
			User           struct {
				Username                   string `json:"username"`
				IsPrivate                  bool   `json:"is_private"`
				FullName                   string `json:"full_name"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				Pk                         int    `json:"pk"`
				ProfilePicID               string `json:"profile_pic_id"`
				IsVerified                 bool   `json:"is_verified"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsFavorite                 bool   `json:"is_favorite"`
				ProfilePicURL              string `json:"profile_pic_url"`
				FriendshipStatus           struct {
					OutgoingRequest bool `json:"outgoing_request"`
					Following       bool `json:"following"`
				} `json:"friendship_status"`
			} `json:"user"`
			ContentType string `json:"content_type"`
		} `json:"caption"`
		CaptionIsEdited bool `json:"caption_is_edited"`
		PhotoOfYou      bool `json:"photo_of_you"`
		UserTags        struct {
			In []struct {
				TimeInVideo interface{} `json:"time_in_video"`
				User        struct {
					Username      string `json:"username"`
					IsPrivate     bool   `json:"is_private"`
					FullName      string `json:"full_name"`
					Pk            int    `json:"pk"`
					ProfilePicID  string `json:"profile_pic_id"`
					IsVerified    bool   `json:"is_verified"`
					ProfilePicURL string `json:"profile_pic_url"`
				} `json:"user"`
				Position []float64 `json:"position"`
			} `json:"in"`
		} `json:"usertags,omitempty"`
		Algorithm          string `json:"algorithm"`
		ExploreContext     string `json:"explore_context"`
		ExploreSourceToken string `json:"explore_source_token"`
		Explore            struct {
			SourceToken string `json:"source_token"`
			ActorID     int    `json:"actor_id"`
			Explanation string `json:"explanation"`
		} `json:"explore"`
		ImpressionToken string  `json:"impression_token"`
		ViewCount       float64 `json:"view_count,omitempty"`
		VideoVersions   []struct {
			Height int    `json:"height"`
			Width  int    `json:"width"`
			URL    string `json:"url"`
			Type   int    `json:"type"`
		} `json:"video_versions,omitempty"`
		HasAudio      bool    `json:"has_audio,omitempty"`
		VideoDuration float64 `json:"video_duration,omitempty"`
	} `json:"items"`
	MoreAvailable bool `json:"more_available"`
}

// DirectListResponse is list of directs
type DirectListResponse struct {
	PendingRequestsTotal int    `json:"pending_requests_total"`
	SeqID                int    `json:"seq_id"`
	Status               string `json:"status"`
	Inbox                struct {
		HasOlder      bool   `json:"has_older"`
		OldestCursor  string `json:"oldest_cursor"`
		UnseenCount   int    `json:"unseen_count"`
		UnseenCountTs int64  `json:"unseen_count_ts"`
		Threads       []struct {
			ThreadType     string `json:"thread_type"`
			LastActivityAt int64  `json:"last_activity_at"`
			LastSeenAt     struct {
				Num4178028611 struct {
					Timestamp string `json:"timestamp"`
					ItemID    string `json:"item_id"`
				} `json:"4178028611"`
			} `json:"last_seen_at"`
			ViewerID     int64         `json:"viewer_id"`
			OldestCursor string        `json:"oldest_cursor"`
			LeftUsers    []interface{} `json:"left_users"`
			ThreadID     string        `json:"thread_id"`
			Inviter      struct {
				Username                   string `json:"username"`
				IsPrivate                  bool   `json:"is_private"`
				ProfilePicURL              string `json:"profile_pic_url"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				Pk                         int    `json:"pk"`
				ProfilePicID               string `json:"profile_pic_id"`
				IsVerified                 bool   `json:"is_verified"`
				FullName                   string `json:"full_name"`
			} `json:"inviter"`
			ThreadTitle string `json:"thread_title"`
			Items       []struct {
				Timestamp  int64  `json:"timestamp"`
				ItemID     string `json:"item_id"`
				MediaShare struct {
					TakenAt         int           `json:"taken_at"`
					Pk              int64         `json:"pk"`
					ID              string        `json:"id"`
					DeviceTimestamp int           `json:"device_timestamp"`
					MediaType       int           `json:"media_type"`
					Code            string        `json:"code"`
					ClientCacheKey  string        `json:"client_cache_key"`
					FilterType      int           `json:"filter_type"`
					ImageVersions2  ImageVersions `json:"image_versions2"`
					OriginalWidth   int           `json:"original_width"`
					OriginalHeight  int           `json:"original_height"`
					ViewCount       float64       `json:"view_count"`
					User            struct {
						Username         string `json:"username"`
						IsUnpublished    bool   `json:"is_unpublished"`
						IsPrivate        bool   `json:"is_private"`
						FriendshipStatus struct {
							Following       bool `json:"following"`
							OutgoingRequest bool `json:"outgoing_request"`
						} `json:"friendship_status"`
						ProfilePicURL              string `json:"profile_pic_url"`
						HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
						IsFavorite                 bool   `json:"is_favorite"`
						Pk                         int    `json:"pk"`
						ProfilePicID               string `json:"profile_pic_id"`
						FullName                   string `json:"full_name"`
					} `json:"user"`
					OrganicTrackingToken         string `json:"organic_tracking_token"`
					LikeCount                    int    `json:"like_count"`
					HasLiked                     bool   `json:"has_liked"`
					CommentLikesEnabled          bool   `json:"comment_likes_enabled"`
					HasMoreComments              bool   `json:"has_more_comments"`
					NextMaxID                    int64  `json:"next_max_id"`
					MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
					PreviewComments              []struct {
						MediaID     int64  `json:"media_id"`
						BitFlags    int    `json:"bit_flags"`
						Type        int    `json:"type"`
						Status      string `json:"status"`
						ContentType string `json:"content_type"`
						UserID      int    `json:"user_id"`
						CreatedAt   int    `json:"created_at"`
						Pk          int64  `json:"pk"`
						User        struct {
							Username      string `json:"username"`
							IsPrivate     bool   `json:"is_private"`
							ProfilePicURL string `json:"profile_pic_url"`
							Pk            int    `json:"pk"`
							ProfilePicID  string `json:"profile_pic_id"`
							IsVerified    bool   `json:"is_verified"`
							FullName      string `json:"full_name"`
						} `json:"user"`
						Text         string `json:"text"`
						CreatedAtUtc int    `json:"created_at_utc"`
					} `json:"preview_comments"`
					CommentCount int `json:"comment_count"`
					Caption      struct {
						MediaID     int64  `json:"media_id"`
						BitFlags    int    `json:"bit_flags"`
						Type        int    `json:"type"`
						Status      string `json:"status"`
						ContentType string `json:"content_type"`
						UserID      int    `json:"user_id"`
						CreatedAt   int    `json:"created_at"`
						Pk          int64  `json:"pk"`
						User        struct {
							Username         string `json:"username"`
							IsUnpublished    bool   `json:"is_unpublished"`
							IsPrivate        bool   `json:"is_private"`
							FriendshipStatus struct {
								Following       bool `json:"following"`
								OutgoingRequest bool `json:"outgoing_request"`
							} `json:"friendship_status"`
							ProfilePicURL              string `json:"profile_pic_url"`
							HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
							IsFavorite                 bool   `json:"is_favorite"`
							Pk                         int    `json:"pk"`
							ProfilePicID               string `json:"profile_pic_id"`
							FullName                   string `json:"full_name"`
						} `json:"user"`
						Text         string `json:"text"`
						CreatedAtUtc int    `json:"created_at_utc"`
					} `json:"caption"`
					CaptionIsEdited bool `json:"caption_is_edited"`
					PhotoOfYou      bool `json:"photo_of_you"`
					VideoVersions   []struct {
						Width  int    `json:"width"`
						Height int    `json:"height"`
						Type   int    `json:"type"`
						URL    string `json:"url"`
					} `json:"video_versions"`
					HasAudio      bool    `json:"has_audio"`
					VideoDuration float64 `json:"video_duration"`
				} `json:"media_share"`
				ItemType string `json:"item_type"`
				UserID   int    `json:"user_id"`
			} `json:"items"`
			Muted     bool `json:"muted"`
			Pending   bool `json:"pending"`
			HasOlder  bool `json:"has_older"`
			Canonical bool `json:"canonical"`
			HasNewer  bool `json:"has_newer"`
			Named     bool `json:"named"`
			Users     []struct {
				Username         string `json:"username"`
				IsPrivate        bool   `json:"is_private"`
				FriendshipStatus struct {
					IsPrivate       bool `json:"is_private"`
					OutgoingRequest bool `json:"outgoing_request"`
					Following       bool `json:"following"`
					Blocking        bool `json:"blocking"`
					IncomingRequest bool `json:"incoming_request"`
				} `json:"friendship_status"`
				ProfilePicURL              string `json:"profile_pic_url"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				Pk                         int    `json:"pk"`
				ProfilePicID               string `json:"profile_pic_id"`
				IsVerified                 bool   `json:"is_verified"`
				FullName                   string `json:"full_name"`
			} `json:"users"`
			IsSpam       bool   `json:"is_spam"`
			NewestCursor string `json:"newest_cursor"`
		} `json:"threads"`
	} `json:"inbox"`
	PendingRequestsUsers []interface{} `json:"pending_requests_users"`
}

type FollowingRecentActivityResponse struct {
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	NextMaxID           int    `json:"next_max_id"`
	Status              string `json:"status"`
	Stories             []struct {
		Pk     string `json:"pk"`
		Counts struct {
		} `json:"counts"`
		Type int `json:"type"`
		Args struct {
			Media []struct {
				Image string `json:"image"`
				ID    string `json:"id"`
			} `json:"media"`
			Text         string `json:"text"`
			CommentID    int64  `json:"comment_id"`
			ProfileImage string `json:"profile_image"`
			Timestamp    int    `json:"timestamp"`
			Links        []struct {
				Start int    `json:"start"`
				ID    string `json:"id"`
				End   int    `json:"end"`
				Type  string `json:"type"`
			} `json:"links"`
			ProfileID int64 `json:"profile_id"`
		} `json:"args"`
	} `json:"stories"`
}
