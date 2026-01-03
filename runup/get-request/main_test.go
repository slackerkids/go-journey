package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=Bob", nil)

	w := httptest.NewRecorder()

	HelloHandler(w, req)

	got := w.Body.String()
	want := `{"message": "Hello, Bob"}`

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
