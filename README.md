# goinsta
[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta)

Unofficial Instagram API written in Golang

***

This repository is a copy of [Instagram-API-Python](https://github.com/LevPasha/Instagram-API-python) , And original source is [Instagram-API](https://github.com/mgp25/Instagram-API)

# How to use ?

The example is very simple !

Note : *every methods return array of byte , they are JSONs , you have to unmarshal*

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("USERNAME","PASSWORD")

	if err := insta.Login(); err != nil {
		panic(err)
	}

	bytes, err := insta.UserFeed()
	if err != nil {
		panic(err)
	}

	type Caption struct {
		Status string `json:"status"`
		Text   string `json:"text"`
	}

	type Item struct {
		Id      string  `json:"id"`
		Caption Caption `json:"caption"`
	}

	var Result struct {
		Status string `json:"status"`
		Items  []Item `json:"items"`
	}

	err = json.Unmarshal(bytes, &Result)
	if err != nil {
		panic(err)
	}

	if Result.Status != "ok" {
		panic("Error occured , " + Result.Status)
	}

	for _, item := range Result.Items {
		if len(item.Caption.Text) > 30 {
			item.Caption.Text = item.Caption.Text[:30]
		}
		fmt.Println(item.Id, item.Caption.Text)
	}
}
```

# Thanks

@levpasha