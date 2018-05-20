// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[2])
	e.CheckErr(err)

	// You can downlaod Stories or Feed images.
	// media := user.Feed(nil)
	media := user.Stories()
	e.CheckErr(err)

	for media.Next() {
		for _, item := range media.Items {
			err = item.Download("./files/", "")
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
	fmt.Println(media.Error())

	err = inst.Logout()
	e.CheckErr(err)
}
