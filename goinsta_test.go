package goinsta

import (
	"os"
	"testing"
)

var (
	username string
	password string
	insta    *Instagram
	skip     bool
)

func TestHandlesNonExistingItems(t *testing.T) {
	username = os.Getenv("INSTA_USERNAME")
	password = os.Getenv("INSTA_PASSWORD")
	if len(username)*len(password) == 0 {
		skip = true
		t.Fatal("Username or Password is empty")
	}
}

func TestDeviceID(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	insta = New(username, password)
	t.Log(insta.Informations.DeviceID)
}

func TestLogin(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	err := insta.Login()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("Logged in user", insta.LoggedInUser.FullName)
}

func TestUserFollowings(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	bytes, err := insta.UserFollowings(insta.Informations.UsernameId, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes)[:100])
}
