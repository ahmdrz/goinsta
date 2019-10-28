package main

import (
	"log"
	"os"

	"github.com/ahmdrz/goinsta"
)

func fetchTag(insta *goinsta.Instagram, tag string) error {
	feedTag, err := insta.Feed.Tags(tag)
	if err != nil {
		return err
	}
	for _, item := range feedTag.RankedItems {
		err = item.Like()
		if err != nil {
			log.Printf("error on liking item %s, %v", item.ID, err)
		} else {
			log.Printf("item %s liked", item.ID)
		}
	}
	return nil
}

func main() {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		log.Println(err)
		return
	}
	defer insta.Logout()

	for _, tag := range []string{
		"golang",
		"pizza",
		"google",
	} {
		if err := fetchTag(insta, tag); err != nil {
			log.Println(tag, err)
		}
	}
}
