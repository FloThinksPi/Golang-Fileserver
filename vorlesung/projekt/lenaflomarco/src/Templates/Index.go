
	var (
		status int
		err error
	)
	defer func() {
		if nil != err {
			http.Error(w, err.Error(), status)
		}
	}

}

func getAbsUserPath(r *http.Request) string {
	cookie, err := r.Cookie("Session")
	Utils.HandlePrint(err)
	session, present := SessionManager.GetSessionRecord(cookie.Value)
	if present {
		user, present, err := UserManager.ReadUser(session.Email)
		Utils.HandlePrint(err)
		if present {
			return Flags.GetWorkDir() + "/" + strconv.Itoa(int(user.UID)) + "/"
		} else {
			Utils.HandlePanic(errors.New("Inconsistency in Session to User Storage!"))
		}
	} else {
		Utils.HandlePanic(errors.New("Inconsistency in Session Storage !"))
	}
	return ""
}