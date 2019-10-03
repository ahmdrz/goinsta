#### Golang + Instagram Private API
<p align="center"><img width=100% src="https://raw.githubusercontent.com/ahmdrz/goinsta/v1/resources/goinsta-image.png"></p>

> Unofficial Instagram API for Golang

[![Build Status](https://travis-ci.org/ahmdrz/goinsta.svg?branch=master)](https://travis-ci.org/ahmdrz/goinsta) [![GoDoc](https://godoc.org/github.com/ahmdrz/goinsta?status.svg)](https://godoc.org/github.com/ahmdrz/goinsta) [![Go Report Card](https://goreportcard.com/badge/github.com/ahmdrz/goinsta)](https://goreportcard.com/report/github.com/ahmdrz/goinsta) [![Gitter chat](https://badges.gitter.im/goinsta/community.png)](https://gitter.im/goinsta/community)

### Features

* **HTTP2 by default. Goinsta uses HTTP2 client enhancing performance.**
* **Object independency. Can handle multiple instagram accounts.**
* **Like Instagram mobile application**. Goinsta is very similar to Instagram official application.
* **Simple**. Goinsta is made by lazy programmers!
* **Backup methods**. You can use `Export` and `Import` functions.
* **Security**. Your password is only required to login. After login your password is deleted.
* **No External Dependencies**. GoInsta will not use any Go packages outside of the standard library.

### Package installation 

`go get -u -v gopkg.in/ahmdrz/goinsta.v2`

### Example

```go
package main

import (
	"fmt"

	"gopkg.in/ahmdrz/goinsta.v2"
)

func main() {  
  insta := goinsta.New("USERNAME", "PASSWORD")

  // Export your configuration
  // after exporting you can use Import function instead of New function.
  // insta, err := goinsta.Import("~/.goinsta")
  // it's useful when you want use goinsta repeatedly.
  insta.Export("~/.goinsta")

  ...
}
```

### Projects using `goinsta`

- [go-instabot](https://github.com/tducasse/go-instabot)
- [nick_bot](https://github.com/icholy/nick_bot)
- [instagraph](https://github.com/ahmdrz/instagraph)
- [icrawler](https://github.com/themester/icrawler)
- [ermes](https://github.com/borteo/ermes)
- [instafeed](https://github.com/falzm/instafeed)
- [goinstadownload](https://github.com/alejoloaiza/goinstadownload)
- [InstagramStoriesDownloader](https://github.com/DiSiqueira/InstagramStoriesDownloader)
- [gridcube-challenge](https://github.com/rodrwan/gridcube-challenge)
- [nyaakitties](https://github.com/gracechang/nyaakitties)
- [InstaFollower](https://github.com/Unanoc/InstaFollower)
- [follow-sync](https://github.com/kirsle/follow-sync)
- [Game DB](https://github.com/gamedb/gamedb)
- ...

### Legal

This code is in no way affiliated with, authorized, maintained, sponsored or endorsed by Instagram or any of its affiliates or subsidiaries. This is an independent and unofficial API. Use at your own risk.

### Versioning

Goinsta used gopkg.in as versioning control. Stable new API is the version v2.0. You can get it using:

```bash
$ go get -u -v gopkg.in/ahmdrz/goinsta.v2
```

Or 

If you have `GO111MODULE=on`

```
$ go get -u github.com/ahmdrz/goinsta/v2
```

### Donate

**Ahmdrz**

![btc](https://raw.githubusercontent.com/reek/anti-adblock-killer/gh-pages/images/bitcoin.png) Bitcoin: `1KjcfrBPJtM4MfBSGTqpC6RcoEW1KBh15X`

**Mester**

![btc](https://raw.githubusercontent.com/reek/anti-adblock-killer/gh-pages/images/bitcoin.png) Bitcoin: `37aogDJYBFkdSJTWG7TgcpgNweGHPCy1Ks`


[![Analytics](https://ga-beacon.appspot.com/UA-107698067-1/readme-page)](https://github.com/igrigorik/ga-beacon)

