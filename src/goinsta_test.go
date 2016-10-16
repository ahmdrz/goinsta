package goinsta

import (
	"encoding/json"
	"os"
	"strconv"
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
	t.Log("Logged in user", insta.LoggedInUser.FullName, insta.Informations.UsernameId)
}

func TestUserFollowings(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.UserFollowings(insta.Informations.UsernameId, "")
	if err != nil {
		t.Log(insta.GetLastJson())
		t.Fatal(err)
		return
	}
	t.Log(resp.Status)
}

func TestUserFollowers(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.UserFollowers(insta.Informations.UsernameId, "")
	if err != nil {
		t.Log(insta.GetLastJson())
		t.Fatal(err)
		return
	}
	t.Log(resp.Status)
}

func TestSelfUserFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.UserFeed()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(resp.Status)
}

func TestMediaLikers(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.UserFeed()
	if err != nil {
		t.Fatal(err)
		return
	}

	if len(resp.Items) > 0 {
		bytes, err := insta.MediaLikers(resp.Items[0].ID)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log(string(bytes)[:30])
	} else {
		t.Skip("Empty feed")
	}
}

func TestFollow(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.Follow("1572292791") // ahmdrz (creator) instagram usernameID
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}

func TestLike(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.Like("1336846574982263293") // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}

func TestMediaInfo(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.MediaInfo("1336846574982263293") // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(resp.Status)
}

func TestTagFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.TagFeed("pizza") // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(resp.Items[0])
}

func TestSetPrivate(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.SetPrivateAccount()
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}

func TestCommentAndDeleteComment(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.Comment("1336846574982263293", "Hello , I'm your Instagram Bot !") // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))

	type Comment struct {
		ID int64 `json:"pk"`
	}

	var Result struct {
		Comment Comment `json:"comment"`
		Status  string  `json:"status"`
	}

	err = json.Unmarshal(bytes, &Result)
	if err != nil {
		t.Fatal(err)
		return
	}

	if Result.Status != "ok" {
		t.Fatalf("Incorrect format for comment")
		return
	}

	bytes, err = insta.DeleteComment("1336846574982263293", strconv.FormatInt(Result.Comment.ID, 10)) // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}

func TestSearchUsername(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.SearchUsername("ahmd.rz")
	if err != nil {
		t.Fatal(err)
	}

	type User struct {
		Id       int64  `json:"pk"`
		Username string `json:"username"`
	}

	var Result struct {
		Status string `json:"status"`
		User   User   `json:"user"`
	}

	err = json.Unmarshal(bytes, &Result)
	if err != nil {
		t.Fatal(err)
	}

	if Result.Status != "ok" {
		t.Fatalf("Incorrect status" + Result.Status)
	}

	if Result.User.Username != "ahmd.rz" {
		t.Fatalf("Incorrect username" + Result.User.Username)
	}

	t.Log(Result)
}

func TestGetProfileData(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.GetProfileData()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}
