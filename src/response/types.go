package goinstaresponse

// UsersResponse
type UsersReponse struct {
	Status    string `json:"status"`
	BigList   bool   `json:"big_list"`
	Users     []User `json:"users"`
	PageSize  int    `json:"page_size"`
	NextMaxID string `json:"next_max_id"`
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
