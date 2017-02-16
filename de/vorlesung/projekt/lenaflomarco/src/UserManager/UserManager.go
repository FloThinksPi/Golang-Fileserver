/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

package UserManager

import (
	"crypto/sha512"
	"encoding/hex"
	"Utils"
	"Flags"
	"strconv"
	"os"
)

//SetHash - einen hash Setzten
func generatePasswordHash(psw string)(hash string,salt string) { //TODO die bcyrpt funktion nutzend da stadart und sicher !
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

func RegisterUser(name, email, password string) bool {
	record, present, _ := ReadUser(email)
	if(present) {
		Utils.LogDebug("Benutzer existiert bereits. Es wird kein neuer Nutzer angelegt")
		return false
	}
	uid := getNextUID()
	hash, salt := generatePasswordHash(password)
	record = UserRecord{
		UID:uid,
		Email:email,
		Name:name,
		HashedPassword:hash,
		Salt:salt}
	writeUser(record, Flags.GetWorkDir())	//TODO Error Handling

	Utils.LogDebug("UID (Foldername): "+strconv.Itoa(int(uid)))
	err := os.Mkdir(workdir + "/" + strconv.Itoa(int(uid)), 0777)

	if err != nil {
		Utils.LogError("Error in creating user directory")
		return false
	}
	return true
}

func ChangePassword(email, passwordNew string)  {
	usr, _,_ := ReadUser(email)
	hash, salt := generatePasswordHash(passwordNew)
	usr.HashedPassword = hash
	usr.Salt = salt
	writeUser(usr, Flags.GetWorkDir())
}