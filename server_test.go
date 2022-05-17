package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"tiny-url/database"
)

func init() {
	go startServer()
}
func TestGet(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/tzions12")
	t.Log(resp)
	if err != nil {
		t.Error(err)
	}
}

func TestRedirect(t *testing.T) {

}

func TestUnsupportedhttpMethods(t *testing.T) {

}

func TestPostNewEntry(t *testing.T) {
	url := "https://github.com/Tzion?tab=repositories"
	resp, err := http.Post("http://localhost:8080/", "text/plain", bytes.NewBufferString(url))
	if err != nil {
		t.Error(err)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	shortUrl := string(bodyBytes)
	entry := database.FetchLongUrl(shortUrl)
	if entry == nil {
		t.Fatalf("Entry for short-url %q does not exsits in the database", shortUrl)
	}
	if *entry != url {
		t.Errorf("stored url does not match the original url\nhave %s\nwant %s", *entry, url)
	}
}
