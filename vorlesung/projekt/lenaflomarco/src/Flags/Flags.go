//WARNING , this Package should not use other Packages to prevent cycling includes. And it should not need to include others

package Flags

import (
	"flag"
	"path/filepath"
)

var (
	verbosity = flag.Int("v", 3, "Verbosity (0=Errors ,3=All)")
	workDir = flag.String("workdir", standardizePath("datastorage"), "Folder in which the Server stores the Userdatabase aswell as the Users uploaded data")
	port = flag.Int("port",8080,"Port for the Webserver to listen to (Port 1-1024 need SuperUser Privileges on Unix systems!)")

)

//Getters , Dereferencing to keep privacy

func GetVerbosity() int {
	return *verbosity
}

func GetWorkDir() string {
	return *workDir
}

func GetPort() int {
	return *port
}

// Init gets called before main()

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
