package examples

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/ahmdrz/goinsta/examples"
	"github.com/howeyc/gopass"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitGoinsta(min int, msg string) (*Instagram, error) {
	var (
		nargs  = len(os.Args)
		config string
		inst   *Instagram
	)
	switch {
	// this parameters changes
	case nargs < min:
		fmt.Printf("%s %s [config file]\n", os.Args[0], msg)
		return
	case nargs == min+1:
		config = os.Args[3]
	}
	var err error

	if config != "" {
		inst, err = goinsta.Import(config)
		if err != nil {
			fmt.Println(err)
			fmt.Print("Using password to login")
			config = ""
		}
	}
	if config == "" {
		fmt.Print("Password: ")
		pass, err := gopass.GetPasswd()
		if err != nil {
			panic(err)
		}

		inst = goinsta.New(os.Args[1], string(pass))

		err = inst.Login()
		checkErr(err)
	}
}
