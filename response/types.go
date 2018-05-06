package response

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
	ID                         int64  `json:"pk"`
	IsVerified                 bool   `json:"is_verified"`
	IsPrivate                  bool   `json:"is_private"`
	IsFavorite                 bool   `json:"is_favorite"`
	IsUnpublished              bool   `json:"is_unpublished"`
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
		ID   int64  `json:"id"`
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
	Pk                           int64             `json:"pk"`
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
	Story               struct {
		ID              string      `json:"id"`
		LatestReelMedia int         `json:"latest_reel_media"`
		ExpiringAt      int         `json:"expiring_at"`
		Seen            interface{} `json:"seen"`
		CanReply        bool        `json:"can_reply"`
		CanReshare      bool        `json:"can_reshare"`
		ReelType        string      `json:"reel_type"`
		Owner           struct {
			Type               string  `json:"type"`
			Pk                 int     `json:"pk"`
			Name               string  `json:"name"`
			ProfilePicURL      string  `json:"profile_pic_url"`
			ProfilePicUsername string  `json:"profile_pic_username"`
			ShortName          string  `json:"short_name"`
			Lng                float64 `json:"lng"`
			Lat                float64 `json:"lat"`
			LocationDict       struct {
				Pk               int     `json:"pk"`
				Name             string  `json:"name"`
				Address          string  `json:"address"`
				City             string  `json:"city"`
				ShortName        string  `json:"short_name"`
				Lng              float64 `json:"lng"`
				Lat              float64 `json:"lat"`
				ExternalSource   string  `json:"external_source"`
				FacebookPlacesID int64   `json:"facebook_places_id"`
			} `json:"location_dict"`
		} `json:"owner"`
		Location struct {
			Type               string  `json:"type"`
			Pk                 int     `json:"pk"`
			Name               string  `json:"name"`
			ProfilePicURL      string  `json:"profile_pic_url"`
			ProfilePicUsername string  `json:"profile_pic_username"`
			ShortName          string  `json:"short_name"`
			Lng                float64 `json:"lng"`
			Lat                float64 `json:"lat"`
			LocationDict       struct {
				Pk               int     `json:"pk"`
				Name             string  `json:"name"`
				Address          string  `json:"address"`
				City             string  `json:"city"`
				ShortName        string  `json:"short_name"`
				Lng              float64 `json:"lng"`
				Lat              float64 `json:"lat"`
				ExternalSource   string  `json:"external_source"`
				FacebookPlacesID int64   `json:"facebook_places_id"`
			} `json:"location_dict"`
		} `json:"location"`
		Items []struct {
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
					Width  int    `json:"width"`
					Height int    `json:"height"`
					URL    string `json:"url"`
				} `json:"candidates"`
			} `json:"image_versions2"`
			OriginalWidth   int     `json:"original_width"`
			OriginalHeight  int     `json:"original_height"`
			CaptionPosition float64 `json:"caption_position"`
			IsReelMedia     bool    `json:"is_reel_media"`
			VideoVersions   []struct {
				Type   int    `json:"type"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
				ID     string `json:"id"`
			} `json:"video_versions"`
			HasAudio      bool    `json:"has_audio"`
			VideoDuration float64 `json:"video_duration"`
			User          struct {
				Pk                         int    `json:"pk"`
				Username                   string `json:"username"`
				FullName                   string `json:"full_name"`
				IsPrivate                  bool   `json:"is_private"`
				ProfilePicURL              string `json:"profile_pic_url"`
				ProfilePicID               string `json:"profile_pic_id"`
				IsVerified                 bool   `json:"is_verified"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsFavorite                 bool   `json:"is_favorite"`
			} `json:"user"`
			Caption              interface{}   `json:"caption"`
			CaptionIsEdited      bool          `json:"caption_is_edited"`
			PhotoOfYou           bool          `json:"photo_of_you"`
			CanViewerSave        bool          `json:"can_viewer_save"`
			OrganicTrackingToken string        `json:"organic_tracking_token"`
			ExpiringAt           int           `json:"expiring_at"`
			ImportedTakenAt      int           `json:"imported_taken_at"`
			IsDashEligible       int           `json:"is_dash_eligible"`
			VideoDashManifest    string        `json:"video_dash_manifest"`
			NumberOfQualities    int           `json:"number_of_qualities"`
			ReelMentions         []interface{} `json:"reel_mentions"`
			StoryLocations       []struct {
				X        float64 `json:"x"`
				Y        float64 `json:"y"`
				Z        int     `json:"z"`
				Width    float64 `json:"width"`
				Height   float64 `json:"height"`
				Rotation float64 `json:"rotation"`
				IsPinned int     `json:"is_pinned"`
				Location struct {
					Pk               int     `json:"pk"`
					Name             string  `json:"name"`
					Address          string  `json:"address"`
					City             string  `json:"city"`
					ShortName        string  `json:"short_name"`
					Lng              float64 `json:"lng"`
					Lat              float64 `json:"lat"`
					ExternalSource   string  `json:"external_source"`
					FacebookPlacesID int64   `json:"facebook_places_id"`
				} `json:"location"`
			} `json:"story_locations"`
			StoryEvents   []interface{} `json:"story_events"`
			StoryHashtags []struct {
				X        float64 `json:"x"`
				Y        float64 `json:"y"`
				Z        int     `json:"z"`
				Width    float64 `json:"width"`
				Height   float64 `json:"height"`
				Rotation float64 `json:"rotation"`
				IsPinned int     `json:"is_pinned"`
				Hashtag  struct {
					Name string `json:"name"`
					ID   int64  `json:"id"`
				} `json:"hashtag"`
			} `json:"story_hashtags"`
			StoryPolls            []interface{} `json:"story_polls"`
			StoryFeedMedia        []interface{} `json:"story_feed_media"`
			StorySoundOn          []interface{} `json:"story_sound_on"`
			StorySliders          []interface{} `json:"story_sliders"`
			CanReshare            bool          `json:"can_reshare"`
			CreativeConfig        interface{}   `json:"creative_config"`
			SupportsReelReactions bool          `json:"supports_reel_reactions"`
		} `json:"items"`
		PrefetchCount       int   `json:"prefetch_count"`
		UniqueIntegerReelID int64 `json:"unique_integer_reel_id"`
	} `json:"story"`
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
	Pk               int64   `json:"pk"`
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
	Pk           int64  `json:"pk"`
	Type         int    `json:"type"`
}

// MediaCommentsResponse struct for get array of comments of a media
type MediaCommentsResponse struct {
	StatusResponse
	StringPagination
	CommentLikesEnabled bool              `json:"comment_likes_enabled"`
	Comments            []CommentResponse `json:"comments"`
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
		ID                   int64  `json:"pk"`
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
				User
				FriendshipStatus struct {
					Following       bool `json:"following"`
					IncomingRequest bool `json:"incoming_request"`
					OutgoingRequest bool `json:"outgoing_request"`
					Blocking        bool `json:"blocking"`
					IsPrivate       bool `json:"is_private"`
				} `json:"friendship_status"`
			} `json:"users"`
			ViewerID         int64            `json:"viewer_id"`
			MoreAvailableMin bool             `json:"more_available_min"`
			ThreadID         string           `json:"thread_id"`
			ImageVersions2   ImageVersions    `json:"image_versions2"`
			LastActivityAt   int64            `json:"last_activity_at"`
			NextMaxID        string           `json:"next_max_id"`
			IsSpam           bool             `json:"is_spam"`
			LeftUsers        []interface{}    `json:"left_users"`
			NextMinID        string           `json:"next_min_id"`
			Muted            bool             `json:"muted"`
			Items            []ItemMediaShare `json:"items"`
			ThreadType       string           `json:"thread_type"`
			MoreAvailableMax bool             `json:"more_available_max"`
			ThreadTitle      string           `json:"thread_title"`
			Canonical        bool             `json:"canonical"`
			Inviter          User             `json:"inviter"`
			Pending          bool             `json:"pending"`
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
				Pk                         int64  `json:"pk"`
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
			Pk            int64  `json:"pk"`
			IsVerified    bool   `json:"is_verified"`
			IsPrivate     bool   `json:"is_private"`
		} `json:"users"`
		ViewerID         int64            `json:"viewer_id"`
		MoreAvailableMin bool             `json:"more_available_min"`
		ThreadID         string           `json:"thread_id"`
		ImageVersions2   ImageVersions    `json:"image_versions2"`
		LastActivityAt   int64            `json:"last_activity_at"`
		NextMaxID        string           `json:"next_max_id"`
		Canonical        bool             `json:"canonical"`
		LeftUsers        []interface{}    `json:"left_users"`
		NextMinID        string           `json:"next_min_id"`
		Muted            bool             `json:"muted"`
		Items            []ItemMediaShare `json:"items"`
		ThreadType       string           `json:"thread_type"`
		MoreAvailableMax bool             `json:"more_available_max"`
		ThreadTitle      string           `json:"thread_title"`
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
			Pk                         int64  `json:"pk"`
			IsVerified                 bool   `json:"is_verified"`
			IsPrivate                  bool   `json:"is_private"`
		} `json:"inviter"`
		Pending bool `json:"pending"`
	} `json:"thread"`
}

type UserTaggedFeedResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []Item `json:"items"`
	MoreAvailable       bool   `json:"more_available"`
	NextMaxID           int64  `json:"next_max_id"`
	TotalCount          int    `json:"total_count"`
	RequiresReview      bool   `json:"requires_review"`
	// maybe this is photos waiting for review?
	// NewPhotos           []interface{} `json:"new_photos"`
}

// UserFeedResponse contains user feeds
type UserFeedResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []Item `json:"items"`
	MoreAvailable       bool   `json:"more_available"`
	NextMaxID           string `json:"next_max_id"`
}

// Item user feeds item
type Item struct {
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
		Pk                         int64  `json:"pk"`
		IsVerified                 bool   `json:"is_verified"`
		IsPrivate                  bool   `json:"is_private"`
	} `json:"user"`
	CarouselMedia []struct {
		ID            string        `json:"id"`
		MediaType     int           `json:"media_type"`
		ImageVersions ImageVersions `json:"image_versions2"`
		VideoVersions []struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Type   int    `json:"type"`
			Height int    `json:"height"`
		} `json:"video_versions"`
		OriginalWidth    int    `json:"original_width"`
		OriginalHeight   int    `json:"original_height"`
		Pk               int64  `json:"pk"`
		CarouselParentID string `json:"carousel_parent_id"`
	} `json:"carousel_media"`
	OrganicTrackingToken         string            `json:"organic_tracking_token"`
	LikeCount                    int               `json:"like_count"`
	TopLikers                    []interface{}     `json:"top_likers"`
	HasLiked                     bool              `json:"has_liked"`
	HasMoreComments              bool              `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int               `json:"max_num_visible_preview_comments"`
	PreviewComments              []CommentResponse `json:"preview_comments"`
	Comments                     []CommentResponse `json:"comments"`
	CommentCount                 int               `json:"comment_count"`
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
			Pk                         int64  `json:"pk"`
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
			Pk            int64  `json:"pk"`
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

type SearchTagsResponse struct {
	Results []struct {
		Name       string `json:"name"`
		MediaCount int    `json:"media_count"`
		ID         int64  `json:"id"`
	} `json:"results"`
	HasMore   interface{} `json:"has_more"`
	RankToken string      `json:"rank_token"`
	Status    string      `json:"status"`
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
		Pk                   int64   `json:"pk"`
		IsVerified           bool    `json:"is_verified"`
		IsPrivate            bool    `json:"is_private"`
		ProfilePicID         string  `json:"profile_pic_id,omitempty"`
	} `json:"users"`
}

// ExploreResponse is data from explore in Instagram
type ExploreResponse struct {
	AutoLoadMoreEnabled bool          `json:"auto_load_more_enabled"`
	Items               []ExploreItem `json:"items"`
	MaxID               string        `json:"max_id"`
	MoreAvailable       bool          `json:"more_available"`
	NextMaxID           string        `json:"next_max_id"`
	NumResults          int64         `json:"num_results"`
	RankToken           string        `json:"rank_token"`
	Status              string        `json:"status"`
}

// ExploreItem contain information for one exploree feed
type ExploreItem struct {
	Channel struct {
		ChannelID   string `json:"channel_id"`
		ChannelType string `json:"channel_type"`
		Context     string `json:"context"`
		Header      string `json:"header"`
		Media       struct {
			Algorithm     string `json:"algorithm"`
			CanViewerSave bool   `json:"can_viewer_save"`
			Caption       struct {
				BitFlags        int64  `json:"bit_flags"`
				ContentType     string `json:"content_type"`
				CreatedAt       int64  `json:"created_at"`
				CreatedAtUtc    int64  `json:"created_at_utc"`
				DidReportAsSpam bool   `json:"did_report_as_spam"`
				HasTranslation  bool   `json:"has_translation"`
				MediaID         int64  `json:"media_id"`
				Pk              int64  `json:"pk"`
				Status          string `json:"status"`
				Text            string `json:"text"`
				Type            int64  `json:"type"`
				User            struct {
					FriendshipStatus struct {
						Following       bool `json:"following"`
						IsBestie        bool `json:"is_bestie"`
						OutgoingRequest bool `json:"outgoing_request"`
					} `json:"friendship_status"`
					FullName                   string `json:"full_name"`
					HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
					IsFavorite                 bool   `json:"is_favorite"`
					IsPrivate                  bool   `json:"is_private"`
					IsUnpublished              bool   `json:"is_unpublished"`
					IsVerified                 bool   `json:"is_verified"`
					LatestReelMedia            int64  `json:"latest_reel_media"`
					Pk                         int64  `json:"pk"`
					ProfilePicID               string `json:"profile_pic_id"`
					ProfilePicURL              string `json:"profile_pic_url"`
					Username                   string `json:"username"`
				} `json:"user"`
				UserID int64 `json:"user_id"`
			} `json:"caption"`
			CaptionIsEdited         bool   `json:"caption_is_edited"`
			ClientCacheKey          string `json:"client_cache_key"`
			Code                    string `json:"code"`
			CommentCount            int64  `json:"comment_count"`
			CommentLikesEnabled     bool   `json:"comment_likes_enabled"`
			CommentThreadingEnabled bool   `json:"comment_threading_enabled"`
			DeviceTimestamp         int64  `json:"device_timestamp"`
			Explore                 struct {
				ActorID     int64  `json:"actor_id"`
				Explanation string `json:"explanation"`
				SourceToken string `json:"source_token"`
			} `json:"explore"`
			ExploreContext     string `json:"explore_context"`
			ExploreSourceToken string `json:"explore_source_token"`
			FilterType         int64  `json:"filter_type"`
			HasAudio           bool   `json:"has_audio"`
			HasLiked           bool   `json:"has_liked"`
			HasMoreComments    bool   `json:"has_more_comments"`
			ID                 string `json:"id"`
			ImageVersions2     struct {
				Candidates []struct {
					Height int64  `json:"height"`
					URL    string `json:"url"`
					Width  int64  `json:"width"`
				} `json:"candidates"`
			} `json:"image_versions2"`
			ImpressionToken string  `json:"impression_token"`
			IsDashEligible  int64   `json:"is_dash_eligible"`
			Lat             float64 `json:"lat"`
			LikeCount       int64   `json:"like_count"`
			Lng             float64 `json:"lng"`
			Location        struct {
				Address          string  `json:"address"`
				City             string  `json:"city"`
				ExternalSource   string  `json:"external_source"`
				FacebookPlacesID int64   `json:"facebook_places_id"`
				Lat              float64 `json:"lat"`
				Lng              float64 `json:"lng"`
				Name             string  `json:"name"`
				Pk               int64   `json:"pk"`
				ShortName        string  `json:"short_name"`
			} `json:"location"`
			MaxNumVisiblePreviewComments int64  `json:"max_num_visible_preview_comments"`
			MediaType                    int64  `json:"media_type"`
			NextMaxID                    int64  `json:"next_max_id"`
			NumberOfQualities            int64  `json:"number_of_qualities"`
			OrganicTrackingToken         string `json:"organic_tracking_token"`
			OriginalHeight               int64  `json:"original_height"`
			OriginalWidth                int64  `json:"original_width"`
			PhotoOfYou                   bool   `json:"photo_of_you"`
			Pk                           int64  `json:"pk"`
			PreviewComments              []struct {
				BitFlags        int64  `json:"bit_flags"`
				ContentType     string `json:"content_type"`
				CreatedAt       int64  `json:"created_at"`
				CreatedAtUtc    int64  `json:"created_at_utc"`
				DidReportAsSpam bool   `json:"did_report_as_spam"`
				HasTranslation  bool   `json:"has_translation"`
				MediaID         int64  `json:"media_id"`
				Pk              int64  `json:"pk"`
				Status          string `json:"status"`
				Text            string `json:"text"`
				Type            int64  `json:"type"`
				User            struct {
					FullName      string `json:"full_name"`
					IsPrivate     bool   `json:"is_private"`
					IsVerified    bool   `json:"is_verified"`
					Pk            int64  `json:"pk"`
					ProfilePicID  string `json:"profile_pic_id"`
					ProfilePicURL string `json:"profile_pic_url"`
					Username      string `json:"username"`
				} `json:"user"`
				UserID int64 `json:"user_id"`
			} `json:"preview_comments"`
			TakenAt int64 `json:"taken_at"`
			User    struct {
				FriendshipStatus struct {
					Following       bool `json:"following"`
					IsBestie        bool `json:"is_bestie"`
					OutgoingRequest bool `json:"outgoing_request"`
				} `json:"friendship_status"`
				FullName                   string `json:"full_name"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsFavorite                 bool   `json:"is_favorite"`
				IsPrivate                  bool   `json:"is_private"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsVerified                 bool   `json:"is_verified"`
				LatestReelMedia            int64  `json:"latest_reel_media"`
				Pk                         int64  `json:"pk"`
				ProfilePicID               string `json:"profile_pic_id"`
				ProfilePicURL              string `json:"profile_pic_url"`
				Username                   string `json:"username"`
			} `json:"user"`
			VideoDashManifest string  `json:"video_dash_manifest"`
			VideoDuration     float64 `json:"video_duration"`
			VideoVersions     []struct {
				Height int64  `json:"height"`
				ID     string `json:"id"`
				Type   int64  `json:"type"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"video_versions"`
			ViewCount float64 `json:"view_count"`
		} `json:"media"`
		MediaCount int64  `json:"media_count"`
		Title      string `json:"title"`
	} `json:"channel"`
	ExploreItemInfo struct {
		AspectRatio     float64 `json:"aspect_ratio"`
		Autoplay        bool    `json:"autoplay"`
		NumColumns      int64   `json:"num_columns"`
		TotalNumColumns int64   `json:"total_num_columns"`
	} `json:"explore_item_info"`
	Media struct {
		Algorithm     string `json:"algorithm"`
		CanViewerSave bool   `json:"can_viewer_save"`
		Caption       struct {
			BitFlags        int64  `json:"bit_flags"`
			ContentType     string `json:"content_type"`
			CreatedAt       int64  `json:"created_at"`
			CreatedAtUtc    int64  `json:"created_at_utc"`
			DidReportAsSpam bool   `json:"did_report_as_spam"`
			HasTranslation  bool   `json:"has_translation"`
			MediaID         int64  `json:"media_id"`
			Pk              int64  `json:"pk"`
			Status          string `json:"status"`
			Text            string `json:"text"`
			Type            int64  `json:"type"`
			User            struct {
				FriendshipStatus struct {
					Following       bool `json:"following"`
					IsBestie        bool `json:"is_bestie"`
					OutgoingRequest bool `json:"outgoing_request"`
				} `json:"friendship_status"`
				FullName                   string `json:"full_name"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsFavorite                 bool   `json:"is_favorite"`
				IsPrivate                  bool   `json:"is_private"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsVerified                 bool   `json:"is_verified"`
				LatestReelMedia            int64  `json:"latest_reel_media"`
				Pk                         int64  `json:"pk"`
				ProfilePicID               string `json:"profile_pic_id"`
				ProfilePicURL              string `json:"profile_pic_url"`
				Username                   string `json:"username"`
			} `json:"user"`
			UserID int64 `json:"user_id"`
		} `json:"caption"`
		CaptionIsEdited bool `json:"caption_is_edited"`
		CarouselMedia   []struct {
			CarouselParentID string `json:"carousel_parent_id"`
			ID               string `json:"id"`
			ImageVersions2   struct {
				Candidates []struct {
					Height int64  `json:"height"`
					URL    string `json:"url"`
					Width  int64  `json:"width"`
				} `json:"candidates"`
			} `json:"image_versions2"`
			MediaType      int64 `json:"media_type"`
			OriginalHeight int64 `json:"original_height"`
			OriginalWidth  int64 `json:"original_width"`
			Pk             int64 `json:"pk"`
		} `json:"carousel_media"`
		ClientCacheKey          string `json:"client_cache_key"`
		Code                    string `json:"code"`
		CommentCount            int64  `json:"comment_count"`
		CommentLikesEnabled     bool   `json:"comment_likes_enabled"`
		CommentThreadingEnabled bool   `json:"comment_threading_enabled"`
		DeviceTimestamp         int64  `json:"device_timestamp"`
		Explore                 struct {
			ActorID     int64  `json:"actor_id"`
			Explanation string `json:"explanation"`
			SourceToken string `json:"source_token"`
		} `json:"explore"`
		ExploreContext     string `json:"explore_context"`
		ExploreSourceToken string `json:"explore_source_token"`
		FilterType         int64  `json:"filter_type"`
		HasAudio           bool   `json:"has_audio"`
		HasLiked           bool   `json:"has_liked"`
		HasMoreComments    bool   `json:"has_more_comments"`
		ID                 string `json:"id"`
		ImageVersions2     struct {
			Candidates []struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		ImpressionToken string  `json:"impression_token"`
		IsDashEligible  int64   `json:"is_dash_eligible"`
		Lat             float64 `json:"lat"`
		LikeCount       int64   `json:"like_count"`
		Likers          []struct {
			FullName      string `json:"full_name"`
			IsPrivate     bool   `json:"is_private"`
			IsVerified    bool   `json:"is_verified"`
			Pk            int64  `json:"pk"`
			ProfilePicID  string `json:"profile_pic_id"`
			ProfilePicURL string `json:"profile_pic_url"`
			Username      string `json:"username"`
		} `json:"likers"`
		Lng      float64 `json:"lng"`
		Location struct {
			Address          string  `json:"address"`
			City             string  `json:"city"`
			ExternalSource   string  `json:"external_source"`
			FacebookPlacesID int64   `json:"facebook_places_id"`
			Lat              float64 `json:"lat"`
			Lng              float64 `json:"lng"`
			Name             string  `json:"name"`
			Pk               int64   `json:"pk"`
			ShortName        string  `json:"short_name"`
		} `json:"location"`
		MaxNumVisiblePreviewComments int64  `json:"max_num_visible_preview_comments"`
		MediaType                    int64  `json:"media_type"`
		NextMaxID                    int64  `json:"next_max_id"`
		NumberOfQualities            int64  `json:"number_of_qualities"`
		OrganicTrackingToken         string `json:"organic_tracking_token"`
		OriginalHeight               int64  `json:"original_height"`
		OriginalWidth                int64  `json:"original_width"`
		PhotoOfYou                   bool   `json:"photo_of_you"`
		Pk                           int64  `json:"pk"`
		PreviewComments              []struct {
			BitFlags        int64  `json:"bit_flags"`
			ContentType     string `json:"content_type"`
			CreatedAt       int64  `json:"created_at"`
			CreatedAtUtc    int64  `json:"created_at_utc"`
			DidReportAsSpam bool   `json:"did_report_as_spam"`
			HasTranslation  bool   `json:"has_translation"`
			MediaID         int64  `json:"media_id"`
			Pk              int64  `json:"pk"`
			Status          string `json:"status"`
			Text            string `json:"text"`
			Type            int64  `json:"type"`
			User            struct {
				FullName      string `json:"full_name"`
				IsPrivate     bool   `json:"is_private"`
				IsVerified    bool   `json:"is_verified"`
				Pk            int64  `json:"pk"`
				ProfilePicID  string `json:"profile_pic_id"`
				ProfilePicURL string `json:"profile_pic_url"`
				Username      string `json:"username"`
			} `json:"user"`
			UserID int64 `json:"user_id"`
		} `json:"preview_comments"`
		TakenAt   int64    `json:"taken_at"`
		TopLikers []string `json:"top_likers"`
		User      struct {
			FriendshipStatus struct {
				Following       bool `json:"following"`
				IsBestie        bool `json:"is_bestie"`
				OutgoingRequest bool `json:"outgoing_request"`
			} `json:"friendship_status"`
			FullName                   string `json:"full_name"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			IsFavorite                 bool   `json:"is_favorite"`
			IsPrivate                  bool   `json:"is_private"`
			IsUnpublished              bool   `json:"is_unpublished"`
			IsVerified                 bool   `json:"is_verified"`
			LatestReelMedia            int64  `json:"latest_reel_media"`
			Pk                         int64  `json:"pk"`
			ProfilePicID               string `json:"profile_pic_id"`
			ProfilePicURL              string `json:"profile_pic_url"`
			Username                   string `json:"username"`
		} `json:"user"`
		Usertags struct {
			In []struct {
				DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
				Position              []float64   `json:"position"`
				StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
				User                  struct {
					FullName      string `json:"full_name"`
					IsPrivate     bool   `json:"is_private"`
					IsVerified    bool   `json:"is_verified"`
					Pk            int64  `json:"pk"`
					ProfilePicID  string `json:"profile_pic_id"`
					ProfilePicURL string `json:"profile_pic_url"`
					Username      string `json:"username"`
				} `json:"user"`
			} `json:"in"`
		} `json:"usertags"`
		VideoDashManifest string  `json:"video_dash_manifest"`
		VideoDuration     float64 `json:"video_duration"`
		VideoVersions     []struct {
			Height int64  `json:"height"`
			ID     string `json:"id"`
			Type   int64  `json:"type"`
			URL    string `json:"url"`
			Width  int64  `json:"width"`
		} `json:"video_versions"`
		ViewCount float64 `json:"view_count"`
	} `json:"media"`
	Stories struct {
		ID         int64 `json:"id"`
		IsPortrait bool  `json:"is_portrait"`
		Tray       []struct {
			CanReply           bool        `json:"can_reply"`
			CanReshare         bool        `json:"can_reshare"`
			ExpiringAt         int64       `json:"expiring_at"`
			ID                 interface{} `json:"id"`
			LatestReelMedia    int64       `json:"latest_reel_media"`
			PrefetchCount      int64       `json:"prefetch_count"`
			RankedPosition     int64       `json:"ranked_position"`
			ReelType           string      `json:"reel_type"`
			Seen               float64     `json:"seen"`
			SeenRankedPosition int64       `json:"seen_ranked_position"`
			User               struct {
				FriendshipStatus struct {
					Following       bool `json:"following"`
					IncomingRequest bool `json:"incoming_request"`
					IsBestie        bool `json:"is_bestie"`
					IsPrivate       bool `json:"is_private"`
					OutgoingRequest bool `json:"outgoing_request"`
				} `json:"friendship_status"`
				FullName      string `json:"full_name"`
				IsPrivate     bool   `json:"is_private"`
				IsVerified    bool   `json:"is_verified"`
				Pk            int64  `json:"pk"`
				ProfilePicID  string `json:"profile_pic_id"`
				ProfilePicURL string `json:"profile_pic_url"`
				Username      string `json:"username"`
			} `json:"user"`
		} `json:"tray"`
	} `json:"stories"`
}

// MediaInfoResponse contains media information
type MediaInfoResponse struct {
	Status              string `json:"status"`
	NumResults          int    `json:"num_results"`
	AutoLoadMoreEnabled bool   `json:"auto_load_more_enabled"`
	Items               []Item `json:"items"`
	MoreAvailable       bool   `json:"more_available"`
	CommentLikesEnabled bool   `json:"comment_likes_enabled"`
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
	Items               []Item `json:"items"`
	MoreAvailable       bool   `json:"more_available"`
}

type ItemMediaShare struct {
	Item
	MediaShare Item `json:"media_share"`
}

type DirectListResponse struct {
	Inbox struct {
		Threads []struct {
			ThreadID   string `json:"thread_id"`
			ThreadV2ID int64  `json:"thread_v2_id"`
			Users      []struct {
				Pk               int64  `json:"pk"`
				Username         string `json:"username"`
				FullName         string `json:"full_name"`
				IsPrivate        bool   `json:"is_private"`
				ProfilePicURL    string `json:"profile_pic_url"`
				ProfilePicID     string `json:"profile_pic_id"`
				FriendshipStatus struct {
					Following       bool `json:"following"`
					Blocking        bool `json:"blocking"`
					IsPrivate       bool `json:"is_private"`
					IncomingRequest bool `json:"incoming_request"`
					OutgoingRequest bool `json:"outgoing_request"`
					IsBestie        bool `json:"is_bestie"`
				} `json:"friendship_status"`
				IsVerified                 bool `json:"is_verified"`
				HasAnonymousProfilePicture bool `json:"has_anonymous_profile_picture"`
			} `json:"users"`
			LeftUsers []interface{} `json:"left_users"`
			Items     []struct {
				ItemID        string `json:"item_id"`
				UserID        int64  `json:"user_id"`
				Timestamp     int64  `json:"timestamp"`
				ItemType      string `json:"item_type"`
				Text          string `json:"text"`
				ClientContext string `json:"client_context"`
			} `json:"items"`
			LastActivityAt            int64  `json:"last_activity_at"`
			Muted                     bool   `json:"muted"`
			IsPin                     bool   `json:"is_pin"`
			Named                     bool   `json:"named"`
			Canonical                 bool   `json:"canonical"`
			Pending                   bool   `json:"pending"`
			ValuedRequest             bool   `json:"valued_request"`
			ThreadType                string `json:"thread_type"`
			ViewerID                  int64  `json:"viewer_id"`
			ThreadTitle               string `json:"thread_title"`
			PendingScore              int64  `json:"pending_score"`
			ReshareSendCount          int    `json:"reshare_send_count"`
			ReshareReceiveCount       int    `json:"reshare_receive_count"`
			ExpiringMediaSendCount    int    `json:"expiring_media_send_count"`
			ExpiringMediaReceiveCount int    `json:"expiring_media_receive_count"`
			Inviter                   struct {
				Pk                         int64  `json:"pk"`
				Username                   string `json:"username"`
				FullName                   string `json:"full_name"`
				IsPrivate                  bool   `json:"is_private"`
				ProfilePicURL              string `json:"profile_pic_url"`
				ProfilePicID               string `json:"profile_pic_id"`
				IsVerified                 bool   `json:"is_verified"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				ReelAutoArchive            string `json:"reel_auto_archive"`
			} `json:"inviter"`
			HasOlder   bool `json:"has_older"`
			HasNewer   bool `json:"has_newer"`
			LastSeenAt struct {
				Num6107347699 struct {
					Timestamp string `json:"timestamp"`
					ItemID    string `json:"item_id"`
				} `json:"6107347699"`
			} `json:"last_seen_at"`
			NewestCursor string `json:"newest_cursor"`
			OldestCursor string `json:"oldest_cursor"`
			IsSpam       bool   `json:"is_spam"`
		} `json:"threads"`
		HasOlder            bool  `json:"has_older"`
		UnseenCount         int   `json:"unseen_count"`
		UnseenCountTs       int64 `json:"unseen_count_ts"`
		BlendedInboxEnabled bool  `json:"blended_inbox_enabled"`
	} `json:"inbox"`
	SeqID                int    `json:"seq_id"`
	PendingRequestsTotal int    `json:"pending_requests_total"`
	Status               string `json:"status"`
}

type FollowingRecentActivityResponse struct {
	AutoLoadMoreEnabled bool `json:"auto_load_more_enabled"`
	NextMaxID           int  `json:"next_max_id"`
	Stories             []struct {
		Type      int `json:"type"`
		StoryType int `json:"story_type"`
		Args      struct {
			Text  string `json:"text"`
			Links []struct {
				Start int    `json:"start"`
				End   int    `json:"end"`
				Type  string `json:"type"`
				ID    string `json:"id"`
			} `json:"links"`
			ProfileID    int    `json:"profile_id"`
			ProfileImage string `json:"profile_image"`
			Media        []struct {
				ID    string `json:"id"`
				Image string `json:"image"`
			} `json:"media"`
			Timestamp int64  `json:"timestamp"`
			UUID      string `json:"tuuid"`
		} `json:"args"`
		Counts struct {
		} `json:"counts"`
		Pk string `json:"pk"`
	} `json:"stories"`
	Status string `json:"status"`
}

type TrayResponse struct {
	Status string `json:"status"`
	Tray   []struct {
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
		ID                 int `json:"id"`
		LatestReelMedia    int `json:"latest_reel_media"`
		Seen               int `json:"seen"`
		RankedPosition     int `json:"ranked_position"`
		SeenRankedPosition int `json:"seen_ranked_position"`
		Muted              int `json:"muted"`
		Media              []struct {
			TakenAt         int    `json:"taken_at"`
			Pk              int64  `json:"pk"`
			ID              string `json:"id"`
			DeviceTimestamp int64  `json:"device_timestamp"`
			MediaType       int    `json:"media_type"`
			Code            string `json:"code"`
			ClientCacheKey  string `json:"client_cache_key"`
			FilterType      int    `json:"filter_type"`
			ImageVersions2  struct {
				Candidates []ImageCandidate `json:"candidates"`
			} `json:"image_versions2"`
			OriginalWidth  int  `json:"original_width"`
			OriginalHeight int  `json:"original_height"`
			HasAudio       bool `json:"has_audio"`
			VideoVersions  []struct {
				URL    string `json:"url"`
				Type   int    `json:"type"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"video_versions"`
			User struct {
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
			OrganicTrackingToken         string            `json:"organic_tracking_token"`
			LikeCount                    int               `json:"like_count"`
			HasLiked                     bool              `json:"has_liked"`
			HasMoreComments              bool              `json:"has_more_comments"`
			NextMaxID                    int64             `json:"next_max_id"`
			MaxNumVisiblePreviewComments int               `json:"max_num_visible_preview_comments"`
			PreviewComments              []CommentResponse `json:"preview_comments"`
			CommentCount                 int               `json:"comment_count"`
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
		} `json:"items"`
	} `json:"tray"`
}

// TrayUserResponse - Response for specific user tray
type TrayUserResponse struct {
	Status     string `json:"status"`
	CanReply   bool   `json:"can_reply"`
	ExpiringAt int    `json:"expiring_at"`
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
	ID                 int `json:"id"`
	LatestReelMedia    int `json:"latest_reel_media"`
	Seen               int `json:"seen"`
	RankedPosition     int `json:"ranked_position"`
	SeenRankedPosition int `json:"seen_ranked_position"`
	Muted              int `json:"muted"`
	Media              []struct {
		TakenAt         int    `json:"taken_at"`
		Pk              int64  `json:"pk"`
		ID              string `json:"id"`
		DeviceTimestamp int64  `json:"device_timestamp"`
		MediaType       int    `json:"media_type"`
		Code            string `json:"code"`
		ClientCacheKey  string `json:"client_cache_key"`
		FilterType      int    `json:"filter_type"`
		ImageVersions2  struct {
			Candidates []ImageCandidate `json:"candidates"`
		} `json:"image_versions2"`
		OriginalWidth  int  `json:"original_width"`
		OriginalHeight int  `json:"original_height"`
		HasAudio       bool `json:"has_audio"`
		VideoVersions  []struct {
			URL    string `json:"url"`
			Type   int    `json:"type"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"video_versions"`
		User struct {
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
	} `json:"media"`
}

type StoryResponse struct {
	ID              int     `json:"id"`
	LatestReelMedia int     `json:"latest_reel_media"`
	ExpiringAt      int     `json:"expiring_at"`
	Seen            float64 `json:"seen"`
	CanReply        bool    `json:"can_reply"`
	CanReshare      bool    `json:"can_reshare"`
	ReelType        string  `json:"reel_type"`
	User            struct {
		Pk               int    `json:"pk"`
		Username         string `json:"username"`
		FullName         string `json:"full_name"`
		IsPrivate        bool   `json:"is_private"`
		ProfilePicURL    string `json:"profile_pic_url"`
		ProfilePicID     string `json:"profile_pic_id"`
		FriendshipStatus struct {
			Following       bool `json:"following"`
			FollowedBy      bool `json:"followed_by"`
			Blocking        bool `json:"blocking"`
			IsPrivate       bool `json:"is_private"`
			IncomingRequest bool `json:"incoming_request"`
			OutgoingRequest bool `json:"outgoing_request"`
			IsBestie        bool `json:"is_bestie"`
		} `json:"friendship_status"`
		IsVerified bool `json:"is_verified"`
	} `json:"user"`
	Items []struct {
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
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"candidates"`
		} `json:"image_versions2"`
		OriginalWidth   int     `json:"original_width"`
		OriginalHeight  int     `json:"original_height"`
		CaptionPosition float64 `json:"caption_position"`
		IsReelMedia     bool    `json:"is_reel_media"`
		User            struct {
			Pk                         int    `json:"pk"`
			Username                   string `json:"username"`
			FullName                   string `json:"full_name"`
			IsPrivate                  bool   `json:"is_private"`
			ProfilePicURL              string `json:"profile_pic_url"`
			ProfilePicID               string `json:"profile_pic_id"`
			IsVerified                 bool   `json:"is_verified"`
			HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
			IsUnpublished              bool   `json:"is_unpublished"`
			IsFavorite                 bool   `json:"is_favorite"`
		} `json:"user"`
		Caption struct {
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
				Pk                         int    `json:"pk"`
				Username                   string `json:"username"`
				FullName                   string `json:"full_name"`
				IsPrivate                  bool   `json:"is_private"`
				ProfilePicURL              string `json:"profile_pic_url"`
				ProfilePicID               string `json:"profile_pic_id"`
				IsVerified                 bool   `json:"is_verified"`
				HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
				IsUnpublished              bool   `json:"is_unpublished"`
				IsFavorite                 bool   `json:"is_favorite"`
			} `json:"user"`
			DidReportAsSpam bool  `json:"did_report_as_spam"`
			MediaID         int64 `json:"media_id"`
			HasTranslation  bool  `json:"has_translation"`
		} `json:"caption"`
		CaptionIsEdited              bool          `json:"caption_is_edited"`
		LikeCount                    int           `json:"like_count"`
		HasLiked                     bool          `json:"has_liked"`
		Likers                       []interface{} `json:"likers"`
		CommentLikesEnabled          bool          `json:"comment_likes_enabled"`
		CommentThreadingEnabled      bool          `json:"comment_threading_enabled"`
		HasMoreComments              bool          `json:"has_more_comments"`
		MaxNumVisiblePreviewComments int           `json:"max_num_visible_preview_comments"`
		PreviewComments              []interface{} `json:"preview_comments"`
		CommentCount                 int           `json:"comment_count"`
		PhotoOfYou                   bool          `json:"photo_of_you"`
		CanViewerSave                bool          `json:"can_viewer_save"`
		OrganicTrackingToken         string        `json:"organic_tracking_token"`
		ExpiringAt                   int           `json:"expiring_at"`
		ImportedTakenAt              int           `json:"imported_taken_at,omitempty"`
		ReelMentions                 []interface{} `json:"reel_mentions"`
		StoryLocations               []interface{} `json:"story_locations"`
		StoryEvents                  []interface{} `json:"story_events"`
		StoryHashtags                []interface{} `json:"story_hashtags"`
		StoryPolls                   []interface{} `json:"story_polls"`
		StoryFeedMedia               []interface{} `json:"story_feed_media"`
		StorySoundOn                 []interface{} `json:"story_sound_on"`
		StorySliders                 []interface{} `json:"story_sliders"`
		CanReshare                   bool          `json:"can_reshare"`
		CreativeConfig               interface{}   `json:"creative_config"`
		SupportsReelReactions        bool          `json:"supports_reel_reactions"`
		VideoVersions                []struct {
			Type   int    `json:"type"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
			ID     string `json:"id"`
		} `json:"video_versions,omitempty"`
		HasAudio          bool    `json:"has_audio,omitempty"`
		VideoDuration     float64 `json:"video_duration,omitempty"`
		IsDashEligible    int     `json:"is_dash_eligible,omitempty"`
		VideoDashManifest string  `json:"video_dash_manifest,omitempty"`
		NumberOfQualities int     `json:"number_of_qualities,omitempty"`
	} `json:"items"`
	PrefetchCount   int    `json:"prefetch_count"`
	HasBestiesMedia int    `json:"has_besties_media"`
	Status          string `json:"status"`
}

// RecentActivityResponse ...
type RecentActivityResponse struct {
	ContinuationToken int64 `json:"continuation_token"`
	Counts            struct {
		PhotosOfYou int64 `json:"photos_of_you"`
		Requests    int64 `json:"requests"`
	} `json:"counts"`
	Status string `json:"status"` // ok
	//"subscription": null
	//"friend_request_stories": [],
	//"new_stories": [],
	OldStories []struct {
		Args struct {
			Actions      []string `json:"actions,omitempty"`
			CommentID    int64    `json:"comment_id,omitempty"`
			CommentIDs   []int64  `json:"comment_ids,omitempty"`
			InlineFollow struct {
				Following       bool `json:"following,omitempty"`
				OutgoingRequest bool `json:"outgoing_request,omitempty"`
				UserInfo        struct {
					ID            int64  `json:"id,omitempty"`
					IsPrivate     bool   `json:"is_private,omitempty"`
					ProfilePicURL string `json:"profile_pic_url,omitempty"`
					Username      string `json:"username,omitempty"`
				} `json:"user_info,omitempty"`
			} `json:"inline_follow,omitempty"`

			Links []struct {
				End   int64  `json:"end,omitempty"`
				ID    string `json:"id,omitempty"`
				Start int64  `json:"start,omitempty"`
				Type  string `json:"type,omitempty"` // "user" or
			} `json:"links,omitempty"`
			Media []struct {
				ID    string `json:"id,omitempty"`
				Image string `json:"image,omitempty"`
			} `json:"media,omitempty"`
			ProfileID    int64   `json:"profile_id,omitempty"`
			ProfileImage string  `json:"profile_image,omitempty"`
			Text         string  `json:"text,omitempty"`
			Timestamp    float64 `json:"timestamp,omitempty"`
			TUUID        string  `json:"tuuid,omitempty"`
		} `json:"args,omitempty"`

		//"counts": {},
		PK        string `json:"pk"`
		StoryType int64  `json:"story_type"`

		Type int64 `json:"type,omitempty"`
	} `json:"old_stories,omitempty"`
}
