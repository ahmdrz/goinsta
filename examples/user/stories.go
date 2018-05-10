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

	stories, err := user.Stories()
	checkErr(err)

	for {
		err := stories.Next()
		if err != nil {
			break
		}

		// getting images URL
		for _, item := range stories.Items {
			if len(item.Images.Candidates) > 0 {
				fmt.Printf("  Image - %s\n", item.Images.Candidates[0].URL)
			}
			if len(item.Videos) > 0 {
				fmt.Printf("  Video - %s\n", item.Videos[0].URL)
			}
		}
	}
	fmt.Println(err)

	err = inst.Logout()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
