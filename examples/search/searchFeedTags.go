// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<username> <query>")
	e.CheckErr(err)

	fmt.Printf("Hello %s!\n", inst.Account.Username)

	// I don't want to make an example of this. Not today.
	tags, err := inst.Search.FeedTags(os.Args[2])
	e.CheckErr(err)

	// TODO

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
