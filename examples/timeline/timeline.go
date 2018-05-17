// +build ignore

package main

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	media := inst.Timeline.Get()

	for i := 0; i < 2; i++ {
		media.Next()

		fmt.Println("Next:", media.NextID)
		for _, item := range media.Items {
			fmt.Printf("  - %s has %d likes\n", item.Caption.Text, item.Likes)
		}
	}

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
