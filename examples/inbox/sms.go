// +build ignore

package main

import (
	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	err = inst.Inbox.Sync()
	e.CheckErr(err)

	if len(inst.Inbox.Conversations) != 0 {
		err = inst.Inbox.Conversations[0].Send("dfghj")
		e.CheckErr(err)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
