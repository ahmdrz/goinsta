// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	users := user.Followers()
	e.CheckErr(err)

	i := 1
	for users.Next() {
		fmt.Println("Next:", users.NextID)
		for _, user := range users.Users {
			i++
			fmt.Printf("  - %s\n", user.Username)
		}
	}
	fmt.Println("Followers:", i)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
