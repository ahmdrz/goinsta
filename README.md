# goinsta

> Golang Instagram API , Unofficial Instagram API for Golang

[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta) [![GoDoc](https://godoc.org/github.com/ahmdrz/goinsta?status.svg)](https://godoc.org/github.com/ahmdrz/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/ahmdrz/goinsta)](https://goreportcard.com/report/github.com/ahmdrz/goinsta) [![Coverage Status](https://coveralls.io/repos/github/ahmdrz/goinsta/badge.svg?branch=master)](https://coveralls.io/github/ahmdrz/goinsta?branch=master)

Unofficial Instagram API written in Golang

## Legal

@mgp25
This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.

**Very simple and no-dependency Instagram API**

This library work like android version of instagram

***

# Installation 

`go get -u -v github.com/ahmdrz/goinsta`

# Methods 

 - [x] Block
 - [x] ChangePassword
 - [x] Comment
 - [x] DeleteComment
 - [x] DeleteMedia
 - [x] EditMedia
 - [x] Expose
 - [x] Follow
 - [x] GetFollowingRecentActivity
 - [x] GetProfileData
 - [x] GetRecentActivity
 - [x] GetRecentRecipients
 - [x] GetUsername
 - [x] Like
 - [x] Login
 - [x] Logout
 - [x] MediaInfo
 - [x] MediaLikers
 - [x] RemoveProfilePicture
 - [x] RemoveSelfTag
 - [x] SelfTotalUserFollowers
 - [x] SelfTotalUserFollowing
 - [x] SelfUserFollowers
 - [x] SelfUserFollowing
 - [x] SetPrivateAccount
 - [x] SetPublicAccount
 - [x] TagFeed
 - [x] UnBlock
 - [x] UnFollow
 - [x] UnLike
 - [x] UploadPhoto 
 - [x] UserFeed
 - [x] UserFollowers
 - [x] UserFollowings
 - [ ] GetGeoMedia
 - [ ] GetUserTags
 - [x] SearchFacebookUsers
 - [x] SearchTags
 - [x] SearchUsername
 - [ ] SelfGeoMedia
 - [ ] SelfUserTags
 - [ ] SyncContacts
 - [ ] UploadVideo

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
