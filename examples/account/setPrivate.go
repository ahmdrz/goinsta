// +build ignore

package main

import (
	"fmt"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	inst, err := e.InitGoinsta(2, "<username>")
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
