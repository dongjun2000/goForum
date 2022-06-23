package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GoWeb~")
}

func main() {

	var router = mux.NewRouter()

	router.HandleFunc("/", hello)

	http.ListenAndServe(":3000", router)
}
