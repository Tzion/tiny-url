package shortener

import (
	"testing"
)

var urls = []string{
	"https://pkg.go.dev/hash/fnv@go1.18.2#New128",
	"https://gobyexample.com/base64-encoding",
	"https://stackoverflow.com/questions/25304080/callback-in-golang",
	"https://golangr.com/golang-http-server/",
	"https://gobyexample.com/base64-encoding",
}

func TestHashedUrlLength(t *testing.T) {
	for _, url := range urls {
		t.Run(string(url), func(t *testing.T) {
			shortUrl := hashUrl(url)
			if urlLen, hashLen := len(shortUrl), hashMethod.Size(); urlLen != hashLen {
				t.Errorf("length of short url is wrong\nhave %d\nwant %d", urlLen, hashLen)
			}
		})
	}
}
func TestDeterminism(t *testing.T) {
	for _, url := range urls {
		t.Run(string(url), func(t *testing.T) {
			if first, second := ShortenUrl(url), ShortenUrl(url); first != second {
				t.Errorf("2 calls over the same url generate different results\nfirst=%s\nsecond=%s", first, second)
			}
		})
	}
}
