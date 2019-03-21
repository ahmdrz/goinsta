package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta"
	"github.com/howeyc/gopass"
	homedir "github.com/mitchellh/go-homedir"
)

// New function returns Instagram object.
// In case of error this function does os.Exit(1) and prints error to Stdout.
//
// This function tries to get config from $HOME/.goinsta. If this file does not
// exists will wait from Stdin to Username and Password.
func New() *goinsta.Instagram {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("error getting home path: %s\n", err)
		os.Exit(1)
	}

	config := fmt.Sprintf("%s%c.goinsta", home, os.PathSeparator)
	if _, err := os.Stat(config); err == nil {
		inst, err := goinsta.Import(config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return inst
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	l, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	user := string(l)

	fmt.Print("Password: ")
	pass, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inst := goinsta.New(user, string(pass))
	err = inst.Login()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inst.Export(config)
	return inst
}
