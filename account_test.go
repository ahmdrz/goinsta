package goinsta

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
	"github.com/howeyc/gopass"
	"gopkg.in/ahmdrz/goinsta.v2"
)

func ExampleAccount_ChangePassword() {
	// See more: example/account/changepass.go
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
	e.CheckErr(err)
	fmt.Printf("Hello %s!\n", inst.Account.Username)

	err = inst.Account.ChangePassword(string(pass), string(pass2))
	if err == nil {
		fmt.Printf("Password have been changed\n")
	} else {
		fmt.Printf("Password cannot be changed: %s\n", err)
	}

	err = inst.Logout()
	e.CheckErr(err)
}

func ExampleAccount_Followers() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	users := inst.Account.Followers()

	for users.Next() {
		fmt.Println("Next:", users.NextID)
		for _, user := range users.Users {
			fmt.Printf("   - %s\n", user.Username)
		}
	}
}

func ExampleAccount_Following() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	users := inst.Account.Following()
	e.CheckErr(err)

	for users.Next() {
		fmt.Println("Next:", users.NextID)
		for _, user := range users.Users {
			fmt.Printf("   - %s\n", user.Username)
		}
	}
}

func ExampleAccount_RemoveProfilePic() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	fmt.Printf("Profile picture URL: %s\n", inst.Account.ProfilePicURL)

	err = inst.Account.RemoveProfilePic()
	e.CheckErr(err)
	fmt.Printf("After calling func: Profile picture URL: %s\n", inst.Account.ProfilePicURL)
}

func ExampleAccount_SetPrivate() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	fmt.Printf("Is private: %v\n", inst.Account.IsPrivate)

	err = inst.Account.SetPrivate()
	e.CheckErr(err)
	fmt.Printf("After calling func: Is private: %v\n", inst.Account.IsPrivate)
}

func ExampleAccount_SetPublic() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	fmt.Printf("Is private: %v\n", inst.Account.IsPrivate)

	err = inst.Account.SetPublic()
	e.CheckErr(err)
	fmt.Printf("After calling func: Is private: %v\n", inst.Account.IsPrivate)
}

func ExampleAccount_Stories() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	stories := inst.Account.Stories()
	e.CheckErr(err)

	for stories.Next() {
		// getting images URL
		for _, item := range stories.Items {
			if len(item.Images.Versions) > 0 {
				fmt.Printf("  Image - %s\n", item.Images.Versions[0].URL)
			}
			if len(item.Videos) > 0 {
				fmt.Printf("  Video - %s\n", item.Videos[0].URL)
			}
		}
	}
	fmt.Println(stories.Error())
}
