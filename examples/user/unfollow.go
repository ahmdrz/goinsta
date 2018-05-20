// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<another user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	fmt.Printf("Unfollowing: %v\n", user.Friendship.Following)
	err = user.Unfollow()
	e.CheckErr(err)
	fmt.Printf("After func call: Unfollowing: %v\n", user.Friendship.Following)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
