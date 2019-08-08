package tests

import "testing"

func TestFeedTagLike(t *testing.T) {
	insta, err := getRandomAccount()
	if err != nil {
		t.Fatal(err)
		return
	}
	feedTag, err := insta.Feed.Tags("golang")
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, item := range feedTag.RankedItems {
		// media, err := insta.GetMedia(item.ID)
		// if err != nil {
		// 	t.Fatal(err)
		// 	return
		// }
		// err = media.Items[0].Like()

		err = item.Like()
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Logf("media %s liked by goinsta", item.ID)
	}
}

func TestFeedTagNext(t *testing.T) {
	insta, err := getRandomAccount()
	if err != nil {
		t.Fatal(err)
		return
	}
	feedTag, err := insta.Feed.Tags("golang")
	if err != nil {
		t.Fatal(err)
		return
	}

	initNextID := feedTag.NextID
	success := feedTag.Next()
	if !success {
		t.Fatal("Failed to fetch next page")
		return
	}
	gotStatus := feedTag.Status

	if gotStatus != "ok" {
		t.Errorf("Status = %s; want ok", gotStatus)
	}

	gotNextID := feedTag.NextID
	if gotNextID == initNextID {
		t.Errorf("NextID must differ after FeedTag.Next() call")
	}
}
