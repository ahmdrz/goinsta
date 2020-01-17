package goinsta

import "testing"

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
}
