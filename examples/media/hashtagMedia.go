// +build ignore

package main

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	h := inst.NewHashtag("pakillo")
	for h.Next() {
		fmt.Println("Items:", h.MediaCount)
		for i := range h.Sections {
			// I don't made that
			for _, i := range h.Sections[i].LayoutContent.Medias {
				if len(i.Item.Images.Versions) != 0 {
					fmt.Println(i.Item.Images.Versions[0].URL)
				}
			}
		}
	}
	e.CheckErr(h.Error())

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
