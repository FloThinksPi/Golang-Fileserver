package main

import (
	"crypto/tls"
	"net/http"
	"Utils"
	"Flags"
	"path/filepath"
	"strconv"
)

func init() {

}

func main() {
	requestMultiplexer := http.NewServeMux()

	requestMultiplexer.HandleFunc("/", root)
	requestMultiplexer.HandleFunc("/doc/", func(w http.ResponseWriter, r *http.Request) {

		path, err := filepath.Abs("res/" + r.URL.Path[1:])
		Utils.HandlePrint(err)

		Utils.LogDebug("File Accessed:	" + path)

		if (r.URL.Path[1:] == "doc/" || r.URL.Path[1:] == "doc") {
			Utils.LogDebug("Redirecting from doc to doc/pkg/fileServer.html")
			http.Redirect(w, r, "pkg/FileServer.html", 300)
		} else {
			http.ServeFile(w, r, path)
		}

	})

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

func root(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	w.Write([]byte("Test Site \n"))
}
