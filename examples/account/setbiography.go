// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<New biography text>")
	e.CheckErr(err)

	newBiography := strings.Join(os.Args, " ")

	fmt.Printf("Your current biography: %s\n", inst.Account.Biography)
	fmt.Printf("Setting biography to: %s\n", newBiography)

	err = inst.Account.SetBiography(newBiography)
	e.CheckErr(err)

	fmt.Printf("Your current biography (after update): %s\n", inst.Account.Biography)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
