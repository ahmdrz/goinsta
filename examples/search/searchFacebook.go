// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<query>")
	e.CheckErr(err)

	res, err := inst.Search.Facebook(os.Args[2])
	e.CheckErr(err)

	for _, user := range res.Users {
		fmt.Printf("    %s\n", user.Username)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
