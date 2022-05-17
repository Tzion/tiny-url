package main

import (
	"encoding/base64"
	"hash"
	"hash/fnv"
)

var hashMethod hash.Hash = fnv.New64a() // FNV-1a if chosen due to low collision rate and high execution speed

func ShortenUrl(url string) string {
	hashedUrl := hashUrl(url)
	encodedUrl := encode(hashedUrl)
	return encodedUrl
}

func hashUrl(url string) []byte {
	hashMethod.Reset()
	hashMethod.Write([]byte(url))
	hashedUrl := hashMethod.Sum(nil)
	return hashedUrl
}

func encode(hashed []byte) string {
	return base64.StdEncoding.EncodeToString(hashed)
}
