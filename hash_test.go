package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	shortUrl := ShortenUrl("https://pkg.go.dev/hash/fnv@go1.18.2#New128")
	t.Log(shortUrl)
}

func TestShortUrlLength(t *testing.T) {
	shortUrl := ShortenUrl("https://gobyexample.com/base64-encoding")
	if urlLen, hashLen := len(shortUrl), hashMethod.Size(); urlLen != hashLen {
		t.Errorf("length of short url is wrong\nhave %d\nwant %d", urlLen, hashLen)
	}
}

func TestDeterminism(t *testing.T) {
	url := "https://golangr.com/golang-http-server/"
	if first, second := ShortenUrl(url), ShortenUrl(url); first != second {
		t.Errorf("2 calls over the same url generate different results\nfirst=%s\nsecond=%s", first, second)
	}
}
