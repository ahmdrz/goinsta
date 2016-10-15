package goinsta

import (
	"net/http"
	"net/url"
	"sync"
)

type jar struct {
	lk      sync.Mutex
	cookies map[string][]*http.Cookie
}

func newJar() *jar {
	tjar := new(jar)
	tjar.cookies = make(map[string][]*http.Cookie)
	return tjar
}

// SetCookies handles the receipt of the cookies in a reply for the
// given URL.  It may or may not choose to save the cookies, depending
// on the jar's policy and implementation.
func (tjar *jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	tjar.lk.Lock()
	tjar.cookies[u.Host] = cookies
	tjar.lk.Unlock()
}

// Cookies returns the cookies to send in a request for the given URL.
// It is up to the implementation to honor the standard cookie use
// restrictions such as in RFC 6265.
func (tjar *jar) Cookies(u *url.URL) []*http.Cookie {
	return tjar.cookies[u.Host]
}
