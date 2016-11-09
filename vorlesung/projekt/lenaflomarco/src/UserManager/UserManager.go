package UserManager

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"github.com/pkg/errors"
)

type usr struct {
	name string
	salt string
	hash string
}

//MakeSalt - Salt generieren
//TODO byte to string error
func makeSalt(numBytes int) (salt string, err error) {
	bytesalt := make([]byte, numBytes)
	_, err = io.ReadFull(rand.Reader, bytesalt)
	if err != nil {
		errors.Wrap(err, "Error in makeSalt while generatig random number")
		return
	}
	salt = string(bytesalt)
	return
}

//SetHash - einen hash Setzten
func SetHash(psw string)(hash string,salt string) {
	var err error

	//salting
	salt , err = makeSalt(16)
	if err != nil {
		return
	}
	var saltedPsw = psw+salt

	//hashing
	shaHasher := sha512.New()
	shaHasher.Write([]byte(saltedPsw))
	hash = hex.EncodeToString(shaHasher.Sum(nil))
	return
}
	/*
	//check if file exists
	if _, err := os.Stat(path); err != nil {
		//create empty file
		lcase := []byte(``)
		perm := os.FileMode(0644)
		err := ioutil.WriteFile(path, lcase, perm)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Read test file
	data, err := ioutil.ReadFile(path) //unit8[]
	if err != nil {
		fmt.Println(err)
		return
	}

	//create hash
	h := sha512.New()
	h.Write([]byte(psw))
	bs := hex.EncodeToString(h.Sum(nil))

	//add to File Text

	//data+/n+u+bs
	userString := "\n" + username + ";" + bs
	dataSlice := data[:]
	dataSlice = append(dataSlice, userString...)

	// write in File
	perm := os.FileMode(0777)
	err = ioutil.WriteFile(path, dataSlice, perm)
	if err != nil {
		fmt.Println(err)
		return
	}
}*/

//VerifyHash - Hash überprüfen
func VerifyHash(psw string, salt string,correctHash string) bool {
	var saltedPsw = psw+salt

	shaHasher := sha512.New()
	shaHasher.Write([]byte(saltedPsw))
	var inputHash = hex.EncodeToString(shaHasher.Sum(nil))

	return inputHash == correctHash
}
/*
	input, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return false
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, username) {
			savedHash = strings.Split(lines[i], ";")[1]
		}
	}

	//create hash
	h.Write([]byte(psw))
	bs := hex.EncodeToString(h.Sum(nil))

	if savedHash == bs {
		return true
	}
	return false
}
*/