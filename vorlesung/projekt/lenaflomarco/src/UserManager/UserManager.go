package UserManager

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type usr struct {
	name string
	salt string
	hash string
}

const (
	//PW_SALT_BYTES = Length of Salt to add
	PW_SALT_BYTES = 32
)


//MakeSalt - Salt generieren
func MakeSalt() []byte {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil
	}
	return salt
}

//SetHash - einen hash Setzten
func SetHash(path string, username string, psw string) {
	//check if name already

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
}

//VerifyHash - Hash überprüfen
func VerifyHash(path string, username string, psw string) bool {
	h := sha512.New()
	var savedHash string

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
