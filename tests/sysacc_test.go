package tests

import (
	"bytes"
	"encoding/base64"
	"errors"
	"math/rand"
	"os"
	"strings"

	"github.com/ahmdrz/goinsta"
)

func readFromBase64(base64EncodedString string) (*goinsta.Instagram, error) {
	base64Bytes, err := base64.StdEncoding.DecodeString(base64EncodedString)
	if err != nil {
		return nil, err
	}
	return goinsta.ImportReader(bytes.NewReader(base64Bytes))
}

func availableEncodedAccounts() []string {
	output := make([]string, 0)

	environ := os.Environ()
	for _, env := range environ {
		if strings.HasPrefix(env, "INSTAGRAM_BASE64_") {
			index := strings.Index(env, "=")
			encodedString := env[index+1:]
			output = append(output, encodedString)
		}
	}

	return output
}

func getRandomAccount() (*goinsta.Instagram, error) {
	accounts := availableEncodedAccounts()
	if len(accounts) == 0 {
		return nil, errors.New("there is no encoded account in environ")
	}

	encodedAccount := accounts[rand.Intn(len(accounts))]
	return readFromBase64(encodedAccount)
}
