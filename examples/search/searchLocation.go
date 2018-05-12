// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(5, "<your user> <lat> <lng> <query>")
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
