// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	err = inst.Inbox.Sync()
	e.CheckErr(err)

	fmt.Printf("You have %d opened conversations\n", len(inst.Inbox.Conversations))

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
