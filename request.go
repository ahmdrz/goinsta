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
	// Endpoint is the request path of instagram api
	Endpoint string

	// IsPost setted to true will send request with POST method.
	//
	// By default this option is false.
	IsPost bool

	// Query is the parameters of the request
	//
	// This parameters are independents of the request method (POST|GET)
	Query map[string]string
}

func (insta *Instagram) sendSimpleRequest(uri string, a ...interface{}) (body []byte, err error) {
	return insta.sendRequest(
		&reqOptions{
			Endpoint: fmt.Sprintf(uri, a...),
		},
	)
}

func (inst *Instagram) sendRequest(o *reqOptions) (body []byte, err error) {
	method := "GET"
	if o.IsPost {
		method = "POST"
	}

	u, err := url.Parse(goInstaAPIUrl + o.Endpoint)
	if err != nil {
		return nil, err
	}

	bf := bytes.NewBuffer([]byte{})

	if o.IsPost {
		bf.WriteString(q.Encode())
	} else {
		u.RawQuery = q.Encode()
	}

	var req *http.Request
	req, err = http.NewRequest(method, u.String(), bf)
	if err != nil {
		return
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", goInstaUserAgent)

	resp, err := inst.c.Do(req)
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	u, _ = url.Parse(goInstaAPIUrl)
	for _, value := range inst.c.Jar.Cookies(u) {
		if strings.Contains(value.Name, "csrftoken") {
			inst.token = value.Value
		}
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	switch resp.StatusCode {
	case 200:
	case 400:
		err = ErrLoggedOut
	case 404:
		err = ErrNotFound
	default:
		err = fmt.Errorf("Invalid status code %s", string(body))
	}

	return body, err
}

func (insta *Instagram) prepareData(other ...map[string]interface{}) ([]byte, error) {
	data := map[string]interface{}{
		"_uuid":      insta.uuid,
		"_uid":       insta.Account.ID,
		"_csrftoken": insta.token,
	}
	for i := range other {
		for key, value := range other[i] {
			data[key] = value
		}
	}
	return json.Marshal(data)
}
