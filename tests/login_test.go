package tests

import (
	"testing"
)

func TestImportAccount(t *testing.T) {
	insta, err := getRandomAccount()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("logged into Instagram as user '%s'", insta.Account.Username)
}
