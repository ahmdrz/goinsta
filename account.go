package goinsta

// Account is personal account object
type Account struct {
	inst *Instagram

	// Activity is recent activity
	Activity *Activity
	// Tray is your disponible friend's stories
	Tray *Tray

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

// NewAccount creates new account structure
func NewAccount(inst *Instagram) (*Account, error) {
	account := &Account{
		inst: inst,
	}
	return account, nil
}
