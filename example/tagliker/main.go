// tagliker
package main

import (
	"fmt"

	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("username", "password")

	if err := insta.Login(); err != nil {
		panic(err)
	}

	defer insta.Logout()

	resp, err := insta.TagFeed("pizza")
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Items {
		_, err := insta.Like(item.ID)
		if err != nil {
			fmt.Println("ID : ", item.ID, "can't like")
		} else {
			fmt.Println("ID : ", item.ID, "liked")
		}
	}
}
