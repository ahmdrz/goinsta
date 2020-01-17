package device

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ahmdrz/goinsta/constants"
)

const (
	// DefaultDeviceName is a one of devices
	DefaultDeviceName = "lg_g5"
)

var (
	// ErrBadDeviceName occures when device name not matched
	// with any devices of goinsta
	ErrBadDeviceName = errors.New("device name not found")
)

var devices = map[string]Device{
	"lg_g5": Device{
		InstagramVersion: constants.InstagramVersion,
		AndroidVersion:   23,
		AndroidRelease:   "6.0.1",
		DPI:              "640dpi",
		Resolution:       "1440x2392",
		Manufacturer:     "LGE/lge",
		DeviceName:       "RS988",
		Model:            "h1",
		CPU:              "h1",
	},
}

// Device is a simulated version of real device which can interact with
// the Instagram
type Device struct {
	InstagramVersion string
	AndroidVersion   uint
	AndroidRelease   string
	DPI              string
	Resolution       string
	Manufacturer     string
	DeviceName       string
	Model            string
	CPU              string
	userAgent        string
}

func (d *Device) UserAgent() string {
	if d.userAgent == "" {
		d.userAgent = fmt.Sprintf("Instagram %s Android (%d/%s; %s; %s; %s; %s; %s; %s; en_US)",
			constants.InstagramVersion,
			d.AndroidVersion, d.AndroidRelease,
			d.DPI, d.Resolution,
			d.Manufacturer, d.DeviceName,
			d.Model, d.CPU,
		)
	}
	return d.userAgent
}

// New create a android device with default parameters
func New(deviceName string) (*Device, error) {
	if device, ok := devices[deviceName]; ok {
		d := &Device{}
		*d = device
		return d, nil
	}
	return nil, ErrBadDeviceName
}

// FromJSON parse json from bytes to device instance
func FromJSON(b []byte) (*Device, error) {
	d := &Device{}
	err := json.Unmarshal(b, d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
