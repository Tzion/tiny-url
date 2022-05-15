package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
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
	fmt.Printf("GET request for %q\n",html.EscapeString(r.URL.Path))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	return

}
