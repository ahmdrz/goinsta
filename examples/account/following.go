// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	users := inst.Account.Following()
	e.CheckErr(err)

	for users.Next() {
		fmt.Println("Next:", users.NextID)
		for _, user := range users.Users {
			fmt.Printf("   - %s\n", user.Username)
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
