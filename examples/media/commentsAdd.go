// +build ignore

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ahmdrz/goinsta"
	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(3, "<username> <media id>")
	e.CheckErr(err)

	media := goinsta.AcquireFeed(inst)
	media.SetID(os.Args[2])
	media.Sync()

	fmt.Printf("Comments: %d\n", media.Items[0].CommentCount)
	err = media.Items[0].Comments.Add("adf")
	e.CheckErr(err)

	fmt.Println("wait 5 seconds...")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println()

	media.Sync()
	fmt.Printf("After calling: Comments: %d\n", media.Items[0].CommentCount)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
