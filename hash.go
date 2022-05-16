package main

import (
	"crypto/sha256"
	// "encoding/base64"
	"hash"
	// "hash/fnv"
)

var hashMethod hash.Hash = sha256.New()

var sha hash.Hash = sha256.New()

func ShortenUrl(url string) string {
	hashMethod.Reset()
	hashMethod.Write([]byte(url))
	shortUrl := hashMethod.Sum(nil)
	return string(shortUrl)
	// return base64.StdEncoding.EncodeToString(shortUrl)
}
