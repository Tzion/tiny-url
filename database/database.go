package database

type database struct {
	shortToLongUrl map[string]string
}

var Database database

func (db database) init() {
	Database.shortToLongUrl = make(map[string]string)
}

func (db database) FetchLongUrl(shortUrl string) *string {
	longUrl, exists := db.shortToLongUrl[shortUrl]
	if exists {
		return &longUrl
	} else {
		return nil
	}
}
