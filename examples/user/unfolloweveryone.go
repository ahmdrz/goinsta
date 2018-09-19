// +build ignore

package main

import (
	"os"
	e "github.com/ahmdrz/goinsta/examples"
	"fmt"
)

func main() {
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[0])
	e.CheckErr(err)

	users_following := user.Following()
	e.CheckErr(err)

	//check if there are users
	if (len(users_following.Users) > 0) {
		for users_following.Next() {
			for _, user := range users_following.Users {
				err = user.Unfollow()
				e.CheckErr(err)
				fmt.Printf("username %s unfollowed\n", user.Username)
			}
		}
	} else {
		fmt.Printf("following list is empty\n")
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
