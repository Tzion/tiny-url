package database

import (
	"testing"
)

func TestDatabase(t *testing.T) {
	t.Log(Database)
	Database.init()
	Database.shortToLongUrl["a"] = "A"
	t.Log(Database)
}
