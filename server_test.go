package main

import (
	"bytes"
	"net/http"
	"testing"
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

func TestPost(t *testing.T) {
	stringData := "URL"
	resp, err := http.Post("http://localhost:8080/tzions12", "text/plain", bytes.NewBufferString(stringData))
	t.Log(resp)
	if err != nil {
		t.Error(err)
	}
}
