package UserManager

import (
	"math/rand"
	"crypto/sha512"
	"encoding/hex"
	"time"
	"Utils"
)

type usr struct {
	name string
	salt string
	hash string
}

//SetHash - einen hash Setzten
func SetHash(psw string)(hash string,salt string) {
	//salting
	salt = Utils.RandString(16)
	var saltedPsw = psw+salt

	//hashing
	shaHasher := sha512.New()
	shaHasher.Write([]byte(saltedPsw))
	hash = hex.EncodeToString(shaHasher.Sum(nil))
	return
}


//VerifyHash - Hash überprüfen
func VerifyHash(psw string, salt string,correctHash string) bool {
	var saltedPsw = psw+salt

	shaHasher := sha512.New()
	shaHasher.Write([]byte(saltedPsw))
	var inputHash = hex.EncodeToString(shaHasher.Sum(nil))

	return inputHash == correctHash
}
