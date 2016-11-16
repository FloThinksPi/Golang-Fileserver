package SessionManager

import (
	"github.com/pkg/errors"
	"reflect"
	"time"
	"math/rand"
)

//TODO Fix copy of code from UserManager
var src = rand.NewSource(time.Now().UnixNano())
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!§$%&/()=?"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generateSessionKey(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n - 1, src.Int63(), letterIdxMax; i >= 0; {
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

func NewSession(record SessionRecord) (err error) {
	managersSessionStorage.RWMutex.Lock()
	defer managersSessionStorage.RWMutex.Unlock()

	record.Session=generateSessionKey(SESSION_KEY_LENGTH)

	//TODO cause error if value is nil and dont add user
	v := reflect.ValueOf(record)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil {
			errors.WithMessage(err, "A Field of the userrecord was Nil")
			return
		}
	}

	managersSessionStorage.SessionMap[record.Session] = record //TODO ADD error handling
	return
}


func ValidateSession(session string) (valid bool, err error) {
	managersSessionStorage.RWMutex.RLock()
	defer managersSessionStorage.RWMutex.RUnlock()

	record := managersSessionStorage.SessionMap[session] //TODO ADD error handling

	// Valid if Session is found in storage and Session didnt timeout!
	valid = record.Session == session && record.SessionLast.After(time.Now().Add(-SESSION_TIMEOUT_IN_SECODS * time.Second))

	return
}
