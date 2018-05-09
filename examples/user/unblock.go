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

	// if you have someone blocked probably you cannot found it with this method
	user, err := inst.Profiles.ByName(os.Args[2])
	checkErr(err)

	err = user.Unblock()
	checkErr(err)

	err = inst.Logout()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
