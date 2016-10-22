package goinsta

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	lastResponse *http.Response
	lastJson     string
	cookie       string
	cookiejar    *jar
)

func (insta *Instagram) sendRequest(endpoint string, post string, login bool) error {
	if !insta.IsLoggedIn && !login {
		return fmt.Errorf("not logged in")
	}

	var req *http.Request
	var err error

	if len(post) > 0 {
		req, err = http.NewRequest("POST", GOINSTA_API_URL+endpoint, bytes.NewBuffer([]byte(post)))
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest("GET", GOINSTA_API_URL+endpoint, nil)
		if err != nil {
			return err
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
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	lastResponse = resp
	cookie = resp.Header.Get("Set-Cookie")

	body, _ := ioutil.ReadAll(resp.Body)

	lastJson = string(body)

	if resp.StatusCode != 200 {
		return fmt.Errorf("Invalid status code %s", string(body))
	}

	return nil
}
