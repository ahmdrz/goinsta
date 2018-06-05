// +build ignore

package main

import (
	"fmt"
	"os"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("<media id>")
	e.CheckErr(err)

	media, err := inst.GetMedia(os.Args[0])
	e.CheckErr(err)

	fmt.Printf("Hashtags:\n")
	for _, h := range media.Items[0].Hashtags() {
		fmt.Println(h.Name)
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
