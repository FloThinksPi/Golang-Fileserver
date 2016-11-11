package SessionManager

import (
	"sync"
	"time"
)

//SessionMap

type SessionMap map[string] SessionRecord

//RWmutex for concurrent access and prevention of Mutual Exclusion,
// the mutex should be unlocked if data got written to permanent storage
type SessionStorage struct {
	SessionMap
	sync.RWMutex
}

//Entry of a single user
type SessionRecord struct {
	UID            int16  //Unique ID
	Email          string //Email
	Session        string //SessionCookie
	SessionLast    time.Time   //Timestamp when last Interaction took place -> for session timeout
}

//Actual Global acessible Variable for saving the user Data
var managersSessionStorage SessionStorage

