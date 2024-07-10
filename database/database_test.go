package database

import "testing"

func TestDbConnection(t *testing.T) {
	db := Connect()
	err := db.Ping()
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
}
