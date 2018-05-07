package goinsta

// Users
type Users struct {
	inst *Instagram
}

// NewUsers
func NewUsers(inst *Instagram) (*Users, error) {
	users := &Users{
		inst: inst,
	}
	return users, nil
}
