package UserManager

import "testing"
import "strings"
import "github.com/stretchr/testify/assert"

const (
	TestFile = ""
)

func TestMakeSalt(t *testing.T) {
	salt, err := MakeSalt(16)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 16, strings.Count(salt, ""), "Salt is not Generating 16 Characters")
}

func TestSetHash(t *testing.T) {
	t.SkipNow()
}

//TestVerifyHash
func TestVerifyHash(t *testing.T) {
	var v bool
	v = VerifyHash("datastorage/userdatabase", "flo", "meinPasswort")
	if v {
		t.Error("Expected false , got ", v)
	}
}
