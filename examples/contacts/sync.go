package main

import (
	"fmt"
	"github.com/ahmdrz/goinsta"
)

func main() {
	//insta, _ := goinsta.Import("session_dump")
	insta := goinsta.New("lnsta_login", "insta_passwordd")

	// also you can use New function from gopkg.in/ahmdrz/goinsta.v2/utils

	// insta.SetProxy("http://localhost:8080", true) // true for insecure connections
	if err := insta.Login(); err != nil {
		fmt.Println(err)
		return
	}
	// export your configuration
	// after exporting you can use Import function instead of New function.
	insta.Export("session_dump")

	empty := make([]string, 0)
	contacts := []goinsta.Contact{
		{
			Name:    "To Search 1",
			Numbers: []string{"+11111111111"},
			Emails:  empty,
		},
		{
			Name:    "To Search 2",
			Numbers: []string{"+22222222222"},
			Emails:  empty,
		},
		{
			Name:    "To Search 3",
			Numbers: empty,
			Emails:  []string{"test@mail.ex"},
		},
	}
	answer, err := insta.Contacts.SyncContacts(&contacts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(answer)
}
