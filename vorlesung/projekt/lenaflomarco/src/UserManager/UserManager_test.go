package UserManager

import "testing"
import "strings"
import "github.com/stretchr/testify/assert"

const (
	TestFile = ""
)

func TestGenerateHash(t *testing.T)  {
	hash,salt := GenerateHash("MeinPasswort")
	hash2,salt2 := GenerateHash("MeinPasswort")
	assert.NotEqual(t,hash,hash2,"Error, Two Hashes of the same Passwort are identical")
	assert.NotEqual(t,salt,salt2,"Error, Two Runs of generate hash returned the same Salt")
}

func TestMakeSalt(t *testing.T) {
	salt, err := makeSalt(16)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 16, strings.Count(salt, ""), "Salt is not Generating 16 Characters")
}

//TestVerifyHash
func TestVerifyHash(t *testing.T) {
	var v bool
	v = VerifyHash("Psw","ABC","25928498b28c3268d911dd78d7ff820e0f14ed32b7ac2d397746f1778038b968d9e6364fd4b3da2e7026bdf574c104779fac9ce9064b6b9ae09ac043f8d131d4")
	if !v {
		t.Error("Expected true , got ", v)
	}
}
