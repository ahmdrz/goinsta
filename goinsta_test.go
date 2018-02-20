package goinsta

import (
	"encoding/json"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
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
		t.Skip("Username or Password is empty")
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
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestUserFollowings(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	_, err := insta.UserFollowing(insta.LoggedInUser.ID, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestUserFollowers(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	_, err := insta.UserFollowers(insta.LoggedInUser.ID, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestSelfUserFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	_, err := insta.LatestFeed()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestSelfUserFeedWithoutRelogin(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	insta2 := New(username, password)
	_, err := insta2.LatestFeed()
	if err == nil {
		t.Fatal("there is no error, but it must be error cuz no login.")
		return
	}

	u, _ := url.Parse(GOINSTA_API_URL)
	cookies := insta.GetSessions(u)

	insta2.IsLoggedIn = true
	insta2.SetCookies(u, cookies)

	insta2.Informations.DeviceID = insta.Informations.DeviceID
	insta2.Informations.UUID = insta.Informations.UUID
	insta2.Informations.Username = insta.Informations.Username
	insta2.Informations.RankToken = insta.Informations.RankToken
	//insta2.LoggedInUser = insta.LoggedInUser

	resp2, err := insta2.LatestFeed()
	for _, item := range resp2.Items {
		t.Log(item.Code)
	}

	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log(resp2.Status)
}

func TestMediaLikers(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.LatestFeed()
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
	time.Sleep(3 * time.Second)
}

func TestMediaComments(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	resp, err := insta.LatestFeed()
	if err != nil {
		t.Fatal(err)
		return
	}

	if len(resp.Items) > 0 {
		result, err := insta.MediaComments(resp.Items[0].ID, "")
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log(result.Status)
	} else {
		t.Skip("Empty feed")
	}
	time.Sleep(3 * time.Second)
}

func TestFollow(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	user, err := insta.GetUserByUsername("elonmusk")
	if err != nil {
		t.Fatal(err)
		return
	}

	_, err = insta.Follow(user.User.ID)
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestUnFollow(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	user, err := insta.GetUserByUsername("elonmusk")
	if err != nil {
		t.Fatal(err)
		return
	}

	_, err = insta.UnFollow(user.User.ID)
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestFullLikeCommentOptions(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.TagFeed("pizza") // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok -> length : ", len(resp.Items))

	for _, item := range resp.Items {
		id := item.ID

		_, err := insta.Like(id)
		if err != nil {
			t.Fatal(err)
			return
		}
		time.Sleep(3 * time.Second)
		t.Log("[LIKE] Finished")

		_, err = insta.UnLike(id)
		if err != nil {
			t.Fatal(err)
			return
		}
		time.Sleep(3 * time.Second)
		t.Log("[UNLIKE] Finished")

		_, err = insta.MediaInfo(id)
		if err != nil {
			t.Fatal(err)
			return
		}
		time.Sleep(3 * time.Second)
		t.Log("[MEDIAINFO] Finished")

		bytes, err := insta.Comment(id, "Wow <3 pizza !")
		if err != nil {
			t.Fatal(err)
			return
		}

		t.Log("[COMMENT] Finished")

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

		bytes, err = insta.DeleteComment(id, strconv.FormatInt(Result.Comment.ID, 10))
		if err != nil {
			t.Fatal(err)
			return
		}
		time.Sleep(3 * time.Second)
		t.Log("[DELETE] Finished")

		break
	}
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
	time.Sleep(3 * time.Second)
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
	time.Sleep(3 * time.Second)
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
	time.Sleep(3 * time.Second)
	t.Log("Finished")
}

func TestGetUserByID(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.GetUserByID(17644112) // ID of "elonmusk"
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "ok" {
		t.Fatalf("Incorrect status" + resp.Status)
	}

	if resp.User.Username != "elonmusk" {
		t.Fatalf("Username mismatch" + resp.User.Username)
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestGetReelsTray(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.GetReelsTrayFeed()
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "ok" {
		t.Fatalf("Incorrect status" + resp.Status)
	}

	t.Log(resp.Status)
}

func TestGetUserByUsername(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.GetUserByUsername("aidenzibaei")
	if err != nil {
		t.Fatal(err)
	}

	if resp.Status != "ok" {
		t.Fatalf("Incorrect status" + resp.Status)
	}

	if resp.User.Username != "aidenzibaei" {
		t.Fatalf("Incorrect username" + resp.User.Username)
	}
	time.Sleep(3 * time.Second)
	t.Log("ok")
}

func TestGetProfileData(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.GetProfileData()
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(3 * time.Second)
	t.Log("Finished")
}

func TestRecentActivity(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	recentActivity, err := insta.GetRecentActivity()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("recentActivity.Status:%s", recentActivity.Status)
	t.Logf("recentActivity.ContinuationToken:%d", recentActivity.ContinuationToken)
	t.Logf("recentActivity.Counts.PhotosOfYou:%d", recentActivity.Counts.PhotosOfYou)
	t.Logf("recentActivity.Counts.Requests:%d", recentActivity.Counts.Requests)
	if len(recentActivity.OldStories) > 0 {
		for _, item := range recentActivity.OldStories {
			t.Logf("PK=%s, type=%d, text=%s ", item.PK, item.Type, item.Args.Text)
		}
	}

	time.Sleep(3 * time.Second)
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
	time.Sleep(3 * time.Second)
	t.Log("Finished")
}

func TestLatestUserFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.LatestUserFeed(17644112) // ID from elonmusk
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("Finished")
}

func TestUserFeedWithMaxID(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserFeed(17644112, "25025320", "") // ID from elonmusk
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("Finished")
}

func TestUserFeedWithMaxIDAndTimestamp(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserFeed(17644112, "25025320", "25025320") // ID from elonmusk
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log("Finished")
}

func TestUserTaggedFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserTaggedFeed(17644112, 0, "") // ID from elonmusk
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("Finished")
}

func TestUserTaggedFeedWithMaxIDAndTimestamp(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UserTaggedFeed(17644112, 25035320, "25025320") // ID from elonmusk
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
	time.Sleep(3 * time.Second)
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
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestUserFriendShip(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	resp, err := insta.UserFriendShip(insta.LoggedInUser.ID)
	if err != nil {
		t.Fatal(err)
		return
	}

	if resp.Status != "ok" {
		t.Fatal(resp.Status)
		return
	}
	time.Sleep(3 * time.Second)
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

	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

/////////// logout

func TestSyncFeatures(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	err := insta.SyncFeatures()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
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
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestSetPublicAndUnPublicAccount(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	// check account status
	if insta.LoggedInUser.IsPrivate {
		helperSetPublicAccount(t)
		helperSetPrivateAccount(t)
	} else {
		helperSetPrivateAccount(t)
		helperSetPublicAccount(t)
	}

}

func helperSetPublicAccount(t *testing.T) {
	_, err := insta.SetPublicAccount()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(10 * time.Second)
	t.Log("status : ok")
}

func helperSetPrivateAccount(t *testing.T) {
	_, err := insta.SetPrivateAccount()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(10 * time.Second)
	t.Log("status : ok")
}

func TestBlock(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.Block(17644112) // ID of "elonmusk"
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestUnBlock(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.UnBlock(17644112) // ID of "elonmusk"
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

var directThreadID string

func TestGetDirectPendingRequests(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	direct, err := insta.GetDirectPendingRequests()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")

	// store threadID for other test
	if len(direct.Inbox.Threads) > 0 {
		directThreadID = direct.Inbox.Threads[0].ThreadID
	}
}

func TestGetDirectThread(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	if directThreadID == "" {
		t.Skip("Empty Direct ThreadID")
	}
	_, err := insta.GetDirectThread(directThreadID)
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestGetDirectThreadMediaShare(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	if directThreadID == "" {
		t.Skip("Empty Direct ThreadID")
	}
	thread, err := insta.GetDirectThread(directThreadID)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, item := range thread.Thread.Items {
		t.Log(item.MediaShare.TakenAt)
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestExplore(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.Explore()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

// NOT USE THIS TEST!!! Warning! Change the password on the old can not be!
// func TestChangePassword(t *testing.T) {
//         if skip {
//                 t.Skip("Empty username or password , Skipping ...")
//         }
//
//         t.Log("change password... Added \"-\"")
//         _, err := insta.ChangePassword(insta.Informations.Password + "-")
//         if err != nil {
//                 t.Fatal(err)
//                 return
//         }
//
//         time.Sleep(3 * time.Second)
//         t.Log("status : ok")
// }

func TestGetFollowingRecentActivity(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.GetFollowingRecentActivity()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}

func TestGetUserStories(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	_, err := insta.GetUserStories(17644112) // ID of "elonmusk"
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}
func TestLogout(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	err := insta.Logout()
	if err != nil {
		t.Fatal(err)
		return
	}
	time.Sleep(3 * time.Second)
	t.Log("status : ok")
}
