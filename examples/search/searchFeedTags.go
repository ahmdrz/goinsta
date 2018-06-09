// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<query>")
	e.CheckErr(err)

	fmt.Printf("Hello %s!\n", inst.Account.Username)

	tags, err := inst.Search.FeedTags(os.Args[0])
	e.CheckErr(err)

	for _, item := range tags.Images {
		fmt.Printf("   Media found with ID: %s from User %s\n", item.ID, item.User.Username)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
