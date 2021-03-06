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
	"os"
	"testing"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const TESTFILEPATH = "testdir/userdatabase"
const TESTFOLDER = "testdir"

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	finally()
	os.Exit(retCode)
}

//Testdata which gets populated in setup()
var TestData UserMap

func setup() {

	TestData = UserMap{
		"flo@myprivatemail.de":   {UID: 1, Email: "flo@myprivatemail.de", Name: "Florian Braun", HashedPassword: "???", Salt: "???"},
		"lena.hoinkis@gmail.com": {UID: 2, Email: "lena.hoinkis@gmail.com", Name: "Lena Hoinkis", HashedPassword: "???", Salt: "???"},
	}

	for k, v := range TestData {
		managersUserStorage.UserMap[k] = v
	}

}

//Test_FileSaveRead Tests if File gets Saved correctly and Folders in its path get created automatically
func Test_FileSave(t *testing.T) {

	var LocalTestStorage UserStorage
	var err error

	LocalTestStorage.UserMap = make(UserMap)

	//Copy Map
	for k, v := range TestData {
		LocalTestStorage.UserMap[k] = v
	}

	os.RemoveAll(TESTFILEPATH)

	LocalTestStorage.RLock()
	err = saveDataToFile(LocalTestStorage.UserMap, TESTFILEPATH)
	if err != nil {
		errors.Wrap(err, "Error while saving user file to disk")
		t.Error(err)
	}
	LocalTestStorage.RUnlock()

	if _, err := os.Stat(TESTFILEPATH); os.IsNotExist(err) {
		t.Error("Testfile was not created automatically")
	}

	if _, err := os.Stat(TESTFILEPATH); os.IsNotExist(err) {
		t.Error("Testfolder was not created automatically")
	}

}
//Test_FileRead Tests if read data from filesystem equals the data in memory
func Test_FileRead(t *testing.T) {

	var LocalTestStorage UserStorage
	var err error

	LocalTestStorage.UserMap = make(UserMap)

	LocalTestStorage.Lock()
	LocalTestStorage.UserMap, err = readDataToMemory(TESTFILEPATH)
	LocalTestStorage.Unlock()

	if err != nil {
		errors.Wrap(err, "Error while reading from file")
		t.Error(err)
	}

	assert.Equal(t, TestData, LocalTestStorage.UserMap, "Saved and Read data is not Equal")
}

// Test_SynchronizedGlobalUserStorage Tests Concurrent Access to GlobalUserStorage via "ReadUser" and "WriteUser" -> uses Random acces to trigger a Concurrent Access
func Test_SynchronizedGlobalUserStorage(t *testing.T) {

	for i := 0; i < 1000; i++ {
		go readWriteTest(t)
	}

}

//readWriteTest is a Helper Function to test Concurrency
func readWriteTest(t *testing.T) {
	aUserRecord, _, err := ReadUser("flo@myprivatemail.de")
	if err != nil {
		errors.Wrap(err, "Error while reading a user with function 'ReadUser'")
		t.Error(err)
	}
	err = writeUser(aUserRecord,TESTFOLDER)
	if err != nil {
		errors.Wrap(err, "Error while writing a user with function 'writeUser'")
		t.Error(err)
	}
}

//Test_SyncToFileSystem Tests Writes to filesystem and NIL handling trough uncomplete UserMap
func Test_SyncToFileSystem(t *testing.T) {

	//Clear all Variables
	setup()

	var err error
	var LocalTestStorage UserStorage

	//Test Variable got modifyed
	err = writeUser(UserRecord{UID: 3, Email: "someone@somemail.com", Name: "someone", HashedPassword: "???", Salt: "???"},TESTFOLDER)
	if err != nil {
		errors.Wrap(err, "Error while writing a user with function 'writeUser'")
		t.Error(err)
	}
	assert.NotEqual(t, TestData, managersUserStorage.UserMap, "TestData and Modified GobalUserStorage should be different but they are equal")

	//Test File and Variable are Equal
	LocalTestStorage.Lock()
	LocalTestStorage.UserMap, err = readDataToMemory(TESTFILEPATH)
	LocalTestStorage.Unlock()

	if err != nil {
		errors.Wrap(err, "Error while reading from file")
		t.Error(err)
	}

	managersUserStorage.RLock()
	LocalTestStorage.RLock()
	defer managersUserStorage.RUnlock()
	defer LocalTestStorage.RUnlock()
	assert.Equal(t, managersUserStorage.UserMap["someone@somemail.com"], LocalTestStorage.UserMap["someone@somemail.com"], "Data in File and globalUserStorage are not Equal, Changes where not wrote to permanent storage!")

}

//Aufräumen nach dem Test
func finally() {
	os.Remove(TESTFILEPATH)
	os.Remove(TESTFOLDER)
}
