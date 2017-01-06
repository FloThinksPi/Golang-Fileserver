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
	"sync"
	"Flags"
	"Utils"
)

//Workdir in which the UserDatabase gets Written and Users uploaded Files are Stored
var workdir = Flags.GetWorkDir()

//The Data which is accessed by Go should be loaded in RAM(Hashmap for best reading performance) and only synced to disk if a change is imminent/done.
type UserMap map[string]UserRecord

//RWmutex for concurrent access and prevention of Mutual Exclusion, the mutex should be unlocked if data got written to permanent storage
type UserStorage struct {
	UserMap
	sync.RWMutex
}

//Entry of a single user
type UserRecord struct {
	UID            int16  //Unique ID
	Email          string //Email (Must be Unique!)
	Name           string //Name
	HashedPassword string //Password Hashed and Salted
	Salt           string //Salt
}

//Actual Global acessible Variable for saving the user Data
var managersUserStorage UserStorage

func init() {
	managersUserStorage.Lock()
	defer managersUserStorage.Unlock()

	//Create Map
	Utils.LogDebug("Init UserDatabase")
	var err error = nil

	managersUserStorage.UserMap = make(UserMap)
	managersUserStorage.UserMap,err = readDataToMemory(Flags.GetWorkDir()+"/userdatabase")

	Utils.HandlePrint(err)
}