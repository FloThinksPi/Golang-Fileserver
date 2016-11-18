package SessionManager

import (
	"github.com/pkg/errors"
	"reflect"
	"time"
	"Utils"
)

//NewSession Checks if a supplied SessionRecord is completely filled, reqeust a random SessionID and adds this to the SessionRecord.
//The SessionRecord is then saved in the packageGlobal managersSessionStorage.
func NewSession(record SessionRecord) (err error) {
	managersSessionStorage.RWMutex.Lock()
	defer managersSessionStorage.RWMutex.Unlock()

	record.Session= Utils.RandString(SESSION_KEY_LENGTH)

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

//ValidateSession checks if a supplied SessionID exists in managersSessionStorage and returns true if the Session didnt time out.
func ValidateSession(session string) (valid bool, err error) {
	managersSessionStorage.RWMutex.RLock()
	defer managersSessionStorage.RWMutex.RUnlock()

	record := managersSessionStorage.SessionMap[session] //TODO ADD error handling

	// Valid if Session is found in storage and Session didnt timeout!
	valid = record.Session == session && record.SessionLast.After(time.Now().Add( time.Duration(-*sessionTimeout) * time.Second))

	return
}
