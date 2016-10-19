# goinsta

> Golang Instagram API , Unofficial Instagram API for Golang

[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta) [![GoDoc](https://godoc.org/github.com/ahmdrz/goinsta?status.svg)](https://godoc.org/github.com/ahmdrz/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/ahmdrz/goinsta)](https://goreportcard.com/report/github.com/ahmdrz/goinsta)

Unofficial Instagram API written in Golang

This library work like android version of instagram

***

# Installation 

`go get -u -v github.com/ahmdrz/goinsta`

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
 - [x] GetProfileData
 - [x] SetPublicAccount
 - [x] SetPrivateAccount
 - [x] RemoveProfilePicture
 - [x] Comment
 - [x] DeleteComment
 - [x] EditMedia
 - [x] DeleteMedia
 - [x] MediaInfo
 - [x] Expose
 - [x] UploadPhoto
 - [ ] UploadVideo
 - [x] RemoveSelfTag
 - [ ] GetUsernameInfo
 - [ ] GetRecentActivity
 - [ ] GetFollowingRecentActivity
 - [x] TagFeed
 - [x] SearchUsername
 - [x] GetRecentRecipients
 - [x] ChangePassword
 - [x] SelfUserFollowers
 - [x] SelfUserFollowing
 - [x] SelfTotalUserFollowers
 - [x] SelfTotalUserFollowing

This repository is a copy of [Instagram-API-Python](https://github.com/LevPasha/Instagram-API-python) , And original source is [Instagram-API](https://github.com/mgp25/Instagram-API)

# How to use ?

The example is very simple !

### GetUserFeed

```go
package main

import (
	"fmt"

	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD")

	if err := insta.Login(); err != nil {
		panic(err)
	}

	defer insta.Logout()

	resp, err := insta.UserFeed()
	if err != nil {
		panic(err)
	}

	if resp.Status != "ok" {
		panic("Error occured , " + resp.Status)
	}

	for _, item := range resp.Items {
		if len(item.Caption.Text) > 30 {
			item.Caption.Text = item.Caption.Text[:30]
		}
		fmt.Println(item.ID, item.Caption.Text)
	}
}

```

### UploadPhoto

```go
package main

import (
	"fmt"

	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD")

	if err := insta.Login(); err != nil {
		panic(err)
	}

	defer insta.Logout()

	resp, _ := insta.UploadPhoto("PATH_TO_IMAGE", "CAPTION", insta.NewUploadID(), 87,goinsta.Filter_Lark) // default quality is 87

	fmt.Println(resp.Status)
}

```

# Thanks

1. [levpasha](https://github.com/LevPasha)
