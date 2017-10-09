package uuid

import (
	"regexp"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.Fatal(err)
	}
	if IsValidUUID(uuid) {
		t.Log(uuid)
	} else {
		t.Fatal("Invalid uuid")
	}
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}")
	return r.MatchString(uuid)
}
