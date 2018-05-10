package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(2, "<your username>")
	e.CheckErr(err)

	media, err := inst.Timeline.Get()
	e.CheckErr(err)

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
