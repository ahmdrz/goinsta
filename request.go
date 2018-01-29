package goinsta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type reqOptions struct {
	Endpoint     string
	PostData     string
	IsLoggedIn   bool
	IgnoreStatus bool
	Query        map[string]string
}

func (insta *Instagram) OptionalRequest(endpoint string, a ...interface{}) (body []byte, err error) {
	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf(endpoint, a...),
	})
}

func (insta *Instagram) sendSimpleRequest(endpoint string, a ...interface{}) (body []byte, err error) {
	return insta.sendRequest(&reqOptions{
		Endpoint: fmt.Sprintf(endpoint, a...),
	})
}

func (insta *Instagram) sendRequest(o *reqOptions) (body []byte, err error) {
	if !insta.isLoggedIn && !o.IsLoggedIn {
		return nil, ErrLoggedOut
	}

	method := "GET"
	if len(o.PostData) > 0 {
		method = "POST"
	}

	u, err := url.Parse(GOINSTA_API_URL + o.Endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for k, v := range o.Query {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	var req *http.Request
	req, err = http.NewRequest(method, u.String(), bytes.NewBuffer([]byte(o.PostData)))
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
		insta.transport.Proxy = http.ProxyURL(proxy)

		client.Transport = &insta.transport
	} else {
		// Remove proxy if insta.Proxy was removed
		insta.transport.Proxy = nil
		client.Transport = &insta.transport
	}

	resp, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	u, _ = url.Parse(GOINSTA_API_URL)
	for _, value := range insta.cookiejar.Cookies(u) {
		if strings.Contains(value.Name, "csrftoken") {
			insta.token = value.Value
		}
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 && !o.IgnoreStatus {
		e := fmt.Errorf("Invalid status code %s", string(body))
		switch resp.StatusCode {
		case 400:
			e = ErrLoggedOut
		case 404:
			e = ErrNotFound
		}
		return nil, e
	}

	return body, err
}

func (insta *Instagram) prepareData(otherData ...map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"_uuid":      insta.uuid,
		"_uid":       insta.CurrentUser.ID,
		"_csrftoken": insta.token,
	}
	if len(otherData) > 0 {
		for i := range otherData {
			for key, value := range otherData[i] {
				data[key] = value
			}
		}
	}
	bytes, err := json.Marshal(data)
	return string(bytes), err
}
