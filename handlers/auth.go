package handlers

import (
	"fmt"
	"goForum/models"
	"log"
	"net/http"
)

// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "login")
}

// GET /signup
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
func SignupAccount(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		Name: r.PostFormValue("name"),
		Email: r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		log.Println("Cannot create user")
	}
	http.Redirect(w, r, "/login", 302)
}

// POST /authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		fmt.Println("Cannot find user")
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			fmt.Println("Cannot create session")
		}
		// 给客户端种cookie
		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// GET /logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}