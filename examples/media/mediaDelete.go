// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<username> <media id>")
	e.CheckErr(err)

	media := goinsta.AcquireFeed(inst)
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
