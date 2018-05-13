// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<username> <media id>")
	e.CheckErr(err)

	media := goinsta.AcquireFeed(inst)
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Printf("Comments disabled: %v\n", media.Items[0].CommentsDisabled)
	err = media.Items[0].Comments.Disable()
	e.CheckErr(err)

	media.Sync()
	fmt.Printf("Comments disabled: %v\n", media.Items[0].CommentsDisabled)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
