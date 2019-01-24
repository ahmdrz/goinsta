// +build ignore

package main

import (
	"os"
	"strings"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<target user> <text message>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[0])
	e.CheckErr(err)

	err = inst.Inbox.Sync()
	e.CheckErr(err)

	err = inst.Inbox.New(user, strings.Join(os.Args[1:], " "))
	e.CheckErr(err)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
