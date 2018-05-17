// +build ignore

package main

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	err = inst.Inbox.Sync()
	e.CheckErr(err)

	fmt.Printf("You have %d opened conversations\n", len(inst.Inbox.Threads))

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
