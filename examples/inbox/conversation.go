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

	c := inst.Inbox.Conversations[0]

	fmt.Printf("Opening conversation with %s\n", c.Inviter.Username)

	for c.Next() {
		for _, i := range c.Items {
			if i.Type == "text" {
				if i.UserID == inst.Account.ID {
					fmt.Printf("Me: %s\n", i.Text)
				} else {
					fmt.Printf("%s: %s\n", c.Inviter.Username, i.Text)
				}
			}
		}
	}
	fmt.Println(c.Error())

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
