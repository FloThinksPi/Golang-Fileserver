package Templates

import (
	"net/http"
)

type SettingData struct {
	Name string
	Size string
	Date string
}

type SettingDaten []SettingData

func SettingHandler(w http.ResponseWriter, r *http.Request, path string) {

	Data := SettingDaten{
		{Name:"TEst",Size:"10kb",Date:"1.1.2017"},
		{Name:"Zweite Datei",Size:"1Gb",Date:"10.3.1990"},
	}

	RenderTemplate(w, path , &Data)

}

