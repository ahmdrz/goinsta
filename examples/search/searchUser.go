// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<query>")
	e.CheckErr(err)

	res, err := inst.Search.User(os.Args[0])
	e.CheckErr(err)

	for _, user := range res.Users {
		fmt.Printf("    %s\n", user.Username)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
