package handlers

import (
	"goForum/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		generateHTML(w, threads, "layout", "navbar", "index")
	}
}