package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	IndexHandler(w, nil)
	s := `{"message":"Hello world!"}`
	msg := strings.TrimSpace(w.Body.String())
	if msg != s {
		t.Fatalf("Return body (%s) does not match expected (%s)\n", msg, s)
	}

}
