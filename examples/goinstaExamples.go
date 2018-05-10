package examples

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/howeyc/gopass"
)

var UsingSession bool

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("error: %s\n")
		os.Exit(1)
	}
}

func InitGoinsta(min int, msg string) (*goinsta.Instagram, error) {
	var (
		nargs  = len(os.Args)
		config string
		inst   *goinsta.Instagram
	)
	switch {
	// this parameters changes
	case nargs < min:
		fmt.Printf("%s %s [config file]\n", os.Args[0], msg)
		os.Exit(0)
	case nargs == min+1:
		config = os.Args[3]
	}
	var err error

	if config != "" {
		inst, err = goinsta.Import(config)
		if err == nil {
			UsingSession = true
		} else {
			fmt.Println(err)
			fmt.Print("Using password to login")
			config = ""
		}
	}
	if config == "" {
		fmt.Print("Password: ")
		pass, err := gopass.GetPasswd()
		if err != nil {
			return nil, err
		}

		inst = goinsta.New(os.Args[1], string(pass))

		err = inst.Login()
		if err != nil {
			return inst, err
		}
	}

	fmt.Printf("Hello %s!\n", inst.Account.Username)
	return inst, nil
}
