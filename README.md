# goinsta
[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta)

Unofficial Instagram API written in Golang

***

# Installation 

`go get -u -v github.com/ahmdrz/goinsta/src`

# Methods 

 - [x] Login
 - [x] Logout
 - [x] UserFollowings
 - [x] UserFollowers
 - [x] UserFeed
 - [x] MediaLikers
 - [x] Follow
 - [x] UnFollow
 - [x] Block
 - [x] UnBlock
 - [x] Like
 - [x] UnLike
 - [ ] GetProfileData
 - [x] SetPublicAccount
 - [x] SetPrivateAccount
 - [x] RemoveProfilePicture
 - [x] Comment
 - [x] DeleteComment
 - [x] EditMedia
 - [x] DeleteMedia
 - [x] MediaInfo
 - [x] Expose
 - [ ] UploadPhoto
 - [ ] UploadVideo
 - [x] RemoveSelfTag
 - [ ] GetUsernameInfo
 - [ ] GetRecentActivity
 - [ ] GetFollowingRecentActivity
 - [x] TagFeed

This repository is a copy of [Instagram-API-Python](https://github.com/LevPasha/Instagram-API-python) , And original source is [Instagram-API](https://github.com/mgp25/Instagram-API)

# How to use ?

The example is very simple !

Note : *every methods return array of byte , they are JSONs , you have to unmarshal*


### GetUserFeed

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahmdrz/goinsta/src"
)

func main() {
	insta := goinsta.New("USERNAME","PASSWORD")

	if err := insta.Login(); err != nil {
		panic(err)
	}

    defer insta.Logout()

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

### GetTagFeed (SearchByTagName) 

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/ahmdrz/goinsta/src"
)

func main() {
	insta := goinsta.New("USERNAME","PASSWORD")	

	if err := insta.Login(); err != nil {
		panic(err)
	}
	
	defer insta.Logout()

	bytes, err := insta.TagFeed("Pizza")
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
		if len(item.Caption.Text) > 30 {
			item.Caption.Text = item.Caption.Text[:30]
		}
		fmt.Println(item.Caption.Text, "-----", item.LikeCount)
	}
}

## For more , see example folder

```

# Thanks

1. [levpasha](https://github.com/LevPasha)