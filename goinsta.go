package goinsta

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"regexp"

	"github.com/ahmdrz/goinsta/device"
)

const (
	usernameRegex = `[A-Za-z0-9_](?:(?:[A-Za-z0-9_]|(?:\.)){0,28}(?:[A-Za-z0-9_]))?`
)

var (
	// ErrInvalidUsername occures when username does not match with
	// regex rules of Instagram.
	ErrInvalidUsername = errors.New("bad username")
	// ErrBadPassword occures when password is less then three letters.
	ErrBadPassword = errors.New("bad password")
)

// Client is the main object of goinsta
type Client struct {
	username        string
	password        string
	device          *device.Device
	phoneID         string
	uuid            string
	deviceID        string
	clientSessionID string
	advertisingID   string
	httpClient      *http.Client
	token           string
}

// New create an instance of Client
func New(username, password string, options ...Option) (*Client, error) {
	re := regexp.MustCompile(usernameRegex)
	matches := re.FindAllString(username, -1)
	if len(matches) != 1 {
		return nil, ErrInvalidUsername
	}
	if matches[0] != username {
		return nil, ErrInvalidUsername
	}
	if len(password) < 3 {
		return nil, ErrBadPassword
	}
	jar, _ := cookiejar.New(nil)
	c := &Client{
		username: username,
		password: password,
		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
			Jar: jar,
		},
	}
	c.device, _ = device.New(device.DefaultDeviceName)
	c.phoneID = generateUUID()
	c.uuid = generateUUID()
	c.clientSessionID = generateUUID()
	c.advertisingID = generateUUID()
	c.deviceID = generateDeviceID(generateSeed(username, password))

	var err error
	for _, option := range options {
		err = option(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) readMsisdnHeader(usage string) error {
	data, _ := json.Marshal(
		map[string]string{
			"device_id":          c.uuid,
			"mobile_subno_usage": usage,
		},
	)
	_, err := c.sendRequest(
		&reqOptions{
			Endpoint:     urlMsisdnHeader,
			IsPost:       true,
			LoginProcess: true,
			Connection:   "keep-alive",
			Headers: map[string]string{
				"X-DEVICE-ID": c.uuid,
			},
			Query: generateSignature(b2s(data)),
		},
	)
	return err
}

func (c *Client) Login() error {
	err := c.readMsisdnHeader("default")
	if err != nil {
		return err
	}
	return nil
}
