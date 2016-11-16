package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"fmt"
	"time"
	"bytes"
	"log"
	"UserManager"
	"net/http"
)


// note, that variables are pointers
var strFlag = flag.String("long-string", "", "Description")
var boolFlag = flag.Bool("bool", false, "Description of flag")

func init() {
	// example with short version for long flag
	flag.StringVar(strFlag, "s", "", "Description")
}

func main() {
	flag.Parse()
	println(*strFlag, *boolFlag)

	http.HandleFunc("/", serveMain)          // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	UserManager.VerifyHash("123", "123", "123")
}

func serveMain(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	StoredAs := r.Form.Get("index.html") // file name
	data, err := ioutil.ReadFile(filepath.FromSlash("res/html/index.html") + StoredAs)
	if err != nil {
		fmt.Fprint(w, err)
	}
	http.ServeContent(w, r, StoredAs, time.Now(), bytes.NewReader(data))
}


