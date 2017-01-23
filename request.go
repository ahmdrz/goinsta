package goinsta

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	lastResponse *http.Response

	cookie       string
	cookiejar    *jar
	proxyUrl     string
)

func (insta *Instagram) NewRequest(endpoint string, post string) ([]byte, error) {
	return insta.sendRequest(endpoint, post, false)
}

func (insta *Instagram) sendRequest(endpoint string, post string, login bool) (body []byte, err error) {
	if !insta.IsLoggedIn && !login {
		return nil, fmt.Errorf("not logged in")
	}

	var req *http.Request

	if len(post) > 0 {
		req, err = http.NewRequest("POST", GOINSTA_API_URL+endpoint, bytes.NewBuffer([]byte(post)))
		if err != nil {
			return
		}
	} else {
		req, err = http.NewRequest("GET", GOINSTA_API_URL+endpoint, nil)
		if err != nil {
			return
		}
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", GOINSTA_USER_AGENT)

	tempjar := newJar()

	if !login {
		for key, value := range cookiejar.cookies { // make a copy of session
			tempjar.cookies[key] = value
		}
	} else {
		tempjar = cookiejar // copy pointers (move sessions)
	}

	client := &http.Client{
		Jar: tempjar,
	}
	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return body, err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy) }
	}

	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	lastResponse = resp
	cookie = resp.Header.Get("Set-Cookie")

	body, _ = ioutil.ReadAll(resp.Body)

	insta.lastJson = body

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid status code %s", string(body))
	}

	return body, err
}
