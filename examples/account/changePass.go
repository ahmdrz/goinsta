// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/howeyc/gopass"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s <your user>\n", os.Args[0])
		return
	}

	fmt.Print("Password: ")
	pass, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}

	fmt.Print("New password: ")
	pass2, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}

	inst := goinsta.New(os.Args[1], string(pass))

	err = inst.Login()
	checkErr(err)
	fmt.Printf("Hello %s!\n", inst.Account.Username)

	err = inst.Account.ChangePassword(string(pass), string(pass2))
	if err == nil {
		fmt.Printf("Password have been changed\n")
	} else {
		fmt.Printf("Password cannot be changed: %s\n", err)
	}

	err = inst.Logout()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
