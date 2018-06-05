// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	media := user.Feed()

	for media.Next() {
		fmt.Printf("Printing %d items\n", len(media.Items))
		for _, item := range media.Items {
			for _, liker := range item.Likers {
				fmt.Printf("%s with username: %s Likes this \n", liker.FullName, liker.Username)
			}
		}
	}
	fmt.Println(media.Error())

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
