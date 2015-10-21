package main

import (
	"net/http"
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

func TestMain(t *testing.T) {
	s := httptest.NewServer(createBaseHandler())
	defer s.Close()
	resp, err := http.Get(s.URL)
	if err != nil {
		t.Fatalf("Failed to query index \n", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("StatusCode is not 200")
	}
}
