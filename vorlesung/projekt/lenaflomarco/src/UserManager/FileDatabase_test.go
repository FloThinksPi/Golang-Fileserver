package UserManager

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/pkg/errors"
)

const TESTFILEPATH = "testdir/tmp.db"
const TESTFOLDER = "testdir"

var TestData = []UserRecord{
	{UID:1, Email:"flo@myprivatemail.de", Name:"Florian Braun", HashedPassword:"???", Salt:"???"},
	{UID:1, Email:"lena.hoinkis@gmail.com", Name:"Lena Hoinkis", HashedPassword:"???", Salt:"???"},
}

type UserRecord struct {
	UID            int64  //Unique ID
	Email          string //Email
	Name           string //Name
	HashedPassword string //Password Hashed and Salted
	Salt           string //Salt
}

func test_FileSave(t *testing.T) {

	response, err := saveToFile(TestData, TESTFILEPATH)
	if err != nil {
		t.Error(response)
	}

	if !exists(TESTFOLDER) || !exists(TESTFILEPATH) {
		t.Error("Testfolder or Testfile was not created automatically")
	}

}

func test_FileRead(t *testing.T) {

	data, err := readFromFile(TESTFILEPATH)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t,TestData,data,"Saved and Read data is not Equal")

}

// test_finaly = Aufräumen nach dem Test
func test_finaly(t *testing.T){
	cleanup(TESTFOLDER)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, errors.Wrap(err,"Some error while checking if file exists")
}

func cleanup(path string) error {
	err := os.Remove(path)
	return err
}