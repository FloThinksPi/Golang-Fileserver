package main

import (
	"crypto/tls"
	"net/http"
	"Utils"
	"Flags"
	"path/filepath"
	"strconv"
	"html/template"
	"UserManager"
	"SessionManager"
)

const (
	rootURL = "/"
	docURL = "/doc/"

	debugging = true; // Disables Login for Debugging
)

func main() {
	requestMultiplexer := http.NewServeMux()

	// General Handlers for Website + Godoc
	requestMultiplexer.HandleFunc(docURL, docHandler)
	requestMultiplexer.HandleFunc(rootURL, sessionCheckHandler)

	//Login,Logout
	requestMultiplexer.HandleFunc(rootURL + "login", authHandler)
	requestMultiplexer.HandleFunc(rootURL + "logout", authHandler)

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

func sessionCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

	r.ParseForm()
	session := r.FormValue("session")

	if (SessionManager.ValidateSession(session) || debugging || r.URL.Path[1:] == "login.html") {
		rootHandler(w, r)
		return
	} else {
		http.Redirect(w, r, "/login.html", 302)
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var url string
	if r.URL.Path[1:] == "" {
		url = "index.html"
	} else {
		url = r.URL.Path[1:]
	}

	path, err := filepath.Abs("res/html/" + url)
	Utils.HandlePrint(err)

	t, err := template.ParseFiles(path)
	Utils.HandlePrint(err)
	files, err := filepath.Glob("*")
	Utils.HandlePrint(err)

	Utils.LogDebug("File Accessed:	" + path)
	t.Execute(w, files)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	intent := r.FormValue("intent")
	if (intent == "login") {
		email := r.FormValue("email")
		password := r.FormValue("password")
		if (UserManager.VerifyUser(email, password)) {
			http.Redirect(w, r, "/index.html", 302)
		} else {
			http.Redirect(w, r, "/login.html?status=failed", 302)
		}
	} else if (intent == "logout") {
		session := r.FormValue("session")
		err := SessionManager.InvalidateSession(session)
		if (err != nil) {
			http.Redirect(w, r, "/login.html?status=error", 302)
			return
		}
		http.Redirect(w, r, "/login.html?status=logout", 302)
		return
	} else {
		http.Redirect(w, r, "/login.html?status=badrequest", 302)
		return
	}
}

func docHandler(w http.ResponseWriter, r *http.Request) {
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

