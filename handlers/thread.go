package handlers

import (
	"goForum/models"
	"net/http"
)

// GET /threads/new  创建群组页面
func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "auth.navbar", "new.thread")
	}
}

// POST /thread/create 创建群组逻辑
func CreateThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		topic := r.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			danger(err, "Cannot create thread")
		}
		http.Redirect(w, r, "/", 302)
	}
}

// GET /thread/read  通过ID渲染指定群组页面
func ReadThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		error_message(w, r, "Cannot read thread")
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(w, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}


