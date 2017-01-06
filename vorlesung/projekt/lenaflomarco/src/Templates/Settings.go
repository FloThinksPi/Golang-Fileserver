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

