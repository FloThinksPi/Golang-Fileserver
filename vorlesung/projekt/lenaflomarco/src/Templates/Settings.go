package Templates

import (
	"github.com/pkg/errors"
	"net/http"
	"Utils"
	"UserManager"
	"SessionManager"
)

type SettingDaten struct {
	UserName string
}

func SettingHandler(w http.ResponseWriter, r *http.Request, path string) {

	var Data SettingDaten

	cookie, err := r.Cookie("Session")
	Utils.HandlePrint(err)
	session, present := SessionManager.GetSessionRecord(cookie.Value)
	if present {
		user, present, err := UserManager.ReadUser(session.Email)
		Utils.HandlePrint(err)
		if present {
			Data.UserName = user.Name
			RenderTemplate(w, path , &Data)
		} else {
			Utils.HandlePanic(errors.New("Inconsistency in Session to User Storage!"))
		}
	} else {
		Utils.HandlePanic(errors.New("Inconsistency in Session Storage !"))
	}
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