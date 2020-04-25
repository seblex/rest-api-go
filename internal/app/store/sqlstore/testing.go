package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

// Test DB
func TestDB(t *testing.T, databaseuRL string) (*sql.DB, func(... string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseuRL)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
