package main

import (
	"log"

	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD")
	if err := insta.Login(); err != nil {
		panic(err)
	}
	defer insta.Logout()

	exploreFeeds, err := insta.Explore()
	if err != nil {
		log.Println(err)
		return
	}

	for _, item := range exploreFeeds.Items {
		_, err = insta.Follow(item.User.Pk)
		log.Printf("User %s followed with %v status.\n", item.User.Username, err == nil)
	}
}
