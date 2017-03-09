package goinsta

import (
	"encoding/json"
	"net/url"
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

func TestSelfUserFeedWithoutRelogin(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	insta2 := New(username, password)
	_, err := insta2.UserFeed()
	if err == nil {
		t.Fatal("there is no error, but it must be error cuz no login.")
		return
	}

	u, _ := url.Parse(GOINSTA_API_URL)
	cookies := insta.GetSessions(u)

	insta2.IsLoggedIn = true
	insta2.SetCookies(u, cookies)
	insta2.Informations.UsernameId = insta.Informations.UsernameId

	insta2.Informations.DeviceID = insta.Informations.DeviceID
	insta2.Informations.UUID = insta.Informations.UUID
	insta2.Informations.Username = insta.Informations.Username
	insta2.Informations.RankToken = insta.Informations.RankToken
	//insta2.LoggedInUser = insta.LoggedInUser

	resp2, err := insta2.UserFeed()
	for _, item := range resp2.Items {
		t.Log(item.Code)
	}

	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(resp2.Status)
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

	resp, err := insta.Follow(strconv.Itoa(user.User.Pk))
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

	resp, err := insta.UnFollow(strconv.Itoa(user.User.Pk))
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

func TestSearchLocation(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	res, err := insta.SearchLocation("37.3874", "122.0575", "大阪")

	if err != nil {
		t.Fatal(err)
		return
	}

	for i, venue := range res.Venues {
		t.Logf("%d: name=%s, address=%s, lat=%f, lng=%f ", i, venue.Name, venue.Address, venue.Lat, venue.Lng)
	}
}

func TestGetLocationFeed(t *testing.T) {

	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	locationFeed, err := insta.GetLocationFeed(108164709212336, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	for i, item := range locationFeed.RankedItems {
		t.Logf("%d: code=%s", i, item.Code)
	}

	for i, item := range locationFeed.Items {
		t.Logf("%d: code=%s", i, item.Code)
	}

	t.Log("Finished")
}

func TestTagRelated(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	tags, err := insta.GetTagRelated("student")
	if err != nil {
		t.Fatal(err)
		return
	}

	for i, tag := range tags.Related {
		t.Logf("%d: name=%s", i, tag.Name)
	}

	t.Log("Finished")
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
		t.Fatal("Incorrect format for comment")
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

func TestFirstUserFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.FirstUserFeed("25025320")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestUserFeedWithMaxID(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserFeed("25025320")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestUserFeedWithMaxIDAndTimestamp(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserFeed("25025320", "25025320")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestUserFeedWithToManyArgs(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserFeed("", "", "", "")
	if err == nil {
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

func TestGetSessions(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	url, _ := url.Parse("https://instagram.com")
	cookies := insta.GetSessions(url)
	for _, cookie := range cookies {
		t.Log(generateMD5Hash(cookie.String()))
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

func TestUserFriendShip(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.UserFriendShip(insta.Informations.UsernameId)
	if err != nil {
		t.Fatal(err)
		return
	}

	if resp.Status != "ok" {
		t.Fatal(resp.Status)
		return
	}

	t.Log("status : ok")
}

func TestGetPopularFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.GetPopularFeed()
	if err != nil {
		t.Fatal(err)
		return
	}

	if resp.Status != "ok" {
		t.Fatal(resp.Status)
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

func TestSyncFeatures(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	err := insta.SyncFeatures()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status : ok")
}

func TestAutoCompleteUserList(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	err := insta.AutoCompleteUserList()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status : ok")
}

func TestMegaphoneLog(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	err := insta.MegaphoneLog()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("status : ok")
}
