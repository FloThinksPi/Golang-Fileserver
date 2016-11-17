package main

import (
	"crypto/tls"
	"net/http"
	"Utils"
	"Flags"
	"path/filepath"
	"strconv"
	"encoding/json"
	"os"
	"log"
	"html/template"
)

const (
	rootURL = "/"
	apiURL = "/API/"
	docURL = "/doc/"
)

func main() {
	requestMultiplexer := http.NewServeMux()

	requestMultiplexer.HandleFunc(docURL, doc)
	requestMultiplexer.HandleFunc(rootURL, root)

	//REST Test
	requestMultiplexer.HandleFunc(apiURL, apiDoc)
	requestMultiplexer.HandleFunc(apiURL + "getFileList", apiGetFileList)

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		},
	}

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(Flags.GetPort()),
		Handler:      requestMultiplexer,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	Utils.HandlePanic(srv.ListenAndServeTLS(Flags.GetTLScert(), Flags.GetTLSkey()))
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	path, err := filepath.Abs("res/html/" + r.URL.Path[1:])
	Utils.HandlePrint(err)

	Utils.LogDebug("File Accessed:	" + path)
	http.ServeFile(w, r, path)

	t := template.Must(template.New("").Parse(`<table>{{range .}}<tr><td>{{.}}</td></tr>{{end}}</table>`))
	names := []string{"john", "jim"}
	if err := t.Execute(os.Stdout, names); err != nil {
		log.Fatal(err)
	}


}

func doc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	path, err := filepath.Abs("res/" + r.URL.Path[1:])
	Utils.HandlePrint(err)

	Utils.LogDebug("File Accessed:	" + path)

	if (r.URL.Path[1:] == "doc/" || r.URL.Path[1:] == "doc") {
		Utils.LogDebug("Redirecting from doc to doc/pkg/fileServer.html")
		http.Redirect(w, r, "pkg/FileServer.html", 300)
	} else {
		http.ServeFile(w, r, path)
	}

}

//apiDoc Shows Documentation of the Rest Api
func apiDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	json.NewEncoder(w).Encode("Documentation Of Api May Follow")
}

func apiGetFileList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	files, _ := filepath.Glob("*")

	json.NewEncoder(w).Encode(files)

}


