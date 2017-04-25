package goinsta

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (insta *Instagram) NewRequest(endpoint string, post string) ([]byte, error) {
	return insta.sendRequest(endpoint, post, false)
}

func (insta *Instagram) sendRequest(endpoint string, post string, options ...bool) (body []byte, err error) {
	isLoggedIn := false // Optional third argument
	checkStatus := true // Optional forth argument
	if len(options) == 1 {
		isLoggedIn = options[0]
	} else if len(options) == 2 {
		isLoggedIn = options[0]
		checkStatus = options[1]
	}

	if !insta.IsLoggedIn && !isLoggedIn {
		return nil, fmt.Errorf("not logged in")
	}

	var req *http.Request

	method := "GET"
	if len(post) > 0 {
		method = "POST"
	}
	req, err = http.NewRequest(method, GOINSTA_API_URL+endpoint, bytes.NewBuffer([]byte(post)))
	if err != nil {
		return
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", GOINSTA_USER_AGENT)

	client := &http.Client{
		Jar: insta.cookiejar,
	}
	if insta.proxy != "" {
		proxy, err := url.Parse(insta.proxy)
		if err != nil {
			return body, err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxy)}
	}

	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	u, _ := url.Parse(GOINSTA_API_URL)
	for _, value := range insta.cookiejar.Cookies(u) {
		if strings.Contains(value.Name, "csrftoken") {
			insta.Informations.Token = value.Value
		}
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 && checkStatus {
		return nil, fmt.Errorf("Invalid status code %s", string(body))
	}

	return body, err
}
