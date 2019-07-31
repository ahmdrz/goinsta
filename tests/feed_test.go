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
