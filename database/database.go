package database

import (
	"log"
)

type database struct {
	shortToLongUrl map[string]string
}

var Database database

func (db database) init() {
	Database.shortToLongUrl = make(map[string]string)
}

func (db database) len() int {
	return len(Database.shortToLongUrl)
}

func (db database) FetchLongUrl(shortUrl string) *string {
	longUrl, exists := db.shortToLongUrl[shortUrl]
	if exists {
		return &longUrl
	} else {
		return nil
	}
}

func (db database) Insert(shortUrl string, longUrl string) {
	currentLongUrl, exists := db.shortToLongUrl[shortUrl]
	if exists {
		if currentLongUrl == longUrl {
			return
		} else {
			log.Panicf("Collision upon entry %s. Overriding with latest value", shortUrl)
			db.shortToLongUrl[shortUrl] = longUrl
		}
	}
	db.shortToLongUrl[shortUrl] = longUrl
}
