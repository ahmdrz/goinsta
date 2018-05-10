// +build ignore

package main

import (
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<your username> <target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	err = user.Block()
	e.CheckErr(err)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
