package goinsta

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

var (
	lastResponse *http.Response
	lastJson     string
	cookie       string
	jar          *Jar
)

func cloneValue(source interface{}, destin interface{}) {
	x := reflect.ValueOf(source)
	if x.Kind() == reflect.Ptr {
		starX := x.Elem()
		y := reflect.New(starX.Type())
		starY := y.Elem()
		starY.Set(starX)
		reflect.ValueOf(destin).Elem().Set(y.Elem())
	} else {
		destin = x.Interface()
	}
}

func (insta *Instagram) sendRequest(endpoint string, post string, login bool) error {
	if !insta.IsLoggedIn && !login {
		return fmt.Errorf("not logged in")
	}

	var req *http.Request
	var err error

	if len(post) > 0 {
		req, err = http.NewRequest("POST", API_URL+endpoint, bytes.NewBuffer([]byte(post)))
		if err != nil {
			return err
		}
	} else {
		req, err = http.NewRequest("GET", API_URL+endpoint, nil)
		if err != nil {
			return err
		}
	}

	req.Header.Set("Connection", "close")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie2", "$Version=1")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", USER_AGENT)

	tempjar := NewJar()

	if !login {
		for key, value := range jar.cookies { // make a copy of session
			tempjar.cookies[key] = value
		}
	} else {
		tempjar = jar // copy pointers (move sessions)
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
		return fmt.Errorf("Invalid status code", string(body))
	}

	return nil
}
