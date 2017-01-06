/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

package SessionManager

import (
	"sync"
	"time"
	"flag"
)


// SESSION_KEY_LENGTH Length of the SessionKey which gets used after Authentication for access control.
const (
	SESSION_KEY_LENGTH = 64
)

// sessionTimeout is a Variable wich holds the timeout in Second and can be set by Flag e.g -sessionTimeout=900
var sessionTimeout = flag.Int("sessionTimeout", 900, "Time in Seconds after which a inactive Session gets invalid")

//SessionMap arranges SessionRecords in a Map for faster Access by SessionID
type SessionMap map[string]SessionRecord

//SessionStorage RWmutex for concurrent access and prevention of Mutual Exclusion. Holds A SessionMap and adds Mutex functionality
type SessionStorage struct {
	SessionMap
	sync.RWMutex
}

//SessionRecord Entry of a single user.
//Email is unique and identifies a user.
//Session is the sessionKey which is used for further authentication.
//SessionLast is a timestamp of the last access to the SessionRecord, this is used to implement a timeout functionality.
type SessionRecord struct {
	Email       string    //Email
	Session     string    //SessionCookie
	SessionLast time.Time //Timestamp when last Interaction took place -> for session timeout
}

//Actual acessible Variable for saving the user Data
var managersSessionStorage SessionStorage

//Init managersSessionStorage and parse Flags
func init() {
	managersSessionStorage.Lock()
	defer managersSessionStorage.Unlock()

	//Init Map
	managersSessionStorage.SessionMap = make(SessionMap)

}