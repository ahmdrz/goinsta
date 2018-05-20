// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<another user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	media, err := user.Tags(nil)
	e.CheckErr(err)

	for media.Next() {
		fmt.Println("Next:", media.NextID)
		for _, item := range media.Items {
			fmt.Printf("  - %s has %d likes\n", item.Caption.Text, item.Likes)
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
