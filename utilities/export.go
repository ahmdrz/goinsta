package utilities

import (
	"bytes"
	"encoding/base64"

	"github.com/ahmdrz/goinsta/v2"
)

// ExportAsBytes exports selected *Instagram object as []byte
func ExportAsBytes(insta *goinsta.Instagram) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := goinsta.Export(insta, buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// ExportAsBase64String exports selected *Instagram object as base64 encoded string
func ExportAsBase64String(insta *goinsta.Instagram) (string, error) {
	bytes, err := ExportAsBytes(insta)
	if err != nil {
		return "", err
	}

	sEnc := base64.StdEncoding.EncodeToString(bytes)
	return sEnc, nil
}
