// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[0])
	e.CheckErr(err)

	// At this context you can use:
	// user.FriendShip()
	// user.Sync(true)
	err = user.Sync(true)
	e.CheckErr(err)

	fmt.Println("Following:", user.Friendship.Following)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
