// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media, err := inst.GetMedia(os.Args[0])
	e.CheckErr(err)

	fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)
	media.Items[0].Unlike()

	media.Sync()
	fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
