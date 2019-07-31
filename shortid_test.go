package goinsta

import (
	"testing"
)

func TestMediaIDFromShortID(t *testing.T) {
	mediaID, err := MediaIDFromShortID("BR_repxhx4O")
	if err != nil {
		t.Fatal(err)
		return
	}
	if mediaID != "1477090425239445006" {
		t.Fatal("Invalid mediaID")
	}
}
