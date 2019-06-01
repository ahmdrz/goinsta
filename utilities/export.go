package utilities

import (
	"bytes"

	"github.com/ahmdrz/goinsta"
)

func ExportAsBytes(insta *goinsta.Instagram) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := goinsta.Export(insta, buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
