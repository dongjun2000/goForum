package handlers

import (
	"fmt"
	"goForum/models"
	"net/http"
)

// POST /thread/post 再指定群组中创建新主题
func PostThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		user, err := sess.User()
		if err != nil {
			danger(err, "Cannot get user from session")
		}
		body := r.PostFormValue("body")
		uuid := r.PostFormValue("uuid")
		thread, err := models.ThreadByUUID(uuid)
		if err != nil {
			error_message(w, r, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			danger(err, "Cannot create post")
		}

		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(w, r, url, 302)
	}
}
