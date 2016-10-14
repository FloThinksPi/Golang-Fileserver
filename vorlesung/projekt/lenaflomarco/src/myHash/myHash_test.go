package myHash

import "testing"
import "strings"

const (
	TestFile = "usr.txt"
	TestUser = "Flo"
	TestPsw  = "PSW"
)

func TestMakeSalt(t *testing.T) {
	var salt string
	salt = MakeSalt(16)
	if strings.Count(salt, "") != 16 {
		t.Error("Salt is not Generating 16 Characters")
	}
}

func TestSetHash(t *testing.T) {
	t.SkipNow()
}

//TestVerifyHash
func TestVerifyHash(t *testing.T) {
	var v bool
	v = VerifyHash("usr.txt", "flo", "meinPasswort")
	if !v {
		t.Error("Expected 1.5, got ", v)
	}
}
