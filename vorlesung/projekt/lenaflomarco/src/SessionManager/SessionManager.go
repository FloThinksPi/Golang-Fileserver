package SessionManager

import (
	"github.com/pkg/errors"
	"reflect"
)

func SetSession(record SessionRecord) (err error) {
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

	managersSessionStorage.SessionMap[record.Session] = record //TODO ADD error handling
	return
}

func ValidateSession(session string) (valid bool, err error) {
	managersSessionStorage.RWMutex.RLock()
	defer managersSessionStorage.RWMutex.RUnlock()

	record := managersSessionStorage.SessionMap[session] //TODO ADD error handling
	valid = record.Session == session

	return
}
