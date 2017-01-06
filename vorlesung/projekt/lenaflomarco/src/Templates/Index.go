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
	"time"
)

type IndexData struct {
	Name       string
	FolderPath string
	ObjectPath string
	Size       int64
	Date       time.Time
	Image      string
	IsFolder   bool
}

type IndexDaten []IndexData

func IndexHandler(w http.ResponseWriter, r *http.Request, path string) {

	var Data IndexDaten

	userPath := getAbsUserPath(r)
	wantedPath := ""
	wantedPath = r.URL.Query().Get("path")

	fullPath := userPath + wantedPath

	files, _ := ioutil.ReadDir(fullPath)

	for _, f := range files {
		if f.IsDir() {
			Data = append(Data, IndexData{Name:f.Name(), FolderPath: wantedPath, Size:f.Size(), Date:f.ModTime(), Image:"folder", ObjectPath: wantedPath +"/"+ f.Name(),IsFolder:true})
		} else {
			Data = append(Data, IndexData{Name:f.Name(), FolderPath: wantedPath, Size:f.Size(), Date:f.ModTime(), Image:"file", ObjectPath: wantedPath +"/"+ f.Name(),IsFolder:false})
		}
	}

	RenderTemplate(w, path, &Data)
}

func DeleteDataHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Query().Get("filepath")
	fullPath := getAbsUserPath(r) + path
	Utils.LogDebug("File Deleted by DeleteDataHandler:	" + fullPath)
	os.Remove(fullPath)
	http.StatusText(http.StatusNoContent)

}

func DownloadDataHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("filepath")
	fullPath := getAbsUserPath(r) + path
	Utils.LogDebug("File Accessed by DownloadDataHandler:	" + fullPath)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-disposition", `attachment; filename="`+filepath.Base(path)+`"`)
	http.ServeFile(w, r, fullPath)

}

func DownloadBasicAuthDataHandler(w http.ResponseWriter, r *http.Request) {
	email, _, _ := r.BasicAuth()
	usr, present, err := UserManager.ReadUser(email)
	Utils.HandlePrint(err)
	if present {
		path := r.URL.Query().Get("filepath")

		fullPath := Flags.GetWorkDir() + "/" + strconv.Itoa(int(usr.UID)) + "/" + path
		Utils.LogDebug("File Accessed by DownloadBasicAuthDataHandler:	" + fullPath)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-disposition", `attachment; filename="`+filepath.Base(path)+`"`)
		http.ServeFile(w, r, fullPath)
	} else {
		Utils.HandlePanic(errors.New("Inconsistency in User Storage!"))
	}
}

func NewFolderHandler(w http.ResponseWriter, r *http.Request) {
	//Todo tested yet
	//read folderName
	r.ParseForm()
	folderName := r.FormValue("folderName")
	targetPath := r.FormValue("filepath")

	//join Foldername to UserPath
	path := filepath.Join(getAbsUserPath(r)+targetPath, folderName)

	//Try make Dir
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		Utils.LogError("Error creating directory")
		log.Println(err)
		return
	}
}

func UploadDataDataHandler(w http.ResponseWriter, r *http.Request) {

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
			os.Chdir("/")
			Utils.LogDebug("Uploading File to: 	" + getAbsUserPath(r) + hdr.Filename)
			if outfile, err = os.Create(getAbsUserPath(r) + hdr.Filename); nil != err {
				status = http.StatusInternalServerError
				return
			}
			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}
			status = http.StatusCreated
			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}

}

func getAbsUserPath(r *http.Request) string {
	cookie, err := r.Cookie("Session")
	Utils.HandlePrint(err)
	session, present := SessionManager.GetSessionRecord(cookie.Value)
	if present {
		user, present, err := UserManager.ReadUser(session.Email)
		Utils.HandlePrint(err)
		if present {
			return Flags.GetWorkDir() + "/" + strconv.Itoa(int(user.UID)) + "/"
		} else {
			Utils.HandlePanic(errors.New("Inconsistency in Session to User Storage!"))
		}
	} else {
		Utils.HandlePanic(errors.New("Inconsistency in Session Storage !"))
	}
	return ""
}