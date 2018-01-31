# GoInsta !
<p align="center"><img width=100% src="https://raw.github.com/ahmdrz/goinsta/master/resources/goinsta-image.png"></p>

> Unofficial Instagram API for Golang

## Features

* **Like Instagram mobile application**. Goinsta is very similar to Instagram official application.
* **Simple**. Goinsta is made by a lazy programmer!
* **Backup methods**. You can use `store` package to export/import `goinsta.Instagram` struct.
* **No External Dependencies**. Goinsta will not use any Go packages outside of the standard library.

## Example

```go
package main

import (
	"log"

	"github.com/ahmdrz/goinsta-alpha"
)

func main() {
	insta := goinsta.New("####", "####")
	if err := insta.Login(); err != nil {
		log.Fatalln(err)
	}
	defer insta.Logout()

	followers, _ := insta.CurrentUser.Followers()
	for _, follower := range followers.Users {
		log.Printf("Follower information, id=%d, username=%s", follower.ID, follower.Username)
	}

	pendingRequests, _ := insta.FriendShip.Pending()
	for _, request := range pendingRequests.Users {
		log.Printf("Request from @%s with name %s", request.Username, request.FullName)
	}
}

```

## Donate

Bitcoin : `1KjcfrBPJtM4MfBSGTqpC6RcoEW1KBh15X`

[![Analytics](https://ga-beacon.appspot.com/UA-107698067-1/readme-page)](https://github.com/igrigorik/ga-beacon)
