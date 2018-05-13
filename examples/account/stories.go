// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	stories := inst.Account.Stories()
	e.CheckErr(err)

	for stories.Next() {
		// getting images URL
		for _, item := range stories.Items {
			if len(item.Images.Versions) > 0 {
				fmt.Printf("  Image - %s\n", item.Images.Versions[0].URL)
			}
			if len(item.Videos) > 0 {
				fmt.Printf("  Video - %s\n", item.Videos[0].URL)
			}
		}
	}
	fmt.Println(stories.Error())

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
