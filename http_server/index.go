package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	const USERNAME = "username"
	const PASSWORD = "password"
	const PORT = ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "home, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "login, %s, %s", r.URL.Query().Get(USERNAME), r.URL.Query().Get(PASSWORD))
	})
	log.Fatal(http.ListenAndServe(PORT, nil))
}
