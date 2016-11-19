package UserManager

import (
	"crypto/sha512"
	"encoding/hex"
	"Utils"

)

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
func verifyHash(psw string, salt string,correctHash string) bool {
	var saltedPsw = psw+salt
	shaHasher := sha512.New()
	shaHasher.Write([]byte(saltedPsw))
	var inputHash = hex.EncodeToString(shaHasher.Sum(nil))

	return inputHash == correctHash
}

//VerifyUser - Read User Data and VerifyHash
func VerifyUser(email string, psw string) bool{
	usr, err := ReadUser(email)
	if err != nil {
		Utils.HandlePrint(err) //TODO kp was ab geht
	}
	return verifyHash(psw,usr.Salt,usr.HashedPassword)
}