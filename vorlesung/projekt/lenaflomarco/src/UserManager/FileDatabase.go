package UserManager

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"reflect"
	"fmt"
)

func ReadDataToMemory(path string) (data UserMap, err error) {
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

	// write the whole body at once with unix permission
	err = ioutil.WriteFile(path, bytedata, 0660)
	if err != nil {
		errors.Wrap(err, "Error in SaveToFile while writing file to disk")
		return
	}
	return
}

func WriteUser(record UserRecord, path string) (err error) {
	managersUserStorage.RWMutex.Lock()
	defer managersUserStorage.RWMutex.Unlock()


	managersUserStorage.UserMap[record.Email] = record //TODO ADD error handling

	err = saveDataToFile(managersUserStorage.UserMap, path)
	if err != nil {
		errors.Wrap(err, "Error in Write User while saving data to Disk")
	}

	return
}

func ReadUser(email string) (record UserRecord, err error) {
	managersUserStorage.RWMutex.RLock()
	defer managersUserStorage.RWMutex.RUnlock()

	record = managersUserStorage.UserMap[email] //TODO ADD error handling
	err = nil

	return
}
