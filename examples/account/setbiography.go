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

	newBiography := strings.Join(os.Args[2:], " ")

	fmt.Printf("Setting biography to: %s", newBiography)

	err = inst.Account.SetBiography(newBiography)
	e.CheckErr(err)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
