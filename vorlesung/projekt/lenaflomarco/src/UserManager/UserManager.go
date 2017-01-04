package UserManager

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"Utils"

)

//SetHash - einen hash Setzten
func GeneratePasswordHash(psw string)(hash string,salt string) { //TODO die bcyrpt funktion nutzend da stadart und sicher !
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
func verifyPasswordHash(psw string, salt string,correctHash string) bool {
	var saltedPsw = psw+salt
	shaHasher := sha512.New()
	shaHasher.Write([]byte(saltedPsw))
	var inputHash = hex.EncodeToString(shaHasher.Sum(nil))

	return inputHash == correctHash
}

//VerifyUser - Read User Data and VerifyHash
func VerifyUser(email string, psw string) bool{
	usr, present,err := ReadUser(email)
	Utils.HandlePrint(err)

	// User Exists ?
	if present {
		// Wrong Password?
		return verifyPasswordHash(psw,usr.Salt,usr.HashedPassword)
	}else {
		return false
	}

}

//basicAuth - Checks submitted user credentials and grants access to handler
func basicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok ||  !VerifyUser(user, pass) {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		handler(w, r)
	}
}