// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<media id>")
	e.CheckErr(err)

	media := inst.AcquireFeed()
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Printf("Comments disabled: %v\n", media.Items[0].CommentsDisabled)
	err = media.Items[0].Comments.Enable()
	e.CheckErr(err)

	media.Sync()
	fmt.Printf("Comments disabled: %v\n", media.Items[0].CommentsDisabled)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
