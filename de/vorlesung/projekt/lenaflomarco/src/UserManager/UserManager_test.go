/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

package UserManager

import "testing"
import "github.com/stretchr/testify/assert"

//TestGenerateHash Tests if two exact same inputs in the hash function get the same output.
func TestGenerateHash(t *testing.T) {
	hash, salt := generatePasswordHash("MeinPasswort")
	hash2, salt2 := generatePasswordHash("MeinPasswort")
	assert.NotEqual(t, hash, hash2, "Error, Two Hashes of the same Passwort are identical")
	assert.NotEqual(t, salt, salt2, "Error, Two Runs of generate hash returned the same Salt")
}

//TestVerifyHash Test if a specific Plaintext+Salt = a specific hash
func TestVerifyHash(t *testing.T) {
	var v bool
	v = verifyPasswordHash("Psw", "ABC", "25928498b28c3268d911dd78d7ff820e0f14ed32b7ac2d397746f1778038b968d9e6364fd4b3da2e7026bdf574c104779fac9ce9064b6b9ae09ac043f8d131d4")
	if !v {
		t.Error("Expected true , got ", v)
	}
}

func TestVerifyUser(t *testing.T) {


}
