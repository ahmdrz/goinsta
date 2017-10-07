package store

import (
	"os"
	"testing"
	"time"

	"github.com/ahmdrz/goinsta"
)

func TestExportImport(t *testing.T) {
	username := os.Getenv("INSTA_USERNAME")
	password := os.Getenv("INSTA_PASSWORD")
	if len(username)*len(password) == 0 && os.Getenv("INSTA_PULL") != "true" {
		t.Skip("Username or Password is empty")
	}

	var key = []byte("RH1tCpR80AQ3WzXJ") //32byte key for AES

	var encoded_string string

	{
		insta := goinsta.New(username, password)
		insta.Login()
		bytes, err := Export(insta, key)
		if err != nil {
			t.Fatal("Error on export")
		}
		encoded_string = string(bytes)
		insta.Logout()
	}

	time.Sleep(3 * time.Second)

	{
		insta, err := Import([]byte(encoded_string), key)
		if err != nil {
			t.Fatal("Error on import")
		}
		_, err = insta.GetUserByUsername("elonmusk")
		if err != nil {
			t.Fatal("search username")
		}
		insta.Logout()
	}

	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}
