package goinsta

import ()

// Activity is the activity of your instagram account
type Activity struct {
	inst *Instagram
}

func newActivity(inst *Instagram) *Activity {
	act := &Activity{
		inst: inst,
	}
	return act
}
