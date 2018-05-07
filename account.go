package goinsta

// Account is personal account object
type Account struct {
	inst *Instagram
}

// NewAccount creates new account structure
func NewAccount(inst *Instagram) (*Account, error) {
	account := &Account{
		inst: inst,
	}
	return account, err
}
