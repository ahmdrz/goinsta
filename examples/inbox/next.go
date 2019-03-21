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

	i := 1
	fmt.Printf("Page %d has %d conversations\n", i, len(inst.Inbox.Conversations))

	for inst.Inbox.Next() {
		i++
		fmt.Printf("Page %d has %d conversations\n", i, len(inst.Inbox.Conversations))
	}
	//inst.Inbox.Reset()

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
