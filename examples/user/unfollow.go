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
	e.CheckErr(err)
	fmt.Printf("Hello %s!\n", inst.Account.Username)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	fmt.Printf("Unfollowing: %v\n", user.Friendship.Unfollowing)
	err = user.Unfollow()
	e.CheckErr(err)
	fmt.Printf("After func call: Unfollowing: %v\n", user.Friendship.Unfollowing)

	err = inst.Logout()
	e.CheckErr(err)
}

func e.CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
