package goinsta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ahmdrz/goinsta/constants"
)

type reqOptions struct {
	// Connection is connection header. Default is "keep-alive".
	Connection string

	// LoginProcess process of login
	LoginProcess bool

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

	// Headers is the headers of the request
	Headers map[string]string
}

func (c *Client) sendRequest(o *reqOptions) (body []byte, err error) {
	method := "GET"
	if o.IsPost {
		method = "POST"
	}
	if o.Connection == "" {
		o.Connection = "keep-alive"
	}

	apiURL := constants.IGAPIUrl
	if o.UseV2 {
		apiURL = constants.IGAPIUrlv2
	}

	u, err := url.Parse(apiURL + o.Endpoint)
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
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", c.device.UserAgent())
	req.Header.Set("X-IG-App-ID", constants.IGAppID)
	req.Header.Set("X-IG-Prefetch-Request", "foreground")
	req.Header.Set("X-IG-VP9-Capable", "false")
	req.Header.Set("X-FB-HTTP-Engine", "Liger")
	req.Header.Set("X-IG-Capabilities", constants.IGCapabilities)
	req.Header.Set("X-IG-Connection-Type", "WIFI")
	req.Header.Set("X-IG-Connection-Speed", "-1.0")
	req.Header.Set("X-IG-Bandwidth-Speed-KBPS", fmt.Sprintf("%dkbps", acquireRand(8000, 10000)))
	req.Header.Set("X-IG-Bandwidth-TotalBytes-B", fmt.Sprintf("%d", acquireRand(500000, 1000000)))
	req.Header.Set("X-IG-Bandwidth-TotalTime-MS", fmt.Sprintf("%d", acquireRand(50, 150)))
	for key, value := range o.Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	u, _ = url.Parse(apiURL)
	for _, value := range c.httpClient.Jar.Cookies(u) {
		if strings.Contains(value.Name, "csrftoken") {
			c.token = value.Value
		}
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err == nil {
		err = isError(resp.StatusCode, body)
	}
	return body, err
}

// ErrorN is general instagram error
type ErrorN struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

// Error503 is instagram API error
type Error503 struct {
	Message string
}

func (e Error503) Error() string {
	return e.Message
}

func (e ErrorN) Error() string {
	return fmt.Sprintf("%s: %s (%s)", e.Status, e.Message, e.ErrorType)
}

// Error400 is error returned by HTTP 400 status code.
type Error400 struct {
	ChallengeError
	Action     string `json:"action"`
	StatusCode string `json:"status_code"`
	Payload    struct {
		ClientContext string `json:"client_context"`
		Message       string `json:"message"`
	} `json:"payload"`
	Status string `json:"status"`
}

func (e Error400) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Payload.Message)
}

// ChallengeError is error returned by HTTP 400 status code.
type ChallengeError struct {
	Message   string `json:"message"`
	Challenge struct {
		URL               string `json:"url"`
		APIPath           string `json:"api_path"`
		HideWebviewHeader bool   `json:"hide_webview_header"`
		Lock              bool   `json:"lock"`
		Logout            bool   `json:"logout"`
		NativeFlow        bool   `json:"native_flow"`
	} `json:"challenge"`
	Status    string `json:"status"`
	ErrorType string `json:"error_type"`
}

func (e ChallengeError) Error() string {
	return fmt.Sprintf("%s: %s", e.Status, e.Message)
}

func isError(code int, body []byte) (err error) {
	switch code {
	case 200:
	case 503:
		return Error503{
			Message: "Instagram API error. Try it later.",
		}
	case 400:
		ierr := Error400{}
		err = json.Unmarshal(body, &ierr)
		if err != nil {
			return err
		}

		if ierr.Message == "challenge_required" {
			return ierr.ChallengeError
		}

		if err == nil && ierr.Message != "" {
			return ierr
		}
	default:
		ierr := ErrorN{}
		err = json.Unmarshal(body, &ierr)
		if err != nil {
			return err
		}
		return ierr
	}
	return nil
}

func acquireRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
