// +build ignore

package main

import (
	"fmt"
	"os"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media := inst.AcquireFeed()
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Println("Deleting", os.Args[2])
	err = media.Items[0].Delete()
	e.CheckErr(err)

	err = media.Sync()
	if err != nil {
		fmt.Println("Deleted!")
	} else {
		fmt.Println("error deleting...")
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
