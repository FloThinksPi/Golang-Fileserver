package Templates

import (
	"net/http"
	"Utils"
	"SessionManager"
	"UserManager"
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

func SettingBackendHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	passwordOld := r.FormValue("passwordOld")
	passwordNew := r.FormValue("passwordNew")
	passwordNew2 := r.FormValue("passwordNew2")

	if (passwordNew != passwordNew2) {
		//Passwords not equal
		Utils.LogDebug("Neue Passwörter stimmen nicht überein. Änderung fehlgeschlagen.")
		http.Redirect(w, r, "settings.html?status=passwordsNotEqual", 302)
		return
	}

	cookie, err := r.Cookie("Session")
	Utils.HandlePrint(err)
	session, present := SessionManager.GetSessionRecord(cookie.Value)
	if(present) {
		email := session.Email

		Utils.LogDebug("EMail: " + email)

		if (UserManager.VerifyUser(email, passwordOld)) {
			UserManager.ChangePassword(email, passwordNew)
			http.Redirect(w, r, "settings.html", 302)
			Utils.LogDebug("Kennwort erfolgreich geändert.")
		} else {
			http.Redirect(w, r, "settings.html?status=oldPasswordNotValid", 302)
			Utils.LogDebug("Altes Kennwort nicht korrekt. Kennwortänderung fehlgeschlagen.")
		}
	}

}