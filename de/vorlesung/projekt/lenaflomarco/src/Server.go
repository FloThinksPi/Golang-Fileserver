/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

package main

import (
	"crypto/tls"
	"net/http"
	"Utils"
	"Flags"
	"path/filepath"
	"strconv"
	"UserManager"
	"SessionManager"
	"regexp"
	"Templates"
	"time"
	"strings"
)

const (
	// Routes
	rootURL = "/"
	docURL = rootURL + "doc/"
	basicAuthURL = rootURL + "download/"
	funcURL = rootURL + "ops/"

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

	//Login,Logout
	requestMultiplexer.HandleFunc(funcURL + "login", authHandler)

	//Settings (Change Password)
	requestMultiplexer.HandleFunc(funcURL + "settings", sessionCheckHandler(Templates.SettingBackendHandler))

	//Index Functions
	//DeleteData
	requestMultiplexer.HandleFunc(funcURL + "delete", sessionCheckHandler(Templates.DeleteDataHandler))

	//DownloadData
	requestMultiplexer.HandleFunc(funcURL + "download", sessionCheckHandler(Templates.DownloadDataHandler))

	//UploadData
	requestMultiplexer.HandleFunc(funcURL + "upload", sessionCheckHandler(Templates.UploadDataDataHandler))

	//NewFolder
	requestMultiplexer.HandleFunc(funcURL + "newFolder", sessionCheckHandler(Templates.NewFolderHandler))

	//Basic Auth
	requestMultiplexer.HandleFunc(basicAuthURL, basicAuthHandler(Templates.DownloadBasicAuthDataHandler))

	// General Handlers for Website + Godoc
	requestMultiplexer.HandleFunc(rootURL, sessionCheckHandler(rootHandler))

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


//basicAuth - Checks submitted user credentials and grants access to handler
func sessionCheckHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		r.ParseForm()
		cookie, err := r.Cookie("Session")
		if err != nil {
			Utils.LogDebug(err.Error())
			cookie := http.Cookie{Name: "Session", Value: "empty", Expires: time.Now().Add(365 * 24 * time.Hour), Path: "/"}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, loginPageURL, 302)
			return
		} else {
			// Public Folder ?
			publicFolderRegex, _ := regexp.Compile("^public")

			if (SessionManager.ValidateSession(cookie.Value) || publicFolderRegex.MatchString(r.URL.EscapedPath()[1:]) || debugging) {
				handler(w, r)
				return
			} else {
				Utils.LogDebug("Access denied for " + r.URL.EscapedPath() + " by Session Check")
				http.Redirect(w, r, loginPageURL, 302)
				return
			}
		}
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var url string

	if r.URL.EscapedPath()[1:] == "" {
		url = mainPageURL
	} else {
		url = r.URL.EscapedPath()
	}

	path, err := filepath.Abs(pivatePath + url)
	Utils.HandlePrint(err)

	if (url[len(url) - 4:] == "html") {
		// Split string at / , switch case files

		switch url {
		case "/index.html":
			Templates.IndexHandler(w, r, path)
			Utils.LogDebug("File Accessed with TemplateEngine:	" + path)
		case "/settings.html":
			Templates.SettingHandler(w, r, path)
			Utils.LogDebug("File Accessed with TemplateEngine:	" + path)
		default:
			http.ServeFile(w, r, path)
			Utils.LogDebug("File Accessed with StaticFileServer:	" + path)
		}

	} else {
		http.ServeFile(w, r, path)
		Utils.LogDebug("File Accessed with StaticFileServer:	" + path)
	}

}

func authHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	intent := r.FormValue("intent")
	if (intent == "login") {
		Utils.LogDebug("Intent=Login")
		email := strings.ToLower(r.FormValue("email"))
		password := r.FormValue("password")
		if (UserManager.VerifyUser(email, password)) {

			session := Utils.RandString(128)
			SessionManager.NewSession(SessionManager.SessionRecord{Email:email, Session: session, SessionLast:time.Now()})
			cookie := http.Cookie{Name: "Session", Value: session, Expires: time.Now().Add(365 * 24 * time.Hour), Path: "/"}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, mainPageURL, 302)

		} else {
			http.Redirect(w, r, loginPageURL + "?status=failed", 302)
		}
	} else if (intent == "register") {
		Utils.LogDebug("Intent=Register")
		name := r.FormValue("name")
		email := strings.ToLower(r.FormValue("email"))
		password := r.FormValue("password")
		password2 := r.FormValue("password2")

		if (password != password2) {
			//Passwords not equal
			Utils.LogDebug("Passwörter stimmen nicht überein. Registrierung fehlgeschlagen.")
			http.Redirect(w, r, loginPageURL + "?status=passwordsNotEqual", 302)
			return
		}

		registerOK := UserManager.RegisterUser(name, email, password)
		if (registerOK) {
			session := Utils.RandString(128)
			SessionManager.NewSession(SessionManager.SessionRecord{Email:email, Session: session, SessionLast:time.Now()})
			cookie := http.Cookie{Name: "Session", Value: session, Expires: time.Now().Add(365 * 24 * time.Hour), Path: "/"}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, mainPageURL, 302)
			Utils.LogDebug("Registrierung erfolgreich.")
		} else {
			http.Redirect(w, r, loginPageURL + "?status=userAlreadyExists", 302)
			Utils.LogDebug("Benutzer mit angegebener E-Mail-Adresse existiert bereits. Registrierung fehlgeschlagen.")
		}
	} else if (r.URL.Query().Get("intent") == "logout") {
		Utils.LogDebug("Intent=Logout")
		cookie, err := r.Cookie("Session")
		Utils.HandlePrint(err)
		session, present := SessionManager.GetSessionRecord(cookie.Value)
		if present {
			err = SessionManager.InvalidateSession(session.Session)
		}
		if (err != nil) {
			http.Redirect(w, r, loginPageURL + "?status=error", 302)
			return
		}
		http.Redirect(w, r, loginPageURL + "?status=logout", 302)
		return
	} else {
		Utils.LogDebug("Intent=BadRequest")
		http.Redirect(w, r, loginPageURL + "?status=badrequest", 302)
		return
	}
}



//basicAuth - Checks submitted user credentials and grants access to handler
func basicAuthHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, pass, ok := r.BasicAuth()

		if ok && UserManager.VerifyUser(email, pass) {
			handler(w, r)
		} else {
			w.Header().Set("WWW-Authenticate", `Basic realm="Fileserver: Bitte mit E-Mail-Adresse und Kennwort anmelden, um auf Dateien zuzugreifen."`)
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		}

	}
}

