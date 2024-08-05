package tugassatu

import (
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}
