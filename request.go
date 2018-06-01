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
	// Connection is connection header. Default is "close".
	Connection string

	// Login url
	Login bool

	// Endpoint is the request path of instagram api
	Endpoint string

	// IsPost set to true will send request with POST method.
	//
	// By default this option is false.
	IsPost bool

	// UseV2 is set when API endpoint uses v2 url.
	UseV2 bool

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
	if o.Connection == "" {
		o.Connection = "close"
	}

	nu := goInstaAPIUrl
	switch {
	case o.UseV2:
		nu = goInstaAPIUrlv2
	case o.Login:
		nu = goInstaAPIUrlLogin
	}

	u, err := url.Parse(nu + o.Endpoint)
	if err != nil {
		return nil, err
	}

	vs := url.Values{}
	bf := bytes.NewBuffer([]byte{})

	for k, v := range o.Query {
		vs.Add(k, v)
	}

	if o.IsPost {
		bf.WriteString(vs.Encode())
	} else {
		for k, v := range u.Query() {
			vs.Add(k, strings.Join(v, " "))
		}

		u.RawQuery = vs.Encode()
	}

	var req *http.Request
	req, err = http.NewRequest(method, u.String(), bf)
	if err != nil {
		return
	}

	req.Header.Set("Connection", o.Connection)
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", goInstaUserAgent)

	resp, err := inst.c.Do(req)
	if err != nil {
		return nil, err
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
		ierr := instaError400{}
		err = json.Unmarshal(body, &ierr)
		if err == nil && ierr.Payload.Message != "" {
			return nil, instaToErr(ierr)
		}
		fallthrough
	default:
		ierr := instaError{}
		err = json.Unmarshal(body, &ierr)
		if err != nil {
			return nil, fmt.Errorf("Invalid status code: %d", resp.StatusCode)
		}
		return nil, instaToErr(ierr)
	}

	return body, err
}

func (insta *Instagram) prepareData(other ...map[string]interface{}) (string, error) {
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
	b, err := json.Marshal(data)
	if err == nil {
		return b2s(b), err
	}
	return "", err
}

func (insta *Instagram) prepareDataQuery(other ...map[string]interface{}) map[string]string {
	data := map[string]string{
		"_uuid":      insta.uuid,
		"_csrftoken": insta.token,
	}
	for i := range other {
		for key, value := range other[i] {
			data[key] = toString(value)
		}
	}
	return data
}
