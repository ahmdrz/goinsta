package goinsta

type accountResp struct {
	Status  string  `json:"status"`
	Account Account `json:"logged_in_user"`
}

// Account is personal account object
type Account struct {
	// Activity is recent activity
	//Activity *Activity
	// Tray is your disponible friend's stories
	//Tray *Tray

	inst *Instagram

	CanSeeOrganicInsights      bool    `json:"can_see_organic_insights"`
	ShowInsightsTerms          bool    `json:"show_insights_terms"`
	IsBusiness                 bool    `json:"is_business"`
	Nametag                    Nametag `json:"nametag"`
	ID                         int64   `json:"pk"`
	Username                   string  `json:"username"`
	FullName                   string  `json:"full_name"`
	HasAnonymousProfilePicture bool    `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool    `json:"is_private"`
	IsVerified                 bool    `json:"is_verified"`
	ProfilePicURL              string  `json:"profile_pic_url"`
	ProfilePicID               string  `json:"profile_pic_id"`
	AllowedCommenterType       string  `json:"allowed_commenter_type"`
	ReelAutoArchive            string  `json:"reel_auto_archive"`
	AllowContactsSync          bool    `json:"allow_contacts_sync"`
	PhoneNumber                string  `json:"phone_number"`
}

// NewAccount creates new account structure
func NewAccount(inst *Instagram) *Account {
	account := &Account{
		inst: inst,
	}
	return account
}
