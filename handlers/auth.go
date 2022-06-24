package handlers

import "net/http"

// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "login")
}

// GET /signup
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}