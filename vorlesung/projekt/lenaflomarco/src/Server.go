

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"UserManager"
	"net/http"
	"time"
)

func serveMain(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	StoredAs := r.Form.Get("index.html") // file name
	data, err := ioutil.ReadFile("res/html/index.html" + StoredAs)
	if err != nil {
		fmt.Fprint(w, err)
	}
	http.ServeContent(w, r, StoredAs, time.Now(), bytes.NewReader(data))
}

func main() {
	http.HandleFunc("/", serveMain)          // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	UserManager.VerifyHash("123", "123", "123")
}
