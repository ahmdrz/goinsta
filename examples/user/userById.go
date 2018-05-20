// +build ignore

package main

import (
	"fmt"
	"os"
	"strconv"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<another user>")
	e.CheckErr(err)

	id, err := strconv.Atoi(os.Args[2])
	e.CheckErr(err)

	user, err := inst.Profiles.ByID(int64(id))
	e.CheckErr(err)
	fmt.Printf("Target username is %s with the id: %d\n", user.Username, user.ID)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
