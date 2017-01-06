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
	"html/template"
)

//RenderTemplate Abstraktion zum Verarbeiten einer Templatedatei (HTML) und Übergeben der entsprechenden Werte
func RenderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t,err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w,p)
}

