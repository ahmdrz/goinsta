package goinstaresponse

import (
	"strconv"
)

// Status struct point to if response is ok or not
type StatusResponse struct {
	Status string `json:"status"`
}

// Pagination every pagination have next_max_id
type Int64Pagination struct {
	NextMaxID int64 `json:"next_max_id"`
}

// Pagination every pagination have next_max_id
type StringPagination struct {
	NextMaxID string `json:"next_max_id"`
}

// UsersResponse
type UsersReponse struct {
	StatusResponse
	BigList  bool   `json:"big_list"`
	Users    []User `json:"users"`
	PageSize int    `json:"page_size"`
	StringPagination
}

// User , Instagram user informations
type User struct {
	Username                   string `json:"username"`
	HasAnanymousProfilePicture bool   `json:"has_anonymouse_profile_picture"`
	ProfilePictureId           string `json:"profile_pic_id"`
	ProfilePictureURL          string `json:"profile_pic_url"`
	FullName                   string `json:"full_name"`
	PK                         int64  `json:"pk"`
	IsVerified                 bool   `json:"is_verified"`
	IsPrivate                  bool   `json:"is_private"`
	IsFavorite                 bool   `json:"is_favorite"`
}

// User.String return PK as string
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
	TopLikers                    []string          `json:"top_likers",omitempty`
	HasLiked                     bool              `json:"has_liked"`
	HasMoreComments              bool              `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int               `json:"max_num_visible_preview_comments"`
	PreviewComments              []CommentResponse `json:"preview_comments",omitempty`
	Comments                     []CommentResponse `json:"comments",omitempty`
	CommentCount                 int               `json:"comment_count"`
	Caption                      Caption           `json:"caption",omitempty`
	CaptionIsEdited              bool              `json:"caption_is_edited"`
	PhotoOfYou                   bool              `json:"photo_of_you"`
	Int64Pagination
}

// ImageVersions struct for image information , urls and etc
type ImageVersions struct {
	Condidates []ImageCondidate `json:"condidates"`
}

// ImageCondidate have urls and image width , height
type ImageCondidate struct {
	Url    string `json:"url"`
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
	City             string  `json:"city",omitempty`
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
// Gender -> 1 male , 2 female , 3 unknow
type ProfileUserResponse struct {
	User
	//Birthday -> what the hell is ?
	PhoneNumber             string           `json:"phone_number"`
	HDProfilePicVersions    []ImageCondidate `json:"hd_profile_pic_versions"`
	Gender                  int              `json:"gender"`
	ShowConversionEditEntry bool             `json:"show_conversion_edit_entry"`
	ExternalLynxUrl         string           `json:"external_lynx_url"`
	Biography               string           `json:"biography"`
	HDProfilePicUrlInfo     ImageCondidate   `json:"hd_profile_pic_url_info"`
	Email                   string           `json:"email"`
	ExternalUrl             string           `json:"external_url"`
}

// ProfileDataResponse have StatusResponse and ProfileUserResponse
type ProfileDataResponse struct {
	StatusResponse
	User ProfileUserResponse `json:"user"`
}

// GetUsernameResponse return special userinformation
type GetUsernameResponse struct {
	StatusResponse
	User UsernameResponse `json:"user"`
}

// UsernameResponse information of each instagram users
type UsernameResponse struct {
	User
	ExternalUrl         string         `json:"external_url"`
	Biography           string         `json:"biography"`
	HDProfilePicUrlInfo ImageCondidate `json:"hd_profile_pic_url_info"`
	UserTagsCount       int            `json:"usertags_count"`
	MediaCount          int            `json:"media_count"`
	FollowingCount      int            `json:"following_count"`
	IsBusiness          bool           `json:"is_business"`
	AutoExpandChaining  bool           `json:"auto_expand_chaining"`
	HasChaining         bool           `json:"has_chaining"`
	FollowerCount       int            `json:"follower_count"`
	GeoMediaCount       int            `json:"geo_media_count"`
}

// UploadResponse struct informatio of upload method
type UploadResponse struct {
	StatusResponse
	UploadID string `json:"upload_id",omitempty`
	Message  string `json:"message"`
}

type UploadPhotoResponse struct {
	StatusResponse
	Media    MediaItemResponse `json:"media"`
	UploadID string            `json:"upload_id"`
}

type FriendShipResponse struct {
	IncomingRequest bool `json:"incoming_request"`
	FollowedBy      bool `json:"followed_by"`
	OutgoingRequest bool `json:"outgoing_request"`
	Following       bool `json:"following"`
	Blocking        bool `json:"blocking"`
	IsPrivate       bool `json:"is_private"`
}

type FollowResponse struct {
	StatusResponse
	FriendShipStatus FriendShipResponse `json:"friendship_status"`
}

type UnFollowResponse struct {
	StatusResponse
	FriendShipStatus FriendShipResponse `json:"friendship_status"`
}

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
			ViewerID         int64  `json:"viewer_id"`
			MoreAvailableMin bool   `json:"more_available_min"`
			ThreadID         string `json:"thread_id"`
			ImageVersions2   struct {
				Candidates []struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"candidates"`
			} `json:"image_versions2"`
			LastActivityAt int64         `json:"last_activity_at"`
			NextMaxID      string        `json:"next_max_id"`
			IsSpam         bool          `json:"is_spam"`
			LeftUsers      []interface{} `json:"left_users"`
			NextMinID      string        `json:"next_min_id"`
			Muted          bool          `json:"muted"`
			Items          []struct {
				ItemID     string `json:"item_id"`
				ItemType   string `json:"item_type"`
				MediaShare struct {
					TakenAt         int    `json:"taken_at"`
					Pk              int64  `json:"pk"`
					ID              string `json:"id"`
					DeviceTimestamp int64  `json:"device_timestamp"`
					MediaType       int    `json:"media_type"`
					Code            string `json:"code"`
					ClientCacheKey  string `json:"client_cache_key"`
					FilterType      int    `json:"filter_type"`
					ImageVersions2  struct {
						Candidates []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"candidates"`
					} `json:"image_versions2"`
					OriginalWidth  int     `json:"original_width"`
					OriginalHeight int     `json:"original_height"`
					ViewCount      float64 `json:"view_count"`
					User           struct {
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
		ViewerID         int64  `json:"viewer_id"`
		MoreAvailableMin bool   `json:"more_available_min"`
		ThreadID         string `json:"thread_id"`
		ImageVersions2   struct {
			Candidates []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		LastActivityAt int64         `json:"last_activity_at"`
		NextMaxID      string        `json:"next_max_id"`
		Canonical      bool          `json:"canonical"`
		LeftUsers      []interface{} `json:"left_users"`
		NextMinID      string        `json:"next_min_id"`
		Muted          bool          `json:"muted"`
		Items          []struct {
			ItemID     string `json:"item_id"`
			ItemType   string `json:"item_type"`
			MediaShare struct {
				TakenAt         int    `json:"taken_at"`
				Pk              int64  `json:"pk"`
				ID              string `json:"id"`
				DeviceTimestamp int    `json:"device_timestamp"`
				MediaType       int    `json:"media_type"`
				Code            string `json:"code"`
				ClientCacheKey  string `json:"client_cache_key"`
				FilterType      int    `json:"filter_type"`
				ImageVersions2  struct {
					Candidates []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"candidates"`
				} `json:"image_versions2"`
				OriginalWidth  int     `json:"original_width"`
				OriginalHeight int     `json:"original_height"`
				ViewCount      float64 `json:"view_count"`
				User           struct {
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

type UserFeedResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []struct {
		TakenAt         int64  `json:"taken_at"`
		Pk              int64  `json:"pk"`
		ID              string `json:"id"`
		DeviceTimestamp int64  `json:"device_timestamp"`
		MediaType       int    `json:"media_type"`
		Code            string `json:"code"`
		ClientCacheKey  string `json:"client_cache_key"`
		FilterType      int    `json:"filter_type"`
		ImageVersions2  struct {
			Candidates []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		OriginalWidth  int `json:"original_width"`
		OriginalHeight int `json:"original_height"`
		User           struct {
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
		Usertags        struct {
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
