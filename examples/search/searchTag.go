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

	res, err := inst.Search.Tags(os.Args[2])
	e.CheckErr(err)

	for _, tag := range res.Tags {
		fmt.Printf("    %s\n", tag.ID)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
