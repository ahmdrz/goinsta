package goinsta

// Account is personal account object
type Account struct {
	inst *Instagram

	// Activity is recent activity
	//Activity *Activity
	// Tray is your disponible friend's stories
	//Tray *Tray

	// Account is also a User
	User
}

// NewAccount creates new account structure
func NewAccount(inst *Instagram) *Account {
	account := &Account{
		inst: inst,
	}
	return account
}
