package main

import (
	"net/http"

	"github.com/sleong110/go-web-programming/ch2_chit_chat_app/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nl {
		danger(err, "Cannot find user")
	}

	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
