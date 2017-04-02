package goinsta

import (
	"testing"
)

func TestMediaFromCode(t *testing.T) {
	mediaID := MediaFromCode("BR_repxhx4O")
	if mediaID == "1477090425239445006" {
		t.Log(mediaID)
	} else {
		t.Fatal("Invalid mediaID")
	}
}
