package goinsta

import ()

type Activity struct {
	inst *Instagram
}

func newActivity(inst *Instagram) *Activity {
	act := &Activity{
		inst: inst,
	}
	return act
}
