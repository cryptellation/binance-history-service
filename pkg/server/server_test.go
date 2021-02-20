package server

import "testing"

func TestNew(t *testing.T) {
	// Create server
	s := New(nil)

	// Check if the server is not null
	if s == nil {
		t.Fatal("New server is nil")
	}

	// Check if version is correct
	if s.Version() != version {
		t.Error("Version does not match")
	}
}
