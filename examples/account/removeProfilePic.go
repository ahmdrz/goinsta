// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(2, "<username>")
	e.CheckErr(err)

	fmt.Printf("Profile picture URL: %s\n", inst.Account.ProfilePicURL)

	err = inst.Account.RemoveProfilePic()
	e.CheckErr(err)
	fmt.Printf("After calling func: Profile picture URL: %s\n", inst.Account.ProfilePicURL)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
