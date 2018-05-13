// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<username> <target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	media := user.Feed(nil)
	e.CheckErr(err)

	for media.Next() {
		fmt.Println("Next:", media.NextID)
		for _, item := range media.Items {
			if len(item.Images.Versions) != 0 {
				fmt.Printf("  %v - %s\n", item.ID, item.Images.Versions[0].URL)
			}
		}
	}
	fmt.Println(media.Error())

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
