package goinsta

import (
	"os"
	"testing"
)

func TestRequest(t *testing.T) {
	username := os.Getenv("INSTA_USERNAME")
	password := os.Getenv("INSTA_PASSWORD")
	if len(username)*len(password) == 0 {
		t.Skip("Empty username or password , Skipping ...")
	}
	insta := New(username, password)
	err := insta.Login()
	if err != nil {
		t.Fatal(err)
		return
	}

	_, err = insta.sendRequest(&reqOptions{Endpoint: "accounts/logout/"})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status : ok")
}
