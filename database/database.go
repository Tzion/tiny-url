package database

/**
	Adding the reversed mapping (longToShortUrl) could be taken as consideration,
	it can be handy to save calculation of the hash in cost of size (Time vs. Space)
	For sake of simplicity I chose to keep single mapping at the moment.
**/

import (
	"log"
)

type database struct {
	shortToLongUrl map[string]string
}

var Database database = database{
	shortToLongUrl: make(map[string]string),
}

func reset() {
	Database.shortToLongUrl = make(map[string]string)
}

func Size() int {
	return len(Database.shortToLongUrl)
}

func FetchLongUrl(shortUrl string) *string {
	longUrl, exists := Database.shortToLongUrl[shortUrl]
	if exists {
		return &longUrl
	} else {
		return nil
	}
}

func Insert(shortUrl string, longUrl string) {
	currentLongUrl, exists := Database.shortToLongUrl[shortUrl]
	if exists {
		if currentLongUrl == longUrl {
			return
		} else {
			log.Printf("Collision in entry %q. Overriding with latest value.", shortUrl)
			Database.shortToLongUrl[shortUrl] = longUrl
		}
	}
	Database.shortToLongUrl[shortUrl] = longUrl
}
