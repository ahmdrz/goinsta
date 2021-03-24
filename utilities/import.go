package utilities

import (
	"bytes"
	"encoding/base64"

	"github.com/ahmdrz/goinsta/v2"
)

// ImportFromBytes imports instagram configuration from an array of bytes.
//
// This function does not set proxy automatically. Use SetProxy after this call.
func ImportFromBytes(inputBytes []byte) (*goinsta.Instagram, error) {
	return goinsta.ImportReader(bytes.NewReader(inputBytes))
}

// ImportFromBase64String imports instagram configuration from a base64 encoded string.
//
// This function does not set proxy automatically. Use SetProxy after this call.
func ImportFromBase64String(base64String string) (*goinsta.Instagram, error) {
	sDec, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return ImportFromBytes(sDec)
}
