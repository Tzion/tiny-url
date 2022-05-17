package database

/**
TODO
	* Find elegant solution to initialization
**/

import (
	"log"
)

type database struct {
	shortToLongUrl map[string]string
}

var Database database
var initialized bool

func (db database) init() {
	if initialized {
		return
	}
	Database.shortToLongUrl = make(map[string]string)
}

func (db database) len() int {
	if !initialized {
		Database.init()
	}
	return len(Database.shortToLongUrl)
}

func (db database) FetchLongUrl(shortUrl string) *string {
	if !initialized {
		Database.init()
	}
	longUrl, exists := db.shortToLongUrl[shortUrl]
	if exists {
		return &longUrl
	} else {
		return nil
	}
}

func (db database) Insert(shortUrl string, longUrl string) {
	if !initialized {
		Database.init()
	}
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
