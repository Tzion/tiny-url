package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"tiny-url/database"
	"tiny-url/shortener"
)

func main() {
	startServer()
}

func startServer() {
	fmt.Printf("Starting tiny-url server\n")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	default:
		http.Error(w, "only GET and POST methods are supported.", http.StatusNotFound)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET request for %q\n", html.EscapeString(r.URL.Path))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err == nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	url := string(body)
	log.Printf("POST request for %q\n", url)
	tinyUrl := shortener.ShortenUrl(url)
	database.Insert(tinyUrl, url)
	w.Write([]byte(tinyUrl))
}
