# goinsta

### This is `new` branch

> Golang Instagram API , Unofficial Instagram API for Golang

[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta) [![GoDoc](https://godoc.org/github.com/ahmdrz/goinsta?status.svg)](https://godoc.org/github.com/ahmdrz/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/ahmdrz/goinsta)](https://goreportcard.com/report/github.com/ahmdrz/goinsta) [![Coverage Status](https://coveralls.io/repos/github/ahmdrz/goinsta/badge.svg?branch=master)](https://coveralls.io/github/ahmdrz/goinsta?branch=master)

Unofficial Instagram API written in Golang

## Legal

[@mgp25](https://github.com/mgp25)
This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.

**Very simple and no-dependency Instagram API**

This library work like android version of instagram

***

## Installation 

`go get -u -v github.com/ahmdrz/goinsta`

## Methods 

There is lot of methods , like uploadPhoto , follow , unfollow , comment , like and etc...

This repository is a copy of [Instagram-API-Python](https://github.com/LevPasha/Instagram-API-python) , And original source is [Instagram-API](https://github.com/mgp25/Instagram-API)

## How to use ?

The example is very simple !

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

	...
}
```

## Export / Import with AES encryption

```go
package main

import "github.com/ahmdrz/goinsta"

func main() {
	...

	bytes, err := insta.Export([]byte("32Byte AES Key"))
	if err != nil {
		panic(err)
	}
	// save bytes to backup file

	...

	// read bytes from backup file
	insta, err := goinsta.Import([]byte(bytes_from_file), []byte("32Byte AES Key"))
	if err != nil {
		panic(err)
	}
}
```

## Does `goinsta` support proxy servers ?
Yes, you may create goinsta object using: 

```go
insta := goinsta.NewViaProxy("USERNAME", "PASSWORD", "http://<ip>:<port>")
```


## Other programs , built with GoInsta

[follow-sync](https://github.com/kirsle/follow-sync) by [@kirsle](https://github.com/kirsle)

[nick_bot](https://github.com/icholy/nick_bot) by [@icholy](https://github.com/icholy)

[unsplash-to-instagram](https://github.com/nguyenvanduocit/unsplash-to-instagram) by [@nguyenvanduocit](https://github.com/nguyenvanduocit)

## Contributors :heart:

1. [@sourcesoft](https://github.com/sourcesoft)
2. [@GhostRussia](https://github.com/GhostRussia)
3. [@icholy](https://github.com/icholy)
4. [@rakd](https://github.com/rakd)
5. [@kemics](https://github.com/kemics)
6. [@sklinkert](https://github.com/sklinkert)
7. [@vitaliikapliuk](https://github.com/vitaliikapliuk)
8. [@glebtv](https://github.com/glebtv)
9. [@neetkee](https://github.com/neetkee)
10. [@daciwei](https://github.com/daciwei)
11. [@aaronarduino](https://github.com/aaronarduino)
12. [@tggo](https://github.com/tggo)
13. [@Albina-art](https://github.com/Albina-art)
14. [@maniack](https://github.com/maniack)
