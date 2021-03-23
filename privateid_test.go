package goinsta

import (
	"testing"
)

func TestMediaIDFromPrivateID(t *testing.T) {
	mediaID, err := MediaIDFromPrivateID("CFPU5jqA7F7gtdLtpBevHl5A6rVwt3fi4zVzs00")
	if err != nil {
		t.Fatal(err)
		return
	}
	if mediaID != "2400229042638008699" {
		t.Fatal("Invalid mediaID")
	}
}
