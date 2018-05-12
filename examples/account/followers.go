// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(2, "<username>")
	e.CheckErr(err)

	users := inst.Account.Followers()

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
