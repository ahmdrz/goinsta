// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<lat> <lng> <query>")
	e.CheckErr(err)

	res, err := inst.Search.Location(
		os.Args[2], os.Args[3], os.Args[4],
	)
	e.CheckErr(err)

	for _, venue := range res.Venues {
		fmt.Printf("    %s\n", venue.Name)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
