package device

import "testing"

func TestNew(t *testing.T) {
	d1, err := New("lg_g5")
	if err != nil {
		t.Fatal(err)
	}
	d2, err := New("lg_g5")
	if err != nil {
		t.Fatal(err)
	}
	d1.CPU = "new cpu name"
	if d2.CPU == d1.CPU {
		t.Fatal("d1 cpu name changed")
	}
}
