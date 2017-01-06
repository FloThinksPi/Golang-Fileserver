package Templates

import (
	"net/http"
	"Utils"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"html/template"
	"UserManager"
	"Flags"
)

type IndexData struct {
	Name string
	Size string
	Date string
	Image string
	AbsPath string
}

type IndexDaten []IndexData

func IndexHandler(w http.ResponseWriter, r *http.Request, path string) {


	Data := IndexDaten{
		{Name:"TEst",Size:"10kb",Date:"1.1.2017",Image:"file",AbsPath:""},
		{Name:"Zweite Datei",Size:"1Gb",Date:"10.3.1990",Image:"folder",AbsPath:""},
	}

	RenderTemplate(w, path , &Data)
}

func DeleteDataHandler(w http.ResponseWriter, r *http.Request) {
 	Utils.LogDebug("DeleteData Not Implemented")
}

func DownloadDataHandler(w http.ResponseWriter, r *http.Request) {
	Utils.LogDebug("DownloadData Not Implemented")
}

func DownloadBasicAuthDataHandler(w http.ResponseWriter, r *http.Request)  {
	email, _, _ := r.BasicAuth()
	usr, _,_ := UserManager.ReadUser(email)
	path := r.URL.Query().Get("filepath")

	fullPath := Flags.GetWorkDir() + "/" + strconv.Itoa(int(usr.UID)) + "/" + path
	Utils.LogDebug("File Accessed by DownloadBasicAuthDataHandler:	" + fullPath)
	http.ServeFile(w, r, fullPath)
}

func NewFolderHandler(w http.ResponseWriter, r *http.Request) {
	Utils.LogDebug("NewFolderHandler Not Implemented")
}

func UploadDataDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			Utils.LogError(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//TODO: Path ./test durch Userpath ersetzen, anlegen, falls nicht vorhanden...
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			Utils.LogError(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}



