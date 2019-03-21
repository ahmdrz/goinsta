// +build ignore

package main

import (
	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	for inst.Inbox.Next() {
		for _, c := range inst.Inbox.Conversations {
			c.Like()
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
