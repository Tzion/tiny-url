package main

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"tiny-url/database"
)

func init() {
	go startServer() // TODO potential bug becuase of race between the server thread and the test thread
}

func TestRedirectToStoredUrl(t *testing.T) {
	tinyUrl, url := store("https://recolabs.dev/")
	resp, err := http.Get("http://localhost:8080/" + tinyUrl)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp.StatusCode)
	t.Log(resp)
	t.Log(url)
}

// TODO break down this flow to multiple test cases
func TestPostNewEntry(t *testing.T) {
	url := "https://github.com/Tzion?tab=repositories"
	resp, err := http.Post("http://localhost:8080/", "text/plain", bytes.NewBufferString(url))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode >= 300 { // TODO can be done better
		t.Fatalf("Received bad status code from server: %d", resp.StatusCode)
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
