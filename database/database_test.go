package database

import (
	"testing"
)

// TODO add teardown/setup to cleanup database

// func TestDatabaseInit(t *testing.T) {
// 	Database.init()
// 	if Database.shortToLongUrl == nil {
// 		t.Error("Database is not initialize")
// 	}
// 	if Database.len() != 0 {
// 		t.Error("Database is not empty")
// 	}
// }

func TestInsert(t *testing.T) {
	Insert("shortUrl_1", "longUrl_1")
	Insert("shortUrl_2", "longUrl_2")
	if entries := Size(); entries != 2 {
		t.Errorf("Database have wrong number of entries\nhave %d\nwant %d\n", entries, 2)
	}
}

/**
func TestReset(t *testing.T) {
	// Database.Insert("shortUrl_1", "longUrl_1")
	// Database.Insert("shortUrl_2", "longUrl_2")
	reset()
	// if entries := Database.len(); entries != 0 {
	// 	t.Error("reset() didn't clean the database")
	// }
}

func TestFetchLongUrl(t *testing.T) {
	Database.reset()
	short, long := "shortUrl_Fetch", "longUrl_Fetch"
	Database.Insert(short, long)
	if inside := Database.FetchLongUrl(short); *inside != long {
		t.Errorf("Url was not fetched correctly\nhave %s\nwant %s\n", *inside, long)
	}
}

func TestCollision(t *testing.T) {
	short, long1, long2 := "sameShortUrl", "LongUrl_1", "LongUrl_2"
	Database.Insert(short, long1)
	Database.Insert(short, long2)
	if inside := Database.FetchLongUrl(short); *inside != long2 {
		t.Errorf("Url was not fetched correctly\nhave %s\nwant %s\n", *inside, long2)
	}
}
**/
