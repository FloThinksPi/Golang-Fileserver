package UserManager

import (
	"math/rand"
	"crypto/sha512"
	"encoding/hex"
	"time"
)

type usr struct {
	name string
	salt string
	hash string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func makeSalt(n int)  string {
	b := make([]byte, n-1)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-2, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

//MakeSalt - Salt generieren

/*func makeSalt(numBytes int) (salt string, err error) {


	bytesalt := make([]byte, numBytes)
	_, err = io.ReadFull(rand.Reader, bytesalt)
	if err != nil {
		errors.Wrap(err, "Error in makeSalt while generatig random number")
		return
	}
	salt = hex.EncodeToString(bytesalt)
	return
}
*/

//SetHash - einen hash Setzten
func SetHash(psw string)(hash string,salt string) {
	//salting
	salt = makeSalt(16)
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