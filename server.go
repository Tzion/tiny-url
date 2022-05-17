package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"strings"
	"tiny-url/database"
	"tiny-url/shortener"
)

var Host string = "localhost"
var Port string = "8080"

func main() {
	startServer()
}

func startServer() {
	address := Host + ":" + Port
	fmt.Printf("Starting tiny-url server on address: %s\n", address)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GET request for %q\n", r.URL.Path)
	tinyUrl := strings.TrimLeft(r.URL.Path, "/")
	url := database.FetchLongUrl(tinyUrl)
	if url == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	} else {
		log.Printf("Redirecting client to %q\n", *url)
		http.Redirect(w, r, *url, http.StatusFound)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	url := string(body)
	log.Printf("POST request for %q\n", url)
	tinyUrl, _ := store(url)
	w.Write([]byte(tinyUrl))
}

func store(url string) (string, string) {
	escapedUrl := html.EscapeString(url)
	shortUrl := shortener.ShortenUrl(escapedUrl)
	database.Insert(shortUrl, escapedUrl)
	return shortUrl, url
}
