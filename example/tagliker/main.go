// tagliker
package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahmdrz/goinsta/src"
)

func main() {
	insta := goinsta.New("username", "password")

	if err := insta.Login(); err != nil {
		panic(err)
	}

	defer insta.Logout()

	bytes, err := insta.TagFeed("pizza")
	if err != nil {
		panic(err)
	}

	type Caption struct {
		Status string `json:"status"`
		Text   string `json:"text"`
	}

	type Item struct {
		Id        string  `json:"id"`
		Caption   Caption `json:"caption"`
		LikeCount int     `json:"like_count"`
	}

	var Result struct {
		Status string `json:"status"`
		Items  []Item `json:"ranked_items"`
	}

	err = json.Unmarshal(bytes, &Result)
	if err != nil {
		panic(err)
	}

	if Result.Status != "ok" {
		panic("Error occured , " + Result.Status)
	}

	for _, item := range Result.Items {
		_, err := insta.Like(item.Id)
		if err != nil {
			fmt.Println("ID : ", item.Id, "can't like")
		} else {
			fmt.Println("ID : ", item.Id, "liked")
		}
	}
}
