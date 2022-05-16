package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	url := "https://pkg.go.dev/hash/fnv@go1.18.2#New128"
	url2 := "https://gobyexample.com/base64-encoding"
	t.Logf("%x", hashMethod.Sum([]byte(url)))
	// t.Logf("%x", hashMethod.Sum([]byte(url)))
	t.Logf("%x", sha.Sum([]byte(url)))
	t.Logf("%x", sha.Sum([]byte(url2)))
	sha.Write([]byte(url2))
	t.Logf("%x", sha.Sum(nil))
	sha.Reset()
	t.Logf("%x", sha.Sum([]byte(url)))
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
