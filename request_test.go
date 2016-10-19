package goinsta

import (
	"os"
	"testing"
)

func TestRequest(t *testing.T) {
	username := os.Getenv("INSTA_USERNAME")
	password := os.Getenv("INSTA_PASSWORD")
	insta := New(username, password)
	err := insta.sendRequest("si/fetch_headers/?challenge_type=signup&guid="+generateUUID(false), "", true)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status : ok")
}
