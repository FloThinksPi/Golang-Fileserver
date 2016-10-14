package myHash

import "testing"


const(
	TEST_FILE = "usr.txt"
	TEST_USER = "Flo"
	TEST_PSW = "PSW"
)

func TestMakeSalt(t *testing.T) {
		t.Error("Not Implemented")
}

func TestSetHash(t *testing.T) {

}

//TestVerifyHash
func TestVerifyHash(t *testing.T) {
	var v bool
	v = VerifyHash("usr.txt", "flo", "meinPasswort")
	if !v{
		t.Error("Expected 1.5, got ", v)
	}
}

func TestUsernameExists(){

}

func TestSaltExists()  {

}