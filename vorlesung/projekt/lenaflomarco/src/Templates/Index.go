package Templates

import (
	"net/http"
)

type IndexData struct {
	Name string
	Size string
	Date string
}

type IndexDaten []IndexData

func IndexHandler(w http.ResponseWriter, r *http.Request, path string) {


	Data := IndexDaten{
		{Name:"TEst",Size:"10kb",Date:"1.1.2017"},
		{Name:"Zweite Datei",Size:"1Gb",Date:"10.3.1990"},
	}

	RenderTemplate(w, path , &Data)

}

