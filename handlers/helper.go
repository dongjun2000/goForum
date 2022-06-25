package handlers

import (
	"errors"
	"fmt"
	"goForum/models"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	// files := []string{"views/layout.html", "views/navbar.html", "views/index.html"}
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}