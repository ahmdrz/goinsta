// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	h := inst.NewHashtag("fanboy")
	stories, err := h.Stories()
	e.CheckErr(err)

	fmt.Println(len(stories.Items))
	for _, item := range stories.Items {
		if len(item.Images.Versions) != 0 {
			fmt.Println(item.Images.Versions[0].URL)
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
