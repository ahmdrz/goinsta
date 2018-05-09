package goinsta

type accountResp struct {
	Status  string  `json:"status"`
	Account Account `json:"user"`
}

// Account is personal account object
type Account struct {
	// Activity is recent activity
	//Activity *Activity
	// Tray is your disponible friend's stories
	//Tray *Tray

	inst *Instagram

	// User values shared between User and Account
	ID            int64  `json:"pk"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	Biography     string `json:"biography"`
	ProfilePicURL string `json:"profile_pic_url"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	IsBusiness    bool   `json:"is_business"`
	Gender        int    `json:"gender"`

	ProfilePicID               string `json:"profile_pic_id"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool   `json:"is_private"`
	IsVerified                 bool   `json:"is_verified"`
	MediaCount                 int    `json:"media_count"`
	FollowerCount              int    `json:"follower_count"`
	FollowingCount             int    `json:"following_count"`
	GeoMediaCount              int    `json:"geo_media_count"`
	ExternalURL                string `json:"external_url"`
}

// NewAccount creates new account structure
func NewAccount(inst *Instagram) *Account {
	account := &Account{
		inst: inst,
	}
	return account
}
