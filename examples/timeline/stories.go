// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	tray, err := inst.Timeline.Stories()
	e.CheckErr(err)

	if s := len(tray.Lives.LiveItems); s != 0 {
		fmt.Printf("%d people are currently on live.\n", s)
	}

	for _, story := range tray.Stories {
		// getting images URL
		for _, item := range story.Items {
			if len(item.Images.Versions) > 0 {
				fmt.Printf("  Image - %s\n", item.Images.Versions[0].URL)
			}
			if len(item.Videos) > 0 {
				fmt.Printf("  Video - %s\n", item.Videos[0].URL)
			}
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
