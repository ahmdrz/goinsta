// +build ignore

package main

import (
	"fmt"

	e "gopkg.in/ahmdrz/goinsta.v2/examples"
)

func main() {
	inst, err := e.InitGoinsta("")
	e.CheckErr(err)

	fmt.Printf("Is private: %v\n", inst.Account.IsPrivate)

	err = inst.Account.SetPrivate()
	e.CheckErr(err)
	fmt.Printf("After calling func: Is private: %v\n", inst.Account.IsPrivate)

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}
