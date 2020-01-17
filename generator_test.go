package goinsta

import (
	"regexp"
	"testing"
)

func TestGenerateMD5Hash(t *testing.T) {
	testCases := [][]string{
		[]string{"test", "098f6bcd4621d373cade4e832627b4f6"},
		[]string{"hello", "5d41402abc4b2a76b9719d911017c592"},
	}
	for _, pair := range testCases {
		hash := generateMD5Hash(pair[0])
		if hash != pair[1] {
			t.Fatalf("got %s, want %s", hash, pair[1])
		}
	}
}

func TestGenerateUUID(t *testing.T) {
	a := generateUUID()
	b := generateUUID()
	if a == b {
		t.Fatalf("two uuids are same")
	}
	matched, err := regexp.MatchString(`[a-f0-9\\-]*`, a)
	if err != nil {
		t.Fatal(err)
	}
	if !matched {
		t.Fatalf("uuid is not valid")
	}
}
