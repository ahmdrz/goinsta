package main

import (
	"log"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/tcnksm/go-input"
)

func main() {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		switch v := err.(type) {
		case goinsta.ChallengeError:
			err := insta.Challenge.Process(v.Challenge.APIPath)
			if err != nil {
				log.Fatalln(err)
			}

			ui := &input.UI{
				Writer: os.Stdout,
				Reader: os.Stdin,
			}

			query := "What is SMS code for instagram?"
			code, err := ui.Ask(query, &input.Options{
				Default:  "000000",
				Required: true,
				Loop:     true,
			})
			if err != nil {
				log.Fatalln(err)
			}

			err = insta.Challenge.SendSecurityCode(code)
			if err != nil {
				log.Fatalln(err)
			}

			insta.Account = insta.Challenge.LoggedInUser
		default:
			log.Fatalln(err)
		}

		log.Printf("logged in as %s \n", insta.Account.Username)
	}

	defer insta.Logout()
}
