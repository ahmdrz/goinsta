package goinsta

import (
	"testing"
)

func TestMD5Hash(t *testing.T) {
	if generateMD5Hash("ahmdrz") != "fb314bb076db69a17670a74cfa0f68f7" {
		t.Fatal("status : failed")
	}
	t.Log("status : ok")
}

func TestHMAC(t *testing.T) {
	if generateHMAC("ahmdrz", "test") != "b669bc2cce9c24f4bdec018430943cbea3a2c8523606305dd31512bc1c5d565e" {
		t.Fatal("status : failed")
	}
	t.Log("status : ok")
}
