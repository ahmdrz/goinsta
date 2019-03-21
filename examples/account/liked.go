// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	media := inst.Account.Liked()

	if media.Next() {
		for _, item := range media.Items {
			fmt.Printf("You liked the media %v of user: %s with total likes of %v\n", item.ID, item.User.Username, item.Likes)
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
