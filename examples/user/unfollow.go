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

	fmt.Printf("Unfollowing: %v\n", user.Friendship.Unfollowing)
	err = user.Unfollow()
	checkErr(err)
	fmt.Printf("After func call: Unfollowing: %v\n", user.Friendship.Unfollowing)

	err = inst.Logout()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
