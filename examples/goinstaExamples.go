package examples

import (
	"fmt"
	"os"
	"strings"

	"github.com/ahmdrz/goinsta"
	"github.com/howeyc/gopass"
)

// UsingSession is used inside InitGoinsta to control if there is already a session created.
var UsingSession bool

// CheckErr is a generic function to validate all errors.
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
}

// InitGoinsta initilises examples
func InitGoinsta(msg string) (*goinsta.Instagram, error) {
	var (
		// min must be args + config file or username
		min   = len(strings.Split(msg, ">")) + 1
		nargs = len(os.Args)
		user  string
		inst  *goinsta.Instagram
	)
	switch {
	// this parameters changes
	case nargs < min:
		fmt.Printf("%s <username or config file> %s\n", os.Args[0], msg)
		os.Exit(0)
	default:
		user = os.Args[1]
	}

	if _, err := os.Stat(user); err == nil {
		inst, err = goinsta.Import(user)
		if err != nil {
			return nil, err
		}

		UsingSession = true
	} else {
		fmt.Print("Password: ")
		pass, err := gopass.GetPasswd()
		if err != nil {
			return nil, err
		}

		inst = goinsta.New(user, string(pass))

		err = inst.Login()
		if err != nil {
			return inst, err
		}
	}
	if min > 0 {
		os.Args = os.Args[2:]
	}

	fmt.Printf("Hello %s!\n", inst.Account.Username)
	return inst, nil
}
