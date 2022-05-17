package database

import (
	"testing"
)

func TestDatabaseInit(t *testing.T) {
	Database.init()
	if Database.shortToLongUrl == nil {
		t.Error("Database is not initialize")
	}
	if Database.len() != 0 {
		t.Error("Database is not empty")
	}
}

func TestInsert(t *testing.T) {
	Database.init()
	t.Log(Database.len())
	Database.Insert("shortUrl", "longUrl")
	t.Log(Database)
	Database.init()
	t.Log(Database)
}
