// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	act := inst.Activity.Recent()
	e.CheckErr(err)

	for act.Next() {
		fmt.Printf("Stories: %d %d\n", len(act.Stories), act.NextID)
	}
	fmt.Println(act.Error())

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
