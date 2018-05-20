// +build ignore

package main

import (
	"fmt"
	"os"
	"time"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media := inst.AcquireFeed()
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Printf("Comments: %d\n", media.Items[0].CommentCount)
	err = media.Items[0].Comments.Add("Awesome pic!")
	e.CheckErr(err)

	fmt.Println("wait 5 seconds...")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println()

	media.Sync()
	fmt.Printf("After calling: Comments: %d\n", media.Items[0].CommentCount)

	/*
		tray, err := inst.Timeline.Stories()
		e.CheckErr(err)

		story := tray.Stories[0]
		// commenting your first timeline story xddxdxd
		fmt.Printf("Sending reply to %s\n", story.Items[0].ID)
		err = story.Items[0].Comments.Add("xd")
		e.CheckErr(err)
		TODO Causes: Media ID is missing
	*/

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
