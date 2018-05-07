package main

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/howeyc/gopass"
)

func main() {
	fmt.Print("Password: ")
	pass, err := gopass.GetPasswd()
	if err != nil {
		panic(err)
	}

	insta := goinsta.New(os.Args[1], string(pass))
	insta.Login()
	insta.GetV2Inbox("")
	insta.Logout()
}
