package Templates

import (
	"net/http"
	"Utils"
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

func NewFolderHandler(w http.ResponseWriter, r *http.Request) {
	Utils.LogDebug("NewFolderHandler Not Implemented")
}

func UploadDataDataHandler(w http.ResponseWriter, r *http.Request) {
	Utils.LogDebug("UploadDataDataHandler Not Implemented")
}



