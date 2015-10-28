package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInvalidDomain(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "localhost/service/ip?domain=dsfkfs", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ipHandler(w, r)
	s := `{"Error":"Invalid domain address."}`
	msg := strings.TrimSpace(w.Body.String())
	if msg != s {
		t.Fatalf("Return body (%s) does not match expected (%s)\n", msg, s)
	}
}

func TestEmptyDomain(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "localhost/service/ip", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ipHandler(w, r)
	s := `{"Error":"Empty domain parameter"}`
	msg := strings.TrimSpace(w.Body.String())
	if msg != s {
		t.Fatalf("Return body (%s) does not match expected (%s)\n", msg, s)
	}
}
func TestNodomain(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "localhost/service/ip?domain=", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ipHandler(w, r)
	s := `{"Error":"Empty domain parameter"}`
	msg := strings.TrimSpace(w.Body.String())
	if msg != s {
		t.Fatalf("Return body (%s) does not match expected (%s)\n", msg, s)
	}
}

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "localhost/service/ip?domain=localhost", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	ipHandler(w, r)
	s := `{"IPs":["127.0.0.1"]}`
	msg := strings.TrimSpace(w.Body.String())
	if msg != s {
		t.Fatalf("Return body (%s) does not match expected (%s)\n", msg, s)
	}
}

func TestWeb(t *testing.T) {
	s := httptest.NewServer(createBaseHandler())
	defer s.Close()
	resp, err := http.Get(s.URL)
	if err != nil {
		t.Fatal("Failed to query index", err)
	}
	if resp.StatusCode != 200 {
		t.Fatal("StatusCode is not 200")
	}
}
