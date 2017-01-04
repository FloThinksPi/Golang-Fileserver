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
	"regexp"
)

const (
	// Routes
	rootURL = "/"
	docURL = rootURL + "doc/"
	funcURL = rootURL + "ops/"
	loginURL = rootURL + funcURL + "login"
	logoutURL = rootURL + funcURL + "login"


	// Paths
	pivatePath = "res/html/"
	publicPath = pivatePath + "/public/"


	// URLs
	mainPageURL = rootURL + "index.html"
	loginPageURL = rootURL + "public/login.html"

	// MISC
	debugging = false; // Disables Login for Debugging
)

func main() {
	requestMultiplexer := http.NewServeMux()

	// General Handlers for Website + Godoc
	requestMultiplexer.HandleFunc(docURL, docHandler)
	requestMultiplexer.HandleFunc(rootURL, sessionCheckHandler)

	//Login,Logout
	requestMultiplexer.HandleFunc(loginURL, authHandler)
	requestMultiplexer.HandleFunc(logoutURL, authHandler)

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

	// Public Folder ?
	publicFolderRegex, _ := regexp.Compile("^public")

	if (SessionManager.ValidateSession(session) || debugging || publicFolderRegex.MatchString(r.URL.Path[1:])) {
		rootHandler(w, r)
		return
	} else {
		Utils.LogDebug("Access denied for " + r.URL.Path + " by Session Check")
		http.Redirect(w, r, loginPageURL, 302)
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var url string

	if r.URL.Path[1:] == "" {
		url = mainPageURL
	} else {
		url = r.URL.Path[1:]
	}

	path, err := filepath.Abs(pivatePath + url)
	Utils.HandlePrint(err)

	if (url[len(url) - 4:] == "html") {
		t, err := template.ParseFiles(path)
		Utils.HandlePrint(err)
		files, err := filepath.Glob("*")
		Utils.HandlePrint(err)

		Utils.LogDebug("File Accessed with TemplateEngine:	" + path)
		t.Execute(w, files)
	} else {
		http.ServeFile(w, r, path)
		Utils.LogDebug("File Accessed with StaticFileServer:	" + path)
	}

}

func authHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	intent := r.FormValue("intent")
	if (intent == "login") {
		email := r.FormValue("email")
		password := r.FormValue("password")
		if (UserManager.VerifyUser(email, password)) {
			http.Redirect(w, r, loginPageURL, 302)
		} else {
			http.Redirect(w, r, loginPageURL + "?status=failed", 302)
		}
	} else if (intent == "logout") {
		session := r.FormValue("session")
		err := SessionManager.InvalidateSession(session)
		if (err != nil) {
			http.Redirect(w, r, loginPageURL + "?status=error", 302)
			return
		}
		http.Redirect(w, r, loginPageURL + "?status=logout", 302)
		return
	} else {
		http.Redirect(w, r, loginPageURL + "?status=badrequest", 302)
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

