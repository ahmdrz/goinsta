package goinsta

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

var (
	username    string
	password    string
	pullrequest bool
	insta       *Instagram
	skip        bool
)

func TestHandlesNonExistingItems(t *testing.T) {
	username = os.Getenv("INSTA_USERNAME")
	password = os.Getenv("INSTA_PASSWORD")
	pullrequest = os.Getenv("INSTA_PULL") == "true"

	t.Log("Pull Request", pullrequest)

	if len(username)*len(password) == 0 && !pullrequest {
		skip = true
		t.Fatal("Username or Password is empty")
	}
	skip = pullrequest
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
	t.Log("status : ok")
}

func TestUserFollowings(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.UserFollowing(insta.Informations.UsernameId, "")
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
		result, err := insta.MediaLikers(resp.Items[0].ID)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log(result.Status)
	} else {
		t.Skip("Empty feed")
	}
}

func TestFollow(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	user, err := insta.GetUsername("elonmusk")
	if err != nil {
		t.Fatal(err)
		return
	}

	resp, err := insta.Follow(user.User.StringID())
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(resp.Status)
}

func TestUnFollow(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	user, err := insta.GetUsername("elonmusk")
	if err != nil {
		t.Fatal(err)
		return
	}

	resp, err := insta.UnFollow(user.User.StringID())
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(resp.Status)
}

func TestLike(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.Like("1363799876794028707") // random image ! from search by tags on pizza
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestMediaInfo(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.MediaInfo("1363799876794028707") // random image ! from search by tags on pizza
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

func TestCommentAndDeleteComment(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.Comment("1363799876794028707", "Wow <3 pizza !") // random image ! from search by tags on pizza
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")

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

	bytes, err = insta.DeleteComment("1363799876794028707", strconv.FormatInt(Result.Comment.ID, 10))
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestGetUserID(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.GetUserID("17644112") // ID of "elonmusk"
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "ok" {
		t.Fatalf("Incorrect status" + resp.Status)
	}

	if resp.User.Username != "elonmusk" {
		t.Fatalf("Username mismatch" + resp.User.Username)
	}

	t.Log(resp.Status)
}

func TestGetUsername(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.GetUsername("ahmd.rz")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "ok" {
		t.Fatalf("Incorrect status" + resp.Status)
	}

	if resp.User.Username != "ahmd.rz" {
		t.Fatalf("Incorrect username" + resp.User.Username)
	}

	t.Log(resp.Status)
}

func TestGetProfileData(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.GetProfileData()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Finished")
}

func TestRecentActivity(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.GetRecentActivity()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Finished")
}

func TestSearchUsername(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.SearchUsername("ahmd.rz")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestSearchTags(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.SearchTags("instagram")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestGetLastJson(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.SearchTags("instagram")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestGetSessions(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	m := insta.GetSessions()
	for _, session := range m {
		for _, cookie := range session {
			t.Log(generateMD5Hash(cookie.String()))
		}
	}
}

func TestExpose(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	err := insta.Expose()
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("status : ok")
}

/////////// logout

func TestLogout(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	err := insta.Logout()
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("status : ok")
}
