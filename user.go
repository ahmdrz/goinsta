package goinsta

type User struct {
	 Json objects and user data
  Username                   string `json:"username"`
  Biography                  string `json:"biography"`
  HasAnonymousProfilePicture bool   `json:"has_anonymouse_profile_picture"`
  ProfilePictureID           string `json:"profile_pic_id"`
  ProfilePictureURL          string `json:"profile_pic_url"`
  FullName                   string `json:"full_name"`
  ID                         int64  `json:"pk"`
  IDStr                      string `json:"-"`
  IsVerified                 bool   `json:"is_verified"`
  IsPrivate                  bool   `json:"is_private"`
  IsFavorite                 bool   `json:"is_favorite"`
  IsUnpublished              bool   `json:"is_unpublished"`
  IsBusiness                 bool   `json:"is_business"`
  ExternalLynxURL            string `json:"external_lynx_url"`
  MediaCount                 int    `json:"media_count"`
  AutoExpandChaining         bool   `json:"auto_expand_chaining"`
  FollowingCount             int    `json:"following_count"`
  FollowerCount              int    `json:"follower_count"`
  ExternalURL                string `json:"external_url"`
  HdProfilePicVersions       []struct {
    Height int    `json:"height"`
    Width  int    `json:"width"`
    URL    string `json:"url"`
  } `json:"hd_profile_pic_versions"`
  UserTagsCount       int `json:"usertags_count"`
  HdProfilePicURLInfo struct {
    Height int    `json:"height"`
    Width  int    `json:"width"`
    URL    string `json:"url"`
  } `json:"hd_profile_pic_url_info"`
  GeoMediaCount int  `json:"geo_media_count"`
  HasChaining   bool `json:"has_chaining"`
}
