package main

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/ahmdrz/goinsta/utilities"
)

func main() {
	var (
		insta *goinsta.Instagram
		err   error
	)
	encodedAccount := os.Getenv("INSTAGRAM_ENCODED")
	if encodedAccount != "" {
		insta, err = utilities.ImportFromBase64String(encodedAccount)
		if err != nil {
			panic(err)
		}
	} else {
		insta = goinsta.New(os.Getenv("INSTAGRAM_USERNAME"), os.Getenv("INSTAGRAM_PASSWORD"))
		if err = insta.Login(); err != nil {
			panic(err)
		}
		result, err := utilities.ExportAsBase64String(insta)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	}

	fmt.Printf("Logged in as %s\n", insta.Account.Username)
}
