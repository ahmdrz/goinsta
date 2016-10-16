// loverbot
package main

import (
	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("username", "password")

	if err := insta.Login(); err != nil {
		panic(err)
	}

	defer insta.Logout()

	_, err := insta.Comment("media id", "fucking cool sentences !") // one of ahmdrz images
	if err != nil {
		panic(err)
	}
}
