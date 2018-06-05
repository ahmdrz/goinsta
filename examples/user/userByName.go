// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<another user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[0])
	e.CheckErr(err)
	fmt.Printf("Target username is %s with the id: %d\n", user.Username, user.ID)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
