/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

package Templates

import (
	"net/http"
	"Utils"
	"io"
	"strconv"
	"os"
	"UserManager"
	"Flags"
	"SessionManager"
	"github.com/pkg/errors"
	"log"
	"path/filepath"
	"mime/multipart"
	"io/ioutil"
	"strings"
)

type IndexData struct {
	Name       string
	FolderPath string
	ObjectPath string
	Size       string
	Date       string
	Image      string
	IsFolder   bool
}

type IndexDaten struct {
	FileData   []IndexData
	FolderPath string
	UserName   string
}

//IndexHandler stellt die Daten für das Template zur Verfügung
func IndexHandler(w http.ResponseWriter, r *http.Request, path string) {

	var Data IndexDaten

	userPath := getAbsUserPath(r)
	wantedPath := ""
	wantedPath = r.URL.Query().Get("path")
	if invalidPath(wantedPath) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	fullPath := filepath.Join(userPath, wantedPath)
	Utils.LogDebug("Showing Files of Folder:	" + fullPath)
	files, _ := ioutil.ReadDir(fullPath)

	cookie, err := r.Cookie("Session")
	Utils.HandlePrint(err)
	session, present := SessionManager.GetSessionRecord(cookie.Value)
	if present {
		user, present, err := UserManager.ReadUser(session.Email)
		Utils.HandlePrint(err)
		if present {

			// All Informations given here!
			Data.FolderPath = wantedPath
			Data.UserName = user.Name

			if(wantedPath != "") {
				pos := strings.LastIndex(wantedPath, "/")
				subfolder := wantedPath[0:pos]

				Data.FileData = append(Data.FileData, IndexData{Name:"../", FolderPath: wantedPath, Size:"", Date:"", Image:"folder", ObjectPath: subfolder, IsFolder:true})
			}

			for _, f := range files {
				if f.IsDir() {
					Data.FileData = append(Data.FileData, IndexData{Name:f.Name(), FolderPath: wantedPath, Size:"", Date:f.ModTime().Format("02.01.2006 15:04:05"), Image:"folder", ObjectPath: wantedPath + "/" + f.Name(), IsFolder:true})
				} else {
					Data.FileData = append(Data.FileData, IndexData{Name:f.Name(), FolderPath: wantedPath, Size:getSize(f.Size()), Date:f.ModTime().Format("02.01.2006 15:04:05"), Image:"file", ObjectPath: wantedPath + "/" + f.Name(), IsFolder:false})
				}
			}

			RenderTemplate(w, path, &Data)

		} else {
			Utils.HandlePanic(errors.New("Inconsistency in Session to User Storage!"))
		}
	} else {
		Utils.HandlePanic(errors.New("Inconsistency in Session Storage !"))
	}

}

//DeleteDataHandler löscht Dateien und Ordner
func DeleteDataHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Query().Get("filepath")
	if invalidPath(path) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	fullPath := filepath.Join(getAbsUserPath(r), path)
	Utils.LogDebug("File Deleted by DeleteDataHandler:	" + fullPath)
	os.RemoveAll(fullPath)
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
}

//DownloadDataHandler bietet die angeforderte Datei zum Download an
func DownloadDataHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("filepath")
	if invalidPath(path) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	fullPath := filepath.Join(getAbsUserPath(r), path)
	Utils.LogDebug("File Accessed by DownloadDataHandler:	" + fullPath)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-disposition", `attachment; filename="` + filepath.Base(path) + `"`)

	http.ServeFile(w, r, fullPath)
}

//DownloadBasicAuthDataHandler bietet die per BasicAuth angeforderte Datei zum Download an
//Der Aufruf erfolgt über /download/?filepath=
func DownloadBasicAuthDataHandler(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()
	usr, present, err := UserManager.ReadUser(email)
	Utils.HandlePrint(err)
	if present {
		path := r.URL.Query().Get("filepath")
		if invalidPath(path) {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
			return
		}

		fullPath := filepath.Join(Flags.GetWorkDir(), strconv.Itoa(int(usr.UID)), path)
		Utils.LogDebug("File Accessed by DownloadBasicAuthDataHandler:	" + fullPath)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-disposition", `attachment; filename="` + filepath.Base(path) + `"`)
		http.ServeFile(w, r, fullPath)
	} else {
		Utils.LogWarning("Inconsistency in User Storage!")
	}
}

//NewFolderHandler erzeugt neue Ordner im Userverzeichnis
func NewFolderHandler(w http.ResponseWriter, r *http.Request) {
	//Todo tested yet
	//read folderName
	r.ParseForm()
	folderName := r.FormValue("folderName")
	targetPath := r.FormValue("folderPath")
	if invalidPath(targetPath) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}
	if invalidPath(folderName) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	//join Foldername to UserPath
	path := filepath.Join(getAbsUserPath(r), targetPath, folderName)

	//Try make Dir
	Utils.LogDebug("Making New Directory with NewFolderHandler:	" + path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		Utils.LogError("Error creating directory")
		log.Println(err)
	}
	http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
	return
}

//UploadDataDataHandler nimmt die hochzuladende Datei entgegen und legt diese im richtigen Userverzeichnis ab
func UploadDataDataHandler(w http.ResponseWriter, r *http.Request) {

	targetPath := r.URL.Query().Get("folderPath")
	if invalidPath(targetPath) {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)
		return
	}

	var (
		status int
		err error
	)
	defer func() {
		if nil != err {
			http.Error(w, err.Error(), status)
		}
	}()
	// parse request
	const _24K = (1 << 20) * 24
	if err = r.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		return
	}
	for _, fheaders := range r.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				return
			}
			// open destination
			var outfile *os.File
			Utils.LogDebug("Uploading File to: 	" + filepath.Join(getAbsUserPath(r), targetPath, hdr.Filename))
			if outfile, err = os.Create(filepath.Join(getAbsUserPath(r), targetPath, hdr.Filename)); nil != err {
				status = http.StatusInternalServerError
				return
			}
			// 32K buffer copy

			if _, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}

			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusTemporaryRedirect)

		}
	}

}

//getAbsUserPath gibt den Pfad zum Userverzeichnis zurück, in welchem sich die hochgeladenen Dateien befinden
func getAbsUserPath(r *http.Request) string {
	cookie, err := r.Cookie("Session")
	Utils.HandlePrint(err)
	session, present := SessionManager.GetSessionRecord(cookie.Value)
	if present {
		user, present, err := UserManager.ReadUser(session.Email)
		Utils.HandlePrint(err)
		if present {
			return filepath.Join(Flags.GetWorkDir(), "/", strconv.Itoa(int(user.UID)), "/")
		} else {
			Utils.LogWarning("Inconsistency in Session to User Storage!")
		}
	} else {
		Utils.LogWarning("Inconsistency in Session Storage !")
	}
	return ""
}

//invalidPath überprüft ob im angegebenen Pfad Path Traversal ausgenutzt wird
func invalidPath(p string) bool {
	if strings.Contains(p, "./") {
		Utils.LogWarning("Path Traversal detected in:	" + p)
		return true
	}
	return false
}

//getSize gibt die Größe abhängig von der Größenordnung in einer passenden Einheit an
func getSize(size int64) string {
	if(size < 1000) {
		return strconv.Itoa(int(size)) + " Byte"
	} else if(size < 1000*1000) {
		i := int(size)/1000
		return strconv.Itoa(i) + " KB"
	} else if(size < 1000*1000*1000) {
		i := int(size)/(1000*1000)
		return strconv.Itoa(i) + " MB"
	} else {
		i := int(size)/(1000*1000*1000)
		return strconv.Itoa(i) + " GB"
	}
}