// +build ignore

package main

import (
	"fmt"
	"os"
	"time"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media, err := inst.GetMedia(os.Args[2])
	e.CheckErr(err)

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

	tray, err := inst.Timeline.Stories()
	e.CheckErr(err)

	story := tray.Stories[1]
	// commenting your first timeline story xddxdxd
	fmt.Printf("Sending reply to %s %s\n", story.Items[0].Images.GetBest(), story.Items[0].MediaToString())
	err = story.Items[0].Comments.Add("xasfdsaf")
	e.CheckErr(err)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
