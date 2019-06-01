package utilities

import (
	"bytes"

	"github.com/ahmdrz/goinsta"
)

func ImportFromBytes(inputBytes []byte) (*goinsta.Instagram, error) {
	return goinsta.ImportReader(bytes.NewReader(inputBytes))
}
