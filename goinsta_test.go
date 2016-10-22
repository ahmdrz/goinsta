package goinsta

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
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
	t.Log("status : ok , " + insta.LoggedInUser.FullName)
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

	bytes, err := insta.Like("1363799876794028707") // random image ! from search by tags on pizza
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

func TestSetPublic(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.SetPublicAccount()
	if err != nil {
		if strings.Contains(err.Error(), "too many requests") {
			t.Log("too many requests")
			return
		}
		t.Fatal(err)
		return
	}

	t.Log(resp.Status)
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

	bytes, err = insta.DeleteComment("1363799876794028707", strconv.FormatInt(Result.Comment.ID, 10))
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
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

	resp, err := insta.GetProfileData()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp.User.FullName)
}

func TestRecentActivity(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.GetRecentActivity()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bytes))
}

func TestSearchUsername(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.SearchUsername("ahmdrz")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}

func TestSearchTags(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.SearchTags("instagram")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
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

	t.Log(insta.GetLastJson())
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

func TestRemoveProfilePicture(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.RemoveProfilePicture()
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(resp.Status)
}

func TestSetPrivate(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.SetPrivateAccount()
	if err != nil {
		if strings.Contains(err.Error(), "too many requests") {
			t.Log("too many requests")
			return
		}
		t.Fatal(err)
		return
	}

	t.Log(resp.Status)
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
