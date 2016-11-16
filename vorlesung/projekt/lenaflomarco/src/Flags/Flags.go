package Flags

import (
	"flag"
	"path/filepath"
)

var (
	Verbosity = flag.Int("v", 3, "Verbosity (0=Errors ,3=All)")
	WorkDir = flag.String("workdir", standardizePath("datastorage"), "Folder in which the Server stores the Userdatabase aswell as the Users uploaded data")
)

func init() {
	flag.Parse()
}


//TODO Cant use Utils/Logging.go here because of cycling includes :/
func standardizePath(path string) string {

	newPath, err := filepath.Abs(path)
	if err != nil {
		println(err)
	}

	return newPath
}