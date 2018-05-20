// +build ignore

package main

import (
	"fmt"
	"os"

	"gopkg.in/ahmdrz/goinsta.v2"
	"github.com/howeyc/gopass"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s <username> <output file>\n", os.Args[0])
		return
	}

	fmt.Print("Password: ")
	pass, err := gopass.GetPasswd()
	checkErr(err)

	inst := goinsta.New(os.Args[1], string(pass))
	err = inst.Login()
	checkErr(err)
	fmt.Printf("Hello %s\n", inst.Account.Username)

	err = inst.Export(os.Args[2])
	checkErr(err)
	// IMPORTANT: DO NOT LOGOUT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
