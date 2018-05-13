// +build ignore

package main

import (
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<another user>")
	e.CheckErr(err)

	// if you have someone blocked probably you cannot found it with this method
	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	err = user.Unblock()
	e.CheckErr(err)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
