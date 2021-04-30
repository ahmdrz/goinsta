package goinsta

import (
	"testing"
)

func TestMediaIDFromShortID(t *testing.T) {
	mediaID := MediaIDFromShortID("BR_repxhx4O")
	if mediaID != "1477090425239445006" {
		t.Fatal("Invalid mediaID")
	}
}
