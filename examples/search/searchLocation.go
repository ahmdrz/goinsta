// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<lat> <lng> <query>")
	e.CheckErr(err)

	res, err := inst.Search.Location(
		os.Args[0], os.Args[1], os.Args[2],
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
