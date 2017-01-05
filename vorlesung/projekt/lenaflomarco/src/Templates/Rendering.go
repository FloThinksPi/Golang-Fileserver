package Templates

import (
	"net/http"
	"html/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t,err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,p)
}

