// +build ignore

package main

import (
	"fmt"
	"os"
	"strings"

	e "github.com/ahmdrz/goinsta/examples"
)

func main() {
	var followers []string
	var followings []string
	inst, err := e.InitGoinsta("<target user>")
	e.CheckErr(err)

	user, err := inst.Profiles.ByName(os.Args[0])
	e.CheckErr(err)

	users_following := user.Following()
	e.CheckErr(err)

	users_follower := user.Followers()
	e.CheckErr(err)

	// collect usernames from Instagram following section
	following_counter := 1
	for users_following.Next() {
		for _, user := range users_following.Users {
			following_counter++
			followings = append(followings, user.Username)
		}
	}

	// collect usernames from Instagram followers section
	follower_counter := 1
	for users_follower.Next() {
		for _, user := range users_follower.Users {
			follower_counter++
			followers = append(followers, user.Username)
		}
	}

	fmt.Println("\nTOTAL FOLLOWING:", following_counter)
	fmt.Println("TOTAL FOLLOWER:", follower_counter)
	fmt.Println("TOTAL DOESN'T FOLLOW YOU BACK:", len(differ(followings, followers)))
	fmt.Println("WHO DOESN'T FOLLOW YOU BACK:\n", strings.Trim(fmt.Sprint(differ(followings, followers)), "[]"))

	if !e.UsingSession {
		err = inst.Logout()
		e.CheckErr(err)
	}
}

func differ(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}
