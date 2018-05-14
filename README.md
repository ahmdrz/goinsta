# GoInsta. Make Goinsta Great Again!
<p align="center"><img width=100% src="https://raw.github.com/ahmdrz/goinsta/master/resources/goinsta-image.png"></p>

> Unofficial Instagram API for Golang

[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta) [![GoDoc](https://godoc.org/github.com/ahmdrz/goinsta?status.svg)](https://godoc.org/github.com/ahmdrz/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/ahmdrz/goinsta)](https://goreportcard.com/report/github.com/ahmdrz/goinsta) [![Coverage Status](https://coveralls.io/repos/github/ahmdrz/goinsta/badge.svg?branch=master)](https://coveralls.io/github/ahmdrz/goinsta?branch=master)

## Features

* **HTTP2 by default. Goinsta uses HTTP2 client enhancing performance.**
* **Object independency. Can handle multiple instagram accounts.**
* **Like Instagram mobile application**. Goinsta is very similar to Instagram official application.
* **Simple**. Goinsta is made by lazy programmers!
* **Backup methods**. You can use `Export` and `Import` functions.
* **Security**. Your password is only required to login. After login your password is deleted.
* **No External Dependencies**. Goinsta will not use any Go packages outside of the standard library.

## New Version !

We are working on `alpha` branch. Try it and tell us your suggestions!

The newer versions will be exported into v2 branch when new features will be well tested.

## Installation 

`go get -u -v gopkg.in/ahmdrz/goinsta.v2`

## Example

```go
package main

import (
	"fmt"

	"github.com/ahmdrz/goinsta"
)

func main() {
	insta := goinsta.New("USERNAME", "PASSWORD")

	if err := insta.Login(); err != nil {
		fmt.Println(err)
		return
	}
	defer insta.Logout()

	...
}
```

In the next examples you can use an optional argument to use cache config.

* [**More Examples**](https://github.com/ahmdrz/goinsta/tree/v2/examples)

## Legal

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.

## Donate

**Ahmdrz**

![btc](https://raw.githubusercontent.com/reek/anti-adblock-killer/gh-pages/images/bitcoin.png) Bitcoin: `1KjcfrBPJtM4MfBSGTqpC6RcoEW1KBh15X`

**Mester**

![btc](https://raw.githubusercontent.com/reek/anti-adblock-killer/gh-pages/images/bitcoin.png) Bitcoin: `37aogDJYBFkdSJTWG7TgcpgNweGHPCy1Ks`



[![Analytics](https://ga-beacon.appspot.com/UA-107698067-1/readme-page)](https://github.com/igrigorik/ga-beacon)


## Schema

Instagram
- Account: Personal information and account interactions.
  - Followers
  - Following
  - Feed
  - Stories
  - Tags
- Profiles: User interaction.
  - Followers
  - Following
  - FriendShip
  - Story
- Media:
  - Comments
  - Likes
  - Likers
- Search:
  - Location
  - Username
  - Tags
  - Location **Deprecated**
  - Facebook
- Activity:
  - Following
  - Recent
