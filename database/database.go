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

var Database database = database{
	shortToLongUrl: make(map[string]string),
}
var initialized bool

// func init() {
// 	setup()
// }

func setup() {
	if initialized {
		return
	}
	Database.shortToLongUrl = make(map[string]string)
	initialized = true
}

func reset() {
	initialized = false
	setup()
}

func Size() int {
	if !initialized {
		setup()
	}
	return len(Database.shortToLongUrl)
}

func FetchLongUrl(shortUrl string) *string {
	if !initialized {
		setup()
	}
	longUrl, exists := Database.shortToLongUrl[shortUrl]
	if exists {
		return &longUrl
	} else {
		return nil
	}
}

func Insert(shortUrl string, longUrl string) {
	if !initialized {
		setup()
	}
	currentLongUrl, exists := Database.shortToLongUrl[shortUrl]
	if exists {
		if currentLongUrl == longUrl {
			return
		} else {
			log.Printf("Collision upon entry %s. Overriding with latest value", shortUrl)
			Database.shortToLongUrl[shortUrl] = longUrl
		}
	}
	Database.shortToLongUrl[shortUrl] = longUrl
}
