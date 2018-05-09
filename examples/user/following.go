// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/howeyc/gopass"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s <your user> <another user>\n", os.Args[0])
		return
	}

	fmt.Print("Password: ")
	pass, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}

	inst := goinsta.New(os.Args[1], string(pass))

	err = inst.Login()
	checkErr(err)
	fmt.Printf("Hello %s!\n", inst.Account.Username)

	user, err := inst.Profiles.ByName(os.Args[2])
	checkErr(err)

	users, err := user.Following()
	checkErr(err)

	i := 1
	for {
		fmt.Println("Next:", users.NextID)
		for _, user := range users.Users {
			i++
			fmt.Printf("  - %s\n", user.Username)
		}

		if err = users.Next(); err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("Following:", i)

	err = inst.Logout()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
