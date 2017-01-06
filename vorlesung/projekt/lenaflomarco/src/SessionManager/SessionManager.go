package SessionManager

import (
	"github.com/pkg/errors"
	"reflect"
	"time"
)

//NewSession Checks if a supplied SessionRecord is completely filled, reqeust a random SessionID and adds this to the SessionRecord.
//The SessionRecord is then saved in the packageGlobal managersSessionStorage.
func NewSession(record SessionRecord) (err error) {
	managersSessionStorage.RWMutex.Lock()
	defer managersSessionStorage.RWMutex.Unlock()

	//TODO cause error if value is nil and dont add user
	v := reflect.ValueOf(record)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil {
			errors.WithMessage(err, "A Field of the userrecord was Nil")
			return
		}
	}

	managersSessionStorage.SessionMap[record.Session] = record
	return
}

//ValidateSession checks if a supplied SessionID exists in managersSessionStorage and returns true if the Session didnt time out.
func ValidateSession(session string) (valid bool) {
	managersSessionStorage.RWMutex.RLock()
	defer managersSessionStorage.RWMutex.RUnlock()

	record,present := managersSessionStorage.SessionMap[session]

	// Valid if Session is found in storage and Session didnt timeout!
	date := record.SessionLast.After(time.Now().Add(time.Duration(-*sessionTimeout) * time.Second))
	valid = present && date

	return
}

//Invalidates a Session by setting SessionLast so that session counts as expired
func InvalidateSession(session string) (err error) {
	managersSessionStorage.RWMutex.RLock()
	defer managersSessionStorage.RWMutex.RUnlock()
	err = nil

	record, present := managersSessionStorage.SessionMap[session]
	if present {
		record.SessionLast = time.Now().Add(time.Duration(-*sessionTimeout+1) * time.Second)
	}else{
		err = errors.New("Session not Found to Invalidate")
	}

	return
}

func GetSessionRecord(session string) (SessionRecord,bool) {
	record, present := managersSessionStorage.SessionMap[session]
	return record,present
}