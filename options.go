package goinsta

import (
	"github.com/ahmdrz/goinsta/device"
)

// Option describes a functional option for configuring the Client.
type Option func(*Client) error

// SetDevice sets device of client
func SetDevice(deviceName string) Option {
	return func(c *Client) error {
		var err error
		c.device, err = device.New(deviceName)
		return err
	}
}

func SetUUID(uuid string) Option {
	return func(c *Client) error {
		c.uuid = uuid
		return nil
	}
}

func SetPhoneID(phoneID string) Option {
	return func(c *Client) error {
		c.phoneID = phoneID
		return nil
	}
}

func SetAdvertisingID(advertisingID string) Option {
	return func(c *Client) error {
		c.advertisingID = advertisingID
		return nil
	}
}

func SetSessionID(clientSessionID string) Option {
	return func(c *Client) error {
		c.clientSessionID = clientSessionID
		return nil
	}
}
