package UserManager

import (
	"encoding/json"
	"io/ioutil"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"reflect"
)

//readDataToMemory Reads a given file as json , decodes this json to a UserMap and returns the UserMap. The Files data must be an Encoded UserMap(see saveDataToFile)
func readDataToMemory(path string) (data UserMap, err error) {
	var bytedata []byte
	data = make(UserMap)
	// read the whole file at once
	bytedata, err = ioutil.ReadFile(path)
	if err != nil {
		errors.Wrap(err, "Error in ReadFromFile")
		return
	}

	//Decode json byte array to hashmap TODO=Find better serialisation
	err = json.Unmarshal(bytedata, &data)

	if err != nil {
		errors.Wrap(err, "Error in ReadFromFile while decoding Json byte array to hashmap")
		return
	}

	return
}

//saveDataToFile saves a given UserMap to a file in the given path by encoding the UserMap as Json and writing the Text to the Filesystem
func saveDataToFile(data UserMap, path string) (err error) {
	var bytedata []byte

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		os.Create(path)
	}

	//Encode file to Json byte array TODO=Find better serialisation
	bytedata, err = json.Marshal(data)
	if err != nil {
		errors.Wrap(err, "Error in SaveToFile while engoding map to json")
		return
	}

	// write the whole body at once with unix permission TODO=CHECK Windows permission working ?
	err = ioutil.WriteFile(path, bytedata, 0660)
	if err != nil {
		errors.Wrap(err, "Error in SaveToFile while writing file to disk")
		return
	}
	return
}

//writeUser adds a UserRecord to the Persistent UserStorage. checks UserRecord if all elements are set, returns error if not.
//Then the Package Global UserMap(in userStorage) gets written to file by SaveDataToFile.
func writeUser(record UserRecord, path string) (err error) {
	managersUserStorage.RWMutex.Lock()
	defer managersUserStorage.RWMutex.Unlock()

	//TODO cause error if value is nil and dont add user
	v := reflect.ValueOf(record)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == nil {
			errors.WithMessage(err, "A Field of the userrecord was Nil")
			return
		}
	}

	managersUserStorage.UserMap[record.Email] = record //TODO ADD error handling

	err = saveDataToFile(managersUserStorage.UserMap, path + "/userdatabase")
	if err != nil {
		errors.Wrap(err, "Error in Write User while saving data to Disk")
		//TODO Revert Changes done to Global UserStorage
	}

	return
}

//ReadUser finds a UserRecord in the PackageGlobal UserStorage by Email and returns nil or an valid entry.
func ReadUser(email string) (record UserRecord, present bool, err error) {
	managersUserStorage.RWMutex.RLock()
	defer managersUserStorage.RWMutex.RUnlock()

	record , present = managersUserStorage.UserMap[email] //TODO ADD error handling
	err = nil

	return
}

func getNextUID() (id int16) {
	id = 0
	for _, value := range managersUserStorage.UserMap {
		if(value.UID > id) {
			id = value.UID
		}
	}
	return id+1;
}