package Flags

import (
	"flag"
	"path/filepath"
)


//noinspection LongLine
var (
	verbosity = flag.Int("v", 3, "Verbosity (0=Errors ,3=All)")
	workDir = flag.String("workdir", standardizePath("datastorage"), "Folder in which the Server stores the Userdatabase aswell as the Users uploaded data")
	port = flag.Int("port", 8080, "Port for the Webserver to listen to (Port 1-1024 need SuperUser Privileges on Unix systems!)")
	tlsCert = flag.String("tlscert", standardizePath("res/certificates/server.pem"), "Path for the TLS Certificate")
	tlsKey = flag.String("tlskey", standardizePath("res/certificates/server.key"), "Path for the TLS Key")
)

//Getters , Dereferencing to keep privacy. Input Ckack and Handling Here

//GetVerbosity returns verbosity
//	verbosity = flag.Int("v", 3, "Verbosity (0=Errors ,3=All)")
func GetVerbosity() int {
	return *verbosity
}

//GetWorkDir returns workDir
//	workDir = flag.String("workdir", standardizePath("datastorage"), "Folder in which the Server stores the Userdatabase aswell as the Users uploaded data")
func GetWorkDir() string {
	return *workDir
}

//GetPort returns port
//	port = flag.Int("port",8080,"Port for the Webserver to listen to (Port 1-1024 need SuperUser Privileges on Unix systems!)")
func GetPort() int {
	return *port
}

//GetTLScert returns tlsCert
//	tlsCert = flag.String("tlscert", standardizePath("res/certificates/server.pem"), "Path for the TLS Certificate")
func GetTLScert() string {
	return *tlsCert
}

//GetTLSkey returns tlsKey
//	tlsKey = flag.String("tlskey", standardizePath("res/certificates/server.key"), "Path for the TLS Key")
func GetTLSkey() string {
	return *tlsKey
}


// Init gets called before main()
func init() {
	flag.Parse()
}

//standardizePath = wrapper for filepath.Abs and handles its error
func standardizePath(path string) string {

	newPath, err := filepath.Abs(path)
	if err != nil {
		println(err)//TODO Cant use Utils/Logging.go here because of cycling includes :/
	}

	return newPath
}
