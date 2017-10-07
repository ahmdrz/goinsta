# Goinsta !
<p align="center"><img width=100% src="https://raw.github.com/ahmdrz/goinsta/master/resources/goinsta-image.png"></p>

> Golang Instagram API , Unofficial Instagram API for Golang

[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta) [![GoDoc](https://godoc.org/github.com/ahmdrz/goinsta?status.svg)](https://godoc.org/github.com/ahmdrz/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/ahmdrz/goinsta)](https://goreportcard.com/report/github.com/ahmdrz/goinsta) [![Coverage Status](https://coveralls.io/repos/github/ahmdrz/goinsta/badge.svg?branch=master)](https://coveralls.io/github/ahmdrz/goinsta?branch=master)

## Features

* **Like Instagram mobile application**. Goinsta is very similar to Instagram official application.
* **Simple**. Goinsta made by lazy programmer !
* **Backup methods**. You can use `store` package to export/import `goinsta.Instagram` struct.
* **No Dependency**. Goinsta will not use any unofficial Golang packages.

## Installation 

`go get -u -v github.com/ahmdrz/goinsta`

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
		panic(err)
	}

	defer insta.Logout()

	...
}
```

## Legal

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.


## Other programs , built with GoInsta

[follow-sync](https://github.com/kirsle/follow-sync) by [@kirsle](https://github.com/kirsle)

[nick_bot](https://github.com/icholy/nick_bot) by [@icholy](https://github.com/icholy)

[unsplash-to-instagram](https://github.com/nguyenvanduocit/unsplash-to-instagram) by [@nguyenvanduocit](https://github.com/nguyenvanduocit)

[go-instabot](https://github.com/tducasse/go-instabot) by [@tducasse](https://github.com/tducasse)

[gopostal](https://github.com/scisci/gopostal) by [@scisci](https://github.com/scisci)

[go_toy](https://github.com/Rhadow/go_toy) by [@Rhadow](https://github.com/Rhadow)

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
15. [@hadidimad](https://github.com/hadidimad)

## Donate

Bitcoin : `1KjcfrBPJtM4MfBSGTqpC6RcoEW1KBh15X`
