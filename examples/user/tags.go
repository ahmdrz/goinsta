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

	media, err := user.Tags(nil)
	checkErr(err)

	for {
		fmt.Println("Next:", media.NextID)
		for _, item := range media.Items {
			fmt.Printf("  - %s has %d likes\n", item.Caption.Text, item.Likes)
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		err = media.Next()
	}

	err = inst.Logout()
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
