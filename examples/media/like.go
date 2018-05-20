// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media := inst.AcquireFeed()
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)
	media.Items[0].Like()

	media.Sync()
	fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
