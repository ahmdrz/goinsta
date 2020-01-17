package goinsta

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var err error
	_, err = New("username", "password")
	if err != nil {
		t.Fatal(err)
	}
	_, err = New("username", "t")
	if err != ErrBadPassword {
		t.Fatal(err)
	}
	_, err = New("((hello))", "password")
	if err != ErrInvalidUsername {
		t.Fatal(err)
	}
	ig, err := New(os.Getenv("INSTAGRAM_USERNAME"), os.Getenv("INSTAGRAM_PASSWORD"))
	if err != nil {
		t.Fatal(err)
	}
	err = ig.Login()
	if err != nil {
		t.Fatal(err)
	}
}
