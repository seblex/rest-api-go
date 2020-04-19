package store

import (
	"fmt"
	"strings"
	"testing"
)

// Test Store
func TestStore(t *testing.T, databaseuRL string) (*Store, func(... string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseUrl = databaseuRL
	s := New(config)
	err := s.Open()
	if err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		_, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ",")))
		if err != nil {
			t.Fatal(err)
		}

		s.Close()
	}
}
